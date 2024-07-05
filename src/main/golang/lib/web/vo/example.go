package vo

import "github.com/starter-go/urls/src/main/golang/lib/web/dto"

// Example ... VO
type Example struct {
	Base

	Items []*dto.Example `json:"items"`
}
