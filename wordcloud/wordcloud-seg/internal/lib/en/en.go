package en

import (
	"context"
	"github.com/jdkato/prose/v2"
	"github.com/unionj-cloud/go-doudou-tutorials/wordcloud/wordcloud-seg/internal/lib/nlp"
	"github.com/unionj-cloud/go-doudou/v2/toolkit/sliceutils"
	"sort"
	"unicode/utf8"
)

type EnglishTokenizer struct {
}

var Pos = []string{"NNP", "NNPS", "NNS", "NN"}

// https://github.com/jdkato/prose#tagging
var posdict = map[string]string{
	"#":    "number sign",
	"$":    "currency",
	"CC":   "conjunction, coordinating",
	"CD":   "cardinal number",
	"DT":   "determiner",
	"EX":   "existential there",
	"FW":   "foreign word",
	"IN":   "conjunction, subordinating or preposition",
	"JJ":   "adjective",
	"JJR":  "adjective, comparative",
	"JJS":  "adjective, superlative",
	"LS":   "list item marker",
	"MD":   "verb, modal auxiliary",
	"NN":   "noun, singular or mass",
	"NNP":  "noun, proper singular",
	"NNPS": "noun, proper plural",
	"NNS":  "noun, plural",
	"PDT":  "predeterminer",
	"POS":  "possessive ending",
	"PRP":  "pronoun, personal",
	"PRP$": "pronoun, possessive",
	"RB":   "adverb",
	"RBR":  "adverb, comparative",
	"RBS":  "adverb, superlative",
	"RP":   "adverb, particle",
	"SYM":  "symbol",
	"TO":   "infinitival to",
	"UH":   "interjection",
	"VB":   "verb, base form",
	"VBD":  "verb, past tense",
	"VBG":  "verb, gerund or present participle",
	"VBN":  "verb, past participle",
	"VBP":  "verb, non-3rd person singular present",
	"VBZ":  "verb, 3rd person singular present",
	"WDT":  "wh-determiner",
	"WP":   "wh-pronoun, personal",
	"WP$":  "wh-pronoun, possessive",
	"WRB":  "wh-adverb",
}

func (self *EnglishTokenizer) DoSeg(ctx context.Context, s string, pos []string, top int, min int) (rs nlp.WordFreqResult, err error) {
	if len(pos) == 0 {
		pos = Pos
	}

	var doc *prose.Document

	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
		doc, err = prose.NewDocument(s)
		if err != nil {
			return nil, err
		}
	}

	rr := &nlp.SegResponse{}
	var result nlp.WordFreqResult
	tokenmap := make(map[nlp.Token]float64)

	for _, item := range doc.Tokens() {
		if min > 0 && utf8.RuneCountInString(item.Text) < min {
			continue
		}
		if len(pos) > 0 && !sliceutils.StringContains(pos, item.Tag) {
			continue
		}
		token := nlp.Token{
			item.Text,
			item.Tag,
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
