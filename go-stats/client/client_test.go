package client

import (
	"context"
	"encoding/json"
	"github.com/go-resty/resty/v2"
	"github.com/unionj-cloud/go-doudou/v2/framework/registry"
	"github.com/unionj-cloud/go-doudou/v2/framework/restclient"
	"go-doudou-tutorials/go-stats/vo"
	"os"
	"reflect"
	"testing"
)

func TestGoStatsClient_LargestRemainder(t *testing.T) {
	os.Setenv("GOSTATS", "http://localhost:6060")
	defer os.Unsetenv("GOSTATS")
	jsonPayload := `{"data":[{"value":20,"key":"apple"},{"value":30,"key":"banana"},{"value":40,"key":"orange"}],"places":2}`
	var payload vo.PercentageReqVo
	json.Unmarshal([]byte(jsonPayload), &payload)
	type fields struct {
		provider registry.IServiceProvider
		client   *resty.Client
		rootPath string
	}
	type args struct {
		ctx      context.Context
		_headers map[string]string
		payload  vo.PercentageReqVo
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		want_resp *resty.Response
		wantData  []vo.PercentageRespVo
		wantErr   bool
	}{
		{
			name: "",
			fields: fields{
				provider: nil,
				client:   restclient.NewClient(),
				rootPath: "",
			},
			args: args{
				ctx:      context.Background(),
				_headers: nil,
				payload:  payload,
			},
			want_resp: nil,
			wantData:  nil,
			wantErr:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			receiver := NewGoStatsClient()
			got_resp, gotData, err := receiver.LargestRemainder(tt.args.ctx, tt.args._headers, tt.args.payload, Options{
				GzipReqBody: true,
			})
			if (err != nil) != tt.wantErr {
				t.Errorf("LargestRemainder() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got_resp, tt.want_resp) {
				t.Errorf("LargestRemainder() got_resp = %v, want %v", got_resp, tt.want_resp)
			}
			if !reflect.DeepEqual(gotData, tt.wantData) {
				t.Errorf("LargestRemainder() gotData = %v, want %v", gotData, tt.wantData)
			}
		})
	}
}
