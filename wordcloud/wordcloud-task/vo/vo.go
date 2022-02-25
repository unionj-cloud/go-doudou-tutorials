package vo

//go:generate go-doudou name --file $GOFILE

type TaskPayload struct {
	// 用户ID
	// user ID
	UserId int `json:"userId"`
	// source file url 原文件链接
	SrcUrl string `json:"srcUrl"`
	Lang   string `json:"lang"`
	Top    int    `json:"top"`
}

type TaskSuccess struct {
	// task ID
	// 任务ID
	TaskId int `json:"taskId"`
	// word cloud image url
	// 词云图链接
	ImgUrl string `json:"imgUrl"`
}

type TaskFail struct {
	// task ID
	// 任务ID
	TaskId int `json:"taskId"`
	// error message
	// 错误信息
	Error string `json:"error"`
}
