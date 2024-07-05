package urls

import (
	"github.com/starter-go/application"
	"github.com/starter-go/libgin/modules/libgin"
	"github.com/starter-go/libgorm/modules/libgorm"
	"github.com/starter-go/module-gorm-mysql/modules/mysql"
	"github.com/starter-go/module-gorm-sqlserver/modules/sqlserver"
	"github.com/starter-go/security/modules/security"
	"github.com/starter-go/urls"
	"github.com/starter-go/urls/gen/main4urls"
	"github.com/starter-go/urls/gen/test4urls"
)

// Module  ...
func Module() application.Module {
	mb := urls.NewMainModule()
	mb.Components(main4urls.ExportComponents)

	mb.Depend(libgin.Module())
	mb.Depend(libgorm.Module())
	mb.Depend(security.Module())

	mb.Depend(mysql.Module())
	mb.Depend(sqlserver.Module())

	return mb.Create()
}

// ModuleForTest ...
func ModuleForTest() application.Module {
	mb := urls.NewTestModule()
	mb.Components(test4urls.ExportComponents)
	return mb.Create()
}
