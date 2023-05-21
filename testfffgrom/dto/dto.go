/**
* Generated by go-doudou v2.1.0.
* You can edit it as your need.
 */
package dto

//go:generate go-doudou name --file $GOFILE --form

type User struct {
	Id    int32  `json:"id" form:"id"`
	Name  string `json:"name" form:"name"`
	Phone string `json:"phone" form:"phone"`
	Dept  string `json:"dept" form:"dept"`
}

// Page result wrapper
type Page struct {
	Items      []interface{} `json:"items" form:"items"`
	Page       int64         `json:"page" form:"page"`
	Size       int64         `json:"size" form:"size"`
	MaxPage    int64         `json:"maxPage" form:"maxPage"`
	TotalPages int64         `json:"totalPages" form:"totalPages"`
	Total      int64         `json:"total" form:"total"`
	Last       bool          `json:"last" form:"last"`
	First      bool          `json:"first" form:"first"`
	Visible    int64         `json:"visible" form:"visible"`
}

// Parameter struct
type Parameter struct {
	Page    string        `json:"page" form:"page"`
	Size    string        `json:"size" form:"size"`
	Sort    string        `json:"sort" form:"sort"`
	Order   string        `json:"order" form:"order"`
	Fields  string        `json:"fields" form:"fields"`
	Filters []interface{} `json:"filters" form:"filters"`
}
