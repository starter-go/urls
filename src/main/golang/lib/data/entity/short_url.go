package entity

import (
	"github.com/starter-go/base/lang"
	"github.com/starter-go/urls/src/main/golang/lib/data/dxo"
)

// ShortLink ...
type ShortLink struct {
	ID dxo.ShortLinkID

	Base

	Magic dxo.ShortLinkMagic `gorm:"unique"`

	Short   string
	Raw     string
	MaxAge  lang.Milliseconds
	Enabled bool
}
