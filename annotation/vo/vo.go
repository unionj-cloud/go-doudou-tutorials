package vo

import "encoding/json"

//go:generate go-doudou name --file $GOFILE

type PageFilter struct {
	// 真实姓名，前缀匹配
	Name string
	// 所属部门ID
	Dept int
}

type Order struct {
	Col  string
	Sort string
}

type Page struct {
	// 排序规则
	Orders []Order
	// 页码
	PageNo int
	// 每页行数
	Size int
}

// 分页筛选条件
type PageQuery struct {
	Filter PageFilter
	Page   Page
}

type PageRet struct {
	Items    interface{}
	PageNo   int
	PageSize int
	Total    int
	HasNext  bool
}

type UserVo struct {
	Id    int
	Name  string
	Phone string
	Dept  string
}

type RoleEnum int

const (
	GUEST RoleEnum = iota
	USER
	ADMIN
)

func (k *RoleEnum) StringSetter(value string) {
	switch value {
	case "GUEST":
		*k = GUEST
	case "USER":
		*k = USER
	case "ADMIN":
		*k = ADMIN
	default:
		*k = GUEST
	}
}

func (k *RoleEnum) StringGetter() string {
	switch *k {
	case GUEST:
		return "GUEST"
	case USER:
		return "USER"
	case ADMIN:
		return "ADMIN"
	default:
		return "GUEST"
	}
}

func (k *RoleEnum) UnmarshalJSON(bytes []byte) error {
	var _k string
	err := json.Unmarshal(bytes, &_k)
	if err != nil {
		return err
	}
	k.StringSetter(_k)
	return nil
}

func (k RoleEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(k.StringGetter())
}

type Auth struct {
	User string
	Pass string
}

type UserStore map[Auth]RoleEnum
