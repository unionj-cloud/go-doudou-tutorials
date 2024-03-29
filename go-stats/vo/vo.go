/**
* Generated by go-doudou v2.0.3.
* You can edit it as your need.
 */
package vo

import "github.com/shopspring/decimal"

//go:generate go-doudou name --file $GOFILE

// request vo
type PercentageReqVo struct {
	// key value pairs
	Data []PercentageVo `json:"data"`
	// digit number after dot
	Places int `json:"places"`
}

// key value pair
type PercentageVo struct {
	// number for something
	Value int `json:"value"`
	// unique key for distinguishing something
	Key string `json:"key"`
}

// result vo
type PercentageRespVo struct {
	Value   int     `json:"value"`
	Key     string  `json:"key"`
	Percent float64 `json:"percent"`
	// formatted percentage
	PercentFormatted string `json:"percentFormatted"`
}

//go:generate go-doudou enum --file $GOFILE

type KeyboardLayout int

const (
	UNKNOWN KeyboardLayout = iota
	QWERTZ
	AZERTY
	QWERTY
)

type DecimalWrapper struct {
	Data decimal.Decimal `json:"data"`
}
