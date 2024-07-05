package ilinks

import (
	"bytes"
	"context"
	"crypto/sha1"
	"fmt"
	"time"

	"github.com/starter-go/base/lang"
	"github.com/starter-go/urls/src/main/golang/lib/classes/links"
	"github.com/starter-go/urls/src/main/golang/lib/data/dao"
	"github.com/starter-go/urls/src/main/golang/lib/data/dxo"
	"github.com/starter-go/urls/src/main/golang/lib/data/entity"
	"github.com/starter-go/urls/src/main/golang/lib/web/dto"
)

// ShortLinkServiceImpl ...
type ShortLinkServiceImpl struct {

	//starter:component
	_as func(links.Service) //starter:as("#")

	BasrURL       string //starter:inject("${starter.urls.base}")
	DefaultMaxAge int64  //starter:inject("${starter.urls.default-max-age}")

	Dao dao.ShortLinkDAO //starter:inject("#")
}

func (inst *ShortLinkServiceImpl) _impl() links.Service {
	return inst
}

// // Find ...
// func (inst *ShortLinkServiceImpl) Find(db *gorm.DB, id dxo.ExampleID) (*entity.Example, error) {
// 	m := inst.model()
// 	db = inst.Agent.DB(db)
// 	res := db.First(m, id)
// 	return inst.makeResult(m, res)
// }

// Insert ...
func (inst *ShortLinkServiceImpl) Insert(c context.Context, item *dto.ShortLink) (*dto.ShortLink, error) {
	o2 := links.ConvertD2E(item)
	inst.prepareItem(o2)
	o3, err := inst.Dao.Insert(nil, o2)
	if err != nil {
		return nil, err
	}
	o4 := links.ConvertE2D(o3)
	return o4, nil
}

// Update ...
func (inst *ShortLinkServiceImpl) Update(c context.Context, id dxo.ShortLinkID, item *dto.ShortLink) (*dto.ShortLink, error) {
	return nil, fmt.Errorf("no impl")
}

// Delete ...
func (inst *ShortLinkServiceImpl) Delete(c context.Context, id dxo.ShortLinkID) error {
	return fmt.Errorf("no impl")
}

// Find ...
func (inst *ShortLinkServiceImpl) Find(c context.Context, id dxo.ShortLinkID) (*dto.ShortLink, error) {
	o1, err := inst.Dao.Find(nil, id)
	if err != nil {
		return nil, err
	}
	o2 := links.ConvertE2D(o1)
	return o2, nil
}

// FindByMagic ...
func (inst *ShortLinkServiceImpl) FindByMagic(c context.Context, id dxo.ShortLinkMagic) (*dto.ShortLink, error) {
	o1, err := inst.Dao.FindByMagic(nil, id)
	if err != nil {
		return nil, err
	}
	if inst.isItemOK(o1) {
		return nil, fmt.Errorf("not found")
	}
	o2 := links.ConvertE2D(o1)
	return o2, nil
}

func (inst *ShortLinkServiceImpl) isItemOK(item *entity.ShortLink) bool {
	if item == nil {
		return false
	}
	if !item.Enabled {
		return false
	}
	t1 := item.CreatedAt
	t2 := t1.Add(item.MaxAge.Duration())
	now := time.Now()
	return now.After(t1) && now.Before(t2)
}

func (inst *ShortLinkServiceImpl) prepareItem(item *entity.ShortLink) {

	now := lang.Now()
	raw := item.Raw
	maxage := item.MaxAge
	magic := dxo.ShortLinkMagic("")

	// check age
	if maxage < 1 {
		maxage = lang.Milliseconds(inst.DefaultMaxAge)
	}

	// make magic
	builder := &bytes.Buffer{}
	builder.WriteString(raw)
	builder.WriteString("&now=" + now.String())
	data := builder.Bytes()
	sum := sha1.Sum(data)
	b64 := lang.Base64FromBytes(sum[0:12])
	magic = dxo.ShortLinkMagic(b64)

	// make short-link
	short := inst.BasrURL + magic.String()

	// update
	item.Magic = magic
	item.Short = short
	item.MaxAge = maxage
}
