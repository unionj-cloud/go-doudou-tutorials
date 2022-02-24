package nlp

import "context"

type WordFreqResult [][]interface{}

func (self WordFreqResult) Len() int {
	return len(self)
}

func (self WordFreqResult) Less(i, j int) bool {
	return self[i][2].(float64) > self[j][2].(float64)
}

func (self WordFreqResult) Swap(i, j int) {
	self[i], self[j] = self[j], self[i]
}

type SegRequest struct {
	Text string   `json:"text"`
	Pos  []string `json:"pos"`
	Top  int      `json:"top"`
	Min  int      `json:"min"`
}

type SegResponse struct {
	Result WordFreqResult `json:"result"`
}

type Token struct {
	Word string
	Pos  string
}

type Tokenizer interface {
	DoSeg(ctx context.Context, s string, pos []string, top int, min int) (rs WordFreqResult, err error)
}
