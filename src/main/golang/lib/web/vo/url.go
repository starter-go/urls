package vo

import "github.com/starter-go/urls/src/main/golang/lib/web/dto"

// ShortLinks ... VO
type ShortLinks struct {
	Base

	Items []*dto.ShortLink `json:"urls"`
}
