package dto

import "github.com/starter-go/urls/src/main/golang/lib/data/dxo"

// Example ... VO
type Example struct {
	ID dxo.ExampleID `json:"id"`
	Base

	Foo string `json:"foo"`
	Bar int    `json:"bar"`
}
