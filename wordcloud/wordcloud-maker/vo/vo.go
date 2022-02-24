package vo

//go:generate go-doudou name --file $GOFILE

type MakePayload struct {
	// 文本文件公开下载链接
	// text file public download url
	SrcUrl string `json:"srcUrl"`
	// 文本语言
	// text language
	Lang string `json:"lang"`
	// 只取词频前Top个词
	// only want Top frequent words
	Top int `json:"top"`
}
