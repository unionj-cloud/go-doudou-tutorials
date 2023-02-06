package vo

//go:generate go-doudou name --file $GOFILE --strategy snake
//go:generate go-doudou enum --file $GOFILE

type PageFilter struct {
	// 真实姓名，前缀匹配
	Name string `json:"name"`
	// 所属部门ID
	Dept int `json:"dept"`
}

type Order struct {
	Col  string `json:"col"`
	Sort string `json:"sort"`
}

type Page struct {
	// 排序规则
	Orders []Order `json:"orders"`
	// 页码
	PageNo int `json:"page_no"`
	// 每页行数
	Size int `json:"size"`
}

// 分页筛选条件
type PageQuery struct {
	Filter PageFilter `json:"filter"`
	Page   Page       `json:"page"`
}

type PageRet struct {
	Items    interface{} `json:"items"`
	PageNo   int         `json:"page_no"`
	PageSize int         `json:"page_size"`
	Total    int         `json:"total"`
	HasNext  bool        `json:"has_next"`
}

type UserVo struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Phone string `json:"phone"`
	Dept  string `json:"dept"`
}

type RoleEnum int

const (
	GUEST RoleEnum = iota
	USER
	ADMIN
)

type KeyboardLayout int

const (
	UNKNOWN KeyboardLayout = iota
	QWERTZ
	AZERTY
	QWERTY
)

type Auth struct {
	User string `json:"user"`
	Pass string `json:"pass"`
}

type UserStore map[Auth]RoleEnum
