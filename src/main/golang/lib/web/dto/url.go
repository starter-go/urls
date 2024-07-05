package dto

import (
	"github.com/starter-go/base/lang"
	"github.com/starter-go/urls/src/main/golang/lib/data/dxo"
)

// ShortLink ... DTO
type ShortLink struct {
	ID dxo.ShortLinkID `json:"id"`

	Base

	Raw     string             `json:"raw"`
	Short   string             `json:"short"`
	Magic   dxo.ShortLinkMagic `json:"magic"`
	MaxAge  lang.Milliseconds  `json:"maxage"`
	Enabled bool               `json:"enabled"`
}
