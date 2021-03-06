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
	Id       int    `json:"id"`
	SrcUrl   string `json:"srcUrl"`
	ImgUrl   string `json:"imgUrl"`
	Lang     string `json:"lang"`
	Top      int    `json:"top"`
	Status   int    `json:"status"`
	Error    string `json:"error"`
	UserId   int    `json:"userId"`
	CreateAt string `json:"createAt"`
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

type Filter struct {
	UserId int `json:"userId"`
}

// 分页筛选条件
type PageQuery struct {
	Filter Filter `json:"filter"`
	Page   Page   `json:"page"`
}

type PageRet struct {
	Items    interface{} `json:"items"`
	PageNo   int         `json:"pageNo"`
	PageSize int         `json:"pageSize"`
	Total    int         `json:"total"`
	HasNext  bool        `json:"hasNext"`
}

type TaskPageRet struct {
	PageRet
	Items []TaskVo `json:"items"`
}
