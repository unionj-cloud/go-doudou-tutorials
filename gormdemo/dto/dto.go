/**
* Generated by go-doudou v2.0.5.
* You can edit it as your need.
 */
package dto

//go:generate go-doudou name --file $GOFILE

type UserDto struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Phone string `json:"phone"`
	Dept  string `json:"dept"`
}

// Page result wrapper
type Page struct {
	Items      interface{} `json:"items"`
	Page       int64       `json:"page"`
	Size       int64       `json:"size"`
	MaxPage    int64       `json:"maxPage"`
	TotalPages int64       `json:"totalPages"`
	Total      int64       `json:"total"`
	Last       bool        `json:"last"`
	First      bool        `json:"first"`
	Visible    int64       `json:"visible"`
}

// Parameter struct
type Parameter struct {
	Page    string      `json:"page"`
	Size    string      `json:"size"`
	Sort    string      `json:"sort"`
	Order   string      `json:"order"`
	Fields  string      `json:"fields"`
	Filters interface{} `json:"filters"`
}
