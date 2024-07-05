package links

import (
	"github.com/starter-go/security-gorm/rbacdb"
	"github.com/starter-go/urls/src/main/golang/lib/data/entity"
	"github.com/starter-go/urls/src/main/golang/lib/web/dto"
)

// ConvertD2E ...
func ConvertD2E(src *dto.ShortLink) *entity.ShortLink {

	dst := &entity.ShortLink{}
	dst.ID = src.ID
	rbacdb.CopyBaseFieldsFromDtoToEntity(&src.BaseDTO, &dst.BaseEntity)

	dst.Raw = src.Raw
	dst.Short = src.Short
	dst.Enabled = src.Enabled
	dst.Magic = src.Magic
	dst.MaxAge = src.MaxAge

	return dst
}

// ConvertE2D 。。。
func ConvertE2D(src *entity.ShortLink) *dto.ShortLink {

	dst := &dto.ShortLink{}
	dst.ID = src.ID
	rbacdb.CopyBaseFieldsFromEntityToDTO(&src.BaseEntity, &dst.BaseDTO)

	dst.Raw = src.Raw
	dst.Short = src.Short
	dst.Enabled = src.Enabled
	dst.Magic = src.Magic
	dst.MaxAge = src.MaxAge

	return dst
}
