package raftx

import (
	"context"
	"fmt"
	transport "github.com/Jille/raft-grpc-transport"
	"github.com/hashicorp/raft"
	"google.golang.org/grpc"
	"net"
	"os"
	"path/filepath"
	"sync/atomic"
	"time"
)

var RaftSrv RaftService

type RaftService struct {
	TcpAddr      *net.TCPAddr
	Servers      []raft.Server
	cfg          *raft.Config
	noticeCh     chan bool
	node         *raft.Raft
	startFunc    func()
	stopFunc     func()
	leaderStatus int32
	checkTimer   *time.Timer
}

func NewRaft(ctx context.Context, myID, myAddress string, fsm raft.FSM) (*raft.Raft, *transport.Manager, error) {
	c := raft.DefaultConfig()
	c.LocalID = raft.ServerID(myID)

	baseDir := filepath.Join(*raftDir, myID)

	ldb, err := boltdb.NewBoltStore(filepath.Join(baseDir, "logs.dat"))
	if err != nil {
		return nil, nil, fmt.Errorf(`boltdb.NewBoltStore(%q): %v`, filepath.Join(baseDir, "logs.dat"), err)
	}

	sdb, err := boltdb.NewBoltStore(filepath.Join(baseDir, "stable.dat"))
	if err != nil {
		return nil, nil, fmt.Errorf(`boltdb.NewBoltStore(%q): %v`, filepath.Join(baseDir, "stable.dat"), err)
	}

	fss, err := raft.NewFileSnapshotStore(baseDir, 3, os.Stderr)
	if err != nil {
		return nil, nil, fmt.Errorf(`raft.NewFileSnapshotStore(%q, ...): %v`, baseDir, err)
	}

	tm := transport.New(raft.ServerAddress(myAddress), []grpc.DialOption{grpc.WithInsecure()})

	r, err := raft.NewRaft(c, fsm, ldb, sdb, fss, tm.Transport())
	if err != nil {
		return nil, nil, fmt.Errorf("raft.NewRaft: %v", err)
	}

	if *raftBootstrap {
		cfg := raft.Configuration{
			Servers: []raft.Server{
				{
					Suffrage: raft.Voter,
					ID:       raft.ServerID(myID),
					Address:  raft.ServerAddress(myAddress),
				},
			},
		}
		f := r.BootstrapCluster(cfg)
		if err := f.Error(); err != nil {
			return nil, nil, fmt.Errorf("raft.Raft.BootstrapCluster: %v", err)
		}
	}

	return r, tm, nil
}

func (this *RaftService) Stop() {
	if this.node != nil {
		this.node.Shutdown()
		close(this.noticeCh)
	}
}

func (this *RaftService) IsLeader() bool {
	if atomic.LoadInt32(&this.leaderStatus) == 1 {
		return true
	} else {
		return false
	}
}

func (this *RaftService) watch() {
	for {
		select {
		case status, ok := <-this.noticeCh:
			if !ok {
				return
			}
			if status {
				atomic.StoreInt32(&this.leaderStatus, 1)
				this.startFunc()
			} else {
				atomic.StoreInt32(&this.leaderStatus, 0)
				this.stopFunc()
				if this.checkTimer != nil {
					this.checkTimer.Stop()
					this.checkTimer = nil
				}
			}
		}
	}
}
