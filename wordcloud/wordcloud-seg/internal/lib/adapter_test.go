package lib_test

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/unionj-cloud/go-doudou-tutorials/wordcloud/wordcloud-seg/internal/lib"
	"github.com/unionj-cloud/go-doudou-tutorials/wordcloud/wordcloud-seg/vo"
	"github.com/unionj-cloud/go-doudou/toolkit/pathutils"
	thulac "github.com/unionj-cloud/thulacgo"
	"path/filepath"
	"testing"
)

var testDir string

func init() {
	testDir = pathutils.Abs("../../testdata")
}

func TestGoThulac_DoSeg(t *testing.T) {
	type fields struct {
		lac *thulac.Thulacgo
		sep string
	}
	type args struct {
		ctx context.Context
		s   string
		pos []string
		top int
		min int
	}
	modelpath := filepath.Join(testDir, "models")
	userpath := filepath.Join(testDir, "userword.txt")
	lac := thulac.NewThulacgo(modelpath, userpath, false, false, false, '_')
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantRs  vo.WordFreqResult
		wantErr bool
	}{
		{
			name: "test1",
			fields: fields{
				lac,
				"_",
			},
			args: args{
				ctx: context.Background(),
				s:   "最近，勇士老将伊戈达拉道出了实情！",
				pos: []string{"ns", "uw",
					"n",
					"a",
					"np",
					"ni",
					"nz"},
				top: -1,
				min: -1,
			},
			wantRs: [][]interface{}{
				{
					"勇士", "名词", float64(1),
				},
				{
					"老将", "名词", float64(1),
				},
				{
					"伊戈达拉", "人名", float64(1),
				},
				{
					"道出了", "保留词", float64(1),
				},
				{
					"实情", "名词", float64(1),
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			self := &lib.GoThulac{
				Lac: tt.fields.lac,
				Sep: tt.fields.sep,
			}
			gotRs, err := self.DoSeg(tt.args.ctx, tt.args.s, tt.args.pos, tt.args.top, tt.args.min)
			if (err != nil) != tt.wantErr {
				t.Errorf("DoSeg() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.ElementsMatch(t, gotRs, tt.wantRs, "DoSeg() gotRs = %v, want %v", gotRs, tt.wantRs)
		})
	}
}
