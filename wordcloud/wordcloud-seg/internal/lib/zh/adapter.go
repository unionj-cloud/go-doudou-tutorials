package zh

import (
	"context"
	"github.com/unionj-cloud/go-doudou-tutorials/wordcloud/wordcloud-seg/internal/lib/nlp"
	"github.com/unionj-cloud/go-doudou/v2/toolkit/sliceutils"
	thulac "github.com/unionj-cloud/thulacgo"
	"sort"
	"strings"
	"unicode/utf8"
)

var Pos = []string{"n", "np", "ns", "ni", "nz", "uw"}

// n/名词 np/人名 ns/地名 ni/机构名 nz/其它专名
//m/数词 q/量词 mq/数量词 t/时间词 f/方位词 s/处所词
//v/动词 a/形容词 d/副词 h/前接成分 k/后接成分
//i/习语 j/简称 r/代词 c/连词 p/介词 u/助词 y/语气助词
//e/叹词 o/拟声词 g/语素 w/标点 x/其它
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

func (self *GoThulac) DoSeg(ctx context.Context, s string, pos []string, top int, min int) (rs nlp.WordFreqResult, err error) {
	if len(pos) == 0 {
		pos = Pos
	}

	var words []string

	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
		words = self.Lac.SegToSlice(s)
	}

	rr := &nlp.SegResponse{}
	var result nlp.WordFreqResult
	tokenmap := make(map[nlp.Token]float64)

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
		token := nlp.Token{
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
		result = append(result, []interface{}{k.Word, posdict[k.Pos], v})
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
