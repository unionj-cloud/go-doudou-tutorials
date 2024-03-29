/**
* Generated by go-doudou v2.1.5.
* You can edit it as your need.
*/
package dto

//go:generate go-doudou name --file $GOFILE --form

type GddUser struct {
	Id    int32
	Name  string
	Phone string
	Dept  string
}

// Page result wrapper
type Page struct {
	Items      []interface{}
	Page       int64
	Size       int64
	MaxPage    int64
	TotalPages int64
	Total      int64
	Last       bool
	First      bool
	Visible    int64
}

// Parameter struct
type Parameter struct {
	Page    string
	Size    string
	Sort    string
	Order   string
	Fields  string
	Filters []interface{}
}
