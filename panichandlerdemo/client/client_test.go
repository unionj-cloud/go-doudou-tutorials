package client

import (
	"context"
	"github.com/go-resty/resty/v2"
	"github.com/unionj-cloud/go-doudou/v2/framework/registry"
	"os"
	"reflect"
	"testing"
	"testsvc/dto"
)

func TestTestsvcClient_GetBookPage(t *testing.T) {
	os.Setenv("TESTSVC", "http://localhost:6060")
	defer func() {
		os.Unsetenv("TESTSVC")
	}()
	type fields struct {
		provider registry.IServiceProvider
		client   *resty.Client
		rootPath string
	}
	type args struct {
		ctx      context.Context
		_headers map[string]string
		name     string
		author   string
		page     dto.Page
		options  Options
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		want_resp *resty.Response
		wantErr   bool
	}{
		{
			name:   "",
			fields: fields{},
			args: args{
				ctx:      context.Background(),
				_headers: nil,
				name:     "aaa",
				author:   "bbb",
				page: dto.Page{
					Orders: []dto.Order{
						{
							Col:  "createAt",
							Sort: "desc",
						},
						{
							Col:  "author",
							Sort: "asc",
						},
					},
					PageNo: 2,
					Size:   10,
					Filter: dto.PageFilter{
						BookName:  "ccc",
						BookShelf: 123,
					},
				},
				options: Options{},
			},
			want_resp: nil,
			wantErr:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			receiver := NewTestsvcClient()
			got_resp, err := receiver.GetBookPage(tt.args.ctx, tt.args._headers, tt.args.name, tt.args.author, tt.args.page, tt.args.options)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetBookPage() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got_resp, tt.want_resp) {
				t.Errorf("GetBookPage() got_resp = %v, want %v", got_resp, tt.want_resp)
			}
		})
	}
}
