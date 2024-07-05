package dao

import (
	"github.com/starter-go/urls/src/main/golang/lib/data/dxo"
	"github.com/starter-go/urls/src/main/golang/lib/data/entity"
	"gorm.io/gorm"
)

// ShortLinkDAO ...
type ShortLinkDAO interface {
	Insert(db *gorm.DB, item *entity.ShortLink) (*entity.ShortLink, error)

	Update(db *gorm.DB, id dxo.ShortLinkID, fn func(*entity.ShortLink)) (*entity.ShortLink, error)

	Delete(db *gorm.DB, id dxo.ShortLinkID) error

	Find(db *gorm.DB, id dxo.ShortLinkID) (*entity.ShortLink, error)
	FindByMagic(db *gorm.DB, magic dxo.ShortLinkMagic) (*entity.ShortLink, error)
}
