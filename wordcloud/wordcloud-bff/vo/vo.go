package vo

//go:generate go-doudou name --file $GOFILE

// UploadResult 上传文件接口返回对象
type UploadResult struct {
	// 任务ID
	TaskId int `json:"taskId"`
	// 原文件链接
	SrcUrl string `json:"srcUrl"`
	// 词云图链接
	ImgUrl string `json:"imgUrl"`
}
