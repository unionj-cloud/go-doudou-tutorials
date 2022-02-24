package lib

import (
	"context"
	"github.com/unionj-cloud/go-doudou-tutorials/wordcloud/wordcloud-seg/vo"
	"github.com/unionj-cloud/go-doudou/toolkit/sliceutils"
	thulac "github.com/unionj-cloud/thulacgo"
	"sort"
	"strings"
	"unicode/utf8"
)

type SegRequest struct {
	Text string   `json:"text"`
	Pos  []string `json:"pos"`
	Top  int      `json:"top"`
	Min  int      `json:"min"`
}

type SegResponse struct {
	Result vo.WordFreqResult `json:"result"`
}

var Pos = []string{"n", "a", "np", "ns", "ni", "nz", "uw", "v"}

var posdict = map[string]string{
	"n":  "名词",
	"a":  "形容词",
	"np": "人名",
	"ns": "地名",
	"ni": "机构名",
	"nz": "其他专名",
	"uw": "保留词",
	"v":  "动词",
}

type token struct {
	word string
	pos  string
}

type Adapter interface {
	DoSeg(ctx context.Context, s string, pos []string, top int, min int) (rs vo.WordFreqResult, err error)
}

type GoThulac struct {
	Lac *thulac.Thulacgo
	Sep string
}

func NewGoThulac(modelpath string, userpath string, justseg bool, t2s bool, ufilter bool, separator string) *GoThulac {
	lac := thulac.NewThulacgo(modelpath, userpath, justseg, t2s, ufilter, []byte(separator)[0])
	return &GoThulac{
		lac,
		separator,
	}
}

func (self *GoThulac) DoSeg(ctx context.Context, s string, pos []string, top int, min int) (rs vo.WordFreqResult, err error) {
	var words []string

	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
		words = self.Lac.SegToSlice(s)
	}

	rr := &SegResponse{}
	var result vo.WordFreqResult
	tokenmap := make(map[token]float64)

	for _, item := range words {
		splits := strings.Split(item, self.Sep)
		w := splits[0]
		if min > 0 && utf8.RuneCountInString(w) < min {
			continue
		}
		p := splits[1]
		if len(pos) > 0 && !sliceutils.StringContains(pos, p) {
			continue
		}
		token := token{
			w,
			p,
		}
		if old, ok := tokenmap[token]; !ok {
			tokenmap[token] = 1
		} else {
			old++
			tokenmap[token] = old
		}
	}
	for k, v := range tokenmap {
		result = append(result, []interface{}{k.word, posdict[k.pos], v})
	}

	sort.Sort(result)

	if top <= 0 || top > len(result) {
		top = len(result)
	}

	rr.Result = result[:top]
	return rr.Result, err
}

func (self *GoThulac) Free() {
	self.Lac.Deinit()
}
