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
