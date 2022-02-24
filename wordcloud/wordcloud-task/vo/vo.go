package vo

//go:generate go-doudou name --file $GOFILE

type TaskSuccess struct {
	// task ID
	// 任务ID
	TaskId int
	// word cloud image url
	// 词云图链接
	ImgUrl string
}

type TaskFail struct {
	// task ID
	// 任务ID
	TaskId int
	// error message
	// 错误信息
	Error string
}
