package vo

//go:generate go-doudou name --file $GOFILE

type SegPayload struct {
	Text string `json:"text"`
}

type SegResult struct {
	WordFreq WordFreqResult
}

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
