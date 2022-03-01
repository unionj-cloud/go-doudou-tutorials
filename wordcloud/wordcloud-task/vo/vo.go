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

type TaskVo struct {
	Id       int
	SrcUrl   string
	ImgUrl   string
	Lang     string
	Top      int
	Status   string
	Error    string
	UserId   int
	CreateAt string
}

type Order struct {
	Col  string `json:"col"`
	Sort string `json:"sort"`
}

type Page struct {
	// 排序规则
	Orders []Order `json:"orders"`
	// 页码
	PageNo int `json:"pageNo"`
	// 每页行数
	Size int `json:"size"`
}

// 分页筛选条件
type PageQuery struct {
	Page Page `json:"page"`
}

type PageRet struct {
	Items    interface{} `json:"items"`
	PageNo   int         `json:"pageNo"`
	PageSize int         `json:"pageSize"`
	Total    int         `json:"total"`
	HasNext  bool        `json:"hasNext"`
}

type TaskPageRet struct {
}
