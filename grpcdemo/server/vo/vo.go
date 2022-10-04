package vo

import "encoding/json"

//go:generate go-doudou name --file $GOFILE

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
	PageNo int `json:"pageNo"`
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
	PageNo   int         `json:"pageNo"`
	PageSize int         `json:"pageSize"`
	Total    int         `json:"total"`
	HasNext  bool        `json:"hasNext"`
}

type UserVo struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Phone string `json:"phone"`
	Dept  string `json:"dept"`
}

type KeyboardLayout int

const (
	UNKNOWN KeyboardLayout = iota
	QWERTZ
	AZERTY
	QWERTY
)

func (k *KeyboardLayout) StringSetter(value string) {
	switch value {
	case "UNKNOWN":
		*k = UNKNOWN
	case "QWERTY":
		*k = QWERTY
	case "QWERTZ":
		*k = QWERTZ
	case "AZERTY":
		*k = AZERTY
	default:
		*k = UNKNOWN
	}
}

func (k *KeyboardLayout) StringGetter() string {
	switch *k {
	case UNKNOWN:
		return "UNKNOWN"
	case QWERTY:
		return "QWERTY"
	case QWERTZ:
		return "QWERTZ"
	case AZERTY:
		return "AZERTY"
	default:
		return "UNKNOWN"
	}
}

func (k *KeyboardLayout) UnmarshalJSON(bytes []byte) error {
	var _k string
	err := json.Unmarshal(bytes, &_k)
	if err != nil {
		return err
	}
	k.StringSetter(_k)
	return nil
}

func (k KeyboardLayout) MarshalJSON() ([]byte, error) {
	return json.Marshal(k.StringGetter())
}

type Keyboard struct {
	Layout  KeyboardLayout `json:"layout"`
	Backlit bool           `json:"backlit"`
}
