package links

import (
	"context"

	"github.com/starter-go/urls/src/main/golang/lib/data/dxo"
	"github.com/starter-go/urls/src/main/golang/lib/web/dto"
)

// Service ...
type Service interface {
	Insert(c context.Context, item *dto.ShortLink) (*dto.ShortLink, error)

	Update(c context.Context, id dxo.ShortLinkID, item *dto.ShortLink) (*dto.ShortLink, error)

	Delete(c context.Context, id dxo.ShortLinkID) error

	Find(c context.Context, id dxo.ShortLinkID) (*dto.ShortLink, error)
	FindByMagic(c context.Context, id dxo.ShortLinkMagic) (*dto.ShortLink, error)
}
