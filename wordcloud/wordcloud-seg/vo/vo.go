package vo

//go:generate go-doudou name --file $GOFILE

type Lang string

const (
	Zh Lang = "zh"
	En Lang = "en"
)

type SegPayload struct {
	// 文本
	Text string `json:"text"`
	// 文本语言
	// 仅支持中文和英文
	// text language
	// only support zh and en
	Lang string `json:"lang"`
}

type SegResult struct {
	WordFreq [][]interface{}
}
