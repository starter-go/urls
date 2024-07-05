package idao

import (
	"github.com/starter-go/security/random"
	"github.com/starter-go/urls/src/main/golang/lib/data/dao"
	"github.com/starter-go/urls/src/main/golang/lib/data/dxo"
	"github.com/starter-go/urls/src/main/golang/lib/data/entity"
	"gorm.io/gorm"
)

// ShortLinkDaoImpl ...
type ShortLinkDaoImpl struct {

	//starter:component
	_as func(dao.ShortLinkDAO) //starter:as("#")

	Agent       dxo.DatabaseAgent  //starter:inject("#")
	UUIDService random.UUIDService //starter:inject("#")

}

func (inst *ShortLinkDaoImpl) _impl() dao.ShortLinkDAO {
	return inst
}

func (inst *ShortLinkDaoImpl) model() *entity.ShortLink {
	return new(entity.ShortLink)
}

func (inst *ShortLinkDaoImpl) modelList() []*entity.ShortLink {
	return make([]*entity.ShortLink, 0)
}

func (inst *ShortLinkDaoImpl) makeResult(o *entity.ShortLink, res *gorm.DB) (*entity.ShortLink, error) {
	err := res.Error
	if err != nil {
		return nil, err
	}
	return o, nil
}

// Find ...
func (inst *ShortLinkDaoImpl) Find(db *gorm.DB, id dxo.ShortLinkID) (*entity.ShortLink, error) {
	m := inst.model()
	db = inst.Agent.DB(db)
	res := db.First(m, id)
	return inst.makeResult(m, res)
}

// FindByMagic ...
func (inst *ShortLinkDaoImpl) FindByMagic(db *gorm.DB, magic dxo.ShortLinkMagic) (*entity.ShortLink, error) {
	m := inst.model()
	db = inst.Agent.DB(db)
	res := db.Where("magic=?", magic).First(m)
	return inst.makeResult(m, res)
}

// Insert ...
func (inst *ShortLinkDaoImpl) Insert(db *gorm.DB, item *entity.ShortLink) (*entity.ShortLink, error) {

	uuid := inst.UUIDService.Build().Class("entity.ShortLink").Generate()

	item.ID = 0
	item.UUID = uuid

	db = inst.Agent.DB(db)
	res := db.Create(item)
	return inst.makeResult(item, res)
}

// Update ...
func (inst *ShortLinkDaoImpl) Update(db *gorm.DB, id dxo.ShortLinkID, fn func(*entity.ShortLink)) (*entity.ShortLink, error) {

	item := inst.model()
	db = inst.Agent.DB(db)
	res := db.First(item, id)
	if res.Error != nil {
		return nil, res.Error
	}

	fn(item)
	res = db.Save(item)
	return inst.makeResult(item, res)
}

// Delete ...
func (inst *ShortLinkDaoImpl) Delete(db *gorm.DB, id dxo.ShortLinkID) error {
	item := inst.model()
	db = inst.Agent.DB(db)
	res := db.Delete(item, id)
	return res.Error
}
