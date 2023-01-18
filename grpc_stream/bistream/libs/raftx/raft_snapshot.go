package raftx

import (
	"github.com/hashicorp/raft"
)

type Snapshot struct {
}

func (s *Snapshot) Persist(sink raft.SnapshotSink) error {
	sink.Close()
	return nil
}
func (s *Snapshot) Release() {}
