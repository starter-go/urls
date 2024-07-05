package main4urls
import (
    pd1a916a20 "github.com/starter-go/libgin"
    p512a30914 "github.com/starter-go/libgorm"
    p9621e8b71 "github.com/starter-go/security/random"
    p8e5339560 "github.com/starter-go/urls/src/main/golang/lib/classes/links"
    p819d3b4bd "github.com/starter-go/urls/src/main/golang/lib/data/dao"
    pf97a3dad4 "github.com/starter-go/urls/src/main/golang/lib/data/database"
    p94da56571 "github.com/starter-go/urls/src/main/golang/lib/data/dxo"
    pb9ef5d19a "github.com/starter-go/urls/src/main/golang/lib/implements/idao"
    p2f326772a "github.com/starter-go/urls/src/main/golang/lib/implements/ilinks"
    p9d1242fce "github.com/starter-go/urls/src/main/golang/lib/web/controllers"
     "github.com/starter-go/application"
)

// type pf97a3dad4.ThisGroup in package:github.com/starter-go/urls/src/main/golang/lib/data/database
//
// id:com-f97a3dad43750a22-database-ThisGroup
// class:class-512a309140d0ad99eb1c95c8dc0d02f9-GroupRegistry
// alias:alias-94da5657115062e8d6f9fc9b6084f5c1-DatabaseAgent
// scope:singleton
//
type pf97a3dad43_database_ThisGroup struct {
}

func (inst* pf97a3dad43_database_ThisGroup) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-f97a3dad43750a22-database-ThisGroup"
	r.Classes = "class-512a309140d0ad99eb1c95c8dc0d02f9-GroupRegistry"
	r.Aliases = "alias-94da5657115062e8d6f9fc9b6084f5c1-DatabaseAgent"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pf97a3dad43_database_ThisGroup) new() any {
    return &pf97a3dad4.ThisGroup{}
}

func (inst* pf97a3dad43_database_ThisGroup) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pf97a3dad4.ThisGroup)
	nop(ie, com)

	
    com.Enabled = inst.getEnabled(ie)
    com.Alias = inst.getAlias(ie)
    com.URI = inst.getURI(ie)
    com.Prefix = inst.getPrefix(ie)
    com.Source = inst.getSource(ie)
    com.SourceManager = inst.getSourceManager(ie)


    return nil
}


func (inst*pf97a3dad43_database_ThisGroup) getEnabled(ie application.InjectionExt)bool{
    return ie.GetBool("${datagroup.default.enabled}")
}


func (inst*pf97a3dad43_database_ThisGroup) getAlias(ie application.InjectionExt)string{
    return ie.GetString("${datagroup.default.alias}")
}


func (inst*pf97a3dad43_database_ThisGroup) getURI(ie application.InjectionExt)string{
    return ie.GetString("${datagroup.default.uri}")
}


func (inst*pf97a3dad43_database_ThisGroup) getPrefix(ie application.InjectionExt)string{
    return ie.GetString("${datagroup.default.table-name-prefix}")
}


func (inst*pf97a3dad43_database_ThisGroup) getSource(ie application.InjectionExt)string{
    return ie.GetString("${datagroup.default.datasource}")
}


func (inst*pf97a3dad43_database_ThisGroup) getSourceManager(ie application.InjectionExt)p512a30914.DataSourceManager{
    return ie.GetComponent("#alias-512a309140d0ad99eb1c95c8dc0d02f9-DataSourceManager").(p512a30914.DataSourceManager)
}



// type pb9ef5d19a.ShortLinkDaoImpl in package:github.com/starter-go/urls/src/main/golang/lib/implements/idao
//
// id:com-b9ef5d19a5f97ee8-idao-ShortLinkDaoImpl
// class:
// alias:alias-819d3b4bd24fda2dea60ee582e554109-ShortLinkDAO
// scope:singleton
//
type pb9ef5d19a5_idao_ShortLinkDaoImpl struct {
}

func (inst* pb9ef5d19a5_idao_ShortLinkDaoImpl) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-b9ef5d19a5f97ee8-idao-ShortLinkDaoImpl"
	r.Classes = ""
	r.Aliases = "alias-819d3b4bd24fda2dea60ee582e554109-ShortLinkDAO"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pb9ef5d19a5_idao_ShortLinkDaoImpl) new() any {
    return &pb9ef5d19a.ShortLinkDaoImpl{}
}

func (inst* pb9ef5d19a5_idao_ShortLinkDaoImpl) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pb9ef5d19a.ShortLinkDaoImpl)
	nop(ie, com)

	
    com.Agent = inst.getAgent(ie)
    com.UUIDService = inst.getUUIDService(ie)


    return nil
}


func (inst*pb9ef5d19a5_idao_ShortLinkDaoImpl) getAgent(ie application.InjectionExt)p94da56571.DatabaseAgent{
    return ie.GetComponent("#alias-94da5657115062e8d6f9fc9b6084f5c1-DatabaseAgent").(p94da56571.DatabaseAgent)
}


func (inst*pb9ef5d19a5_idao_ShortLinkDaoImpl) getUUIDService(ie application.InjectionExt)p9621e8b71.UUIDService{
    return ie.GetComponent("#alias-9621e8b71013b0fc25942a1749ed3652-UUIDService").(p9621e8b71.UUIDService)
}



// type p2f326772a.ShortLinkServiceImpl in package:github.com/starter-go/urls/src/main/golang/lib/implements/ilinks
//
// id:com-2f326772a6acb0f9-ilinks-ShortLinkServiceImpl
// class:
// alias:alias-8e5339560ba50986c3246c299d38a6ab-Service
// scope:singleton
//
type p2f326772a6_ilinks_ShortLinkServiceImpl struct {
}

func (inst* p2f326772a6_ilinks_ShortLinkServiceImpl) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-2f326772a6acb0f9-ilinks-ShortLinkServiceImpl"
	r.Classes = ""
	r.Aliases = "alias-8e5339560ba50986c3246c299d38a6ab-Service"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p2f326772a6_ilinks_ShortLinkServiceImpl) new() any {
    return &p2f326772a.ShortLinkServiceImpl{}
}

func (inst* p2f326772a6_ilinks_ShortLinkServiceImpl) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p2f326772a.ShortLinkServiceImpl)
	nop(ie, com)

	
    com.BasrURL = inst.getBasrURL(ie)
    com.DefaultMaxAge = inst.getDefaultMaxAge(ie)
    com.Dao = inst.getDao(ie)


    return nil
}


func (inst*p2f326772a6_ilinks_ShortLinkServiceImpl) getBasrURL(ie application.InjectionExt)string{
    return ie.GetString("${starter.urls.base}")
}


func (inst*p2f326772a6_ilinks_ShortLinkServiceImpl) getDefaultMaxAge(ie application.InjectionExt)int64{
    return ie.GetInt64("${starter.urls.default-max-age}")
}


func (inst*p2f326772a6_ilinks_ShortLinkServiceImpl) getDao(ie application.InjectionExt)p819d3b4bd.ShortLinkDAO{
    return ie.GetComponent("#alias-819d3b4bd24fda2dea60ee582e554109-ShortLinkDAO").(p819d3b4bd.ShortLinkDAO)
}



// type p9d1242fce.ShortLinkController in package:github.com/starter-go/urls/src/main/golang/lib/web/controllers
//
// id:com-9d1242fce2169832-controllers-ShortLinkController
// class:class-d1a916a203352fd5d33eabc36896b42e-Controller
// alias:
// scope:singleton
//
type p9d1242fce2_controllers_ShortLinkController struct {
}

func (inst* p9d1242fce2_controllers_ShortLinkController) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-9d1242fce2169832-controllers-ShortLinkController"
	r.Classes = "class-d1a916a203352fd5d33eabc36896b42e-Controller"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p9d1242fce2_controllers_ShortLinkController) new() any {
    return &p9d1242fce.ShortLinkController{}
}

func (inst* p9d1242fce2_controllers_ShortLinkController) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p9d1242fce.ShortLinkController)
	nop(ie, com)

	
    com.Sender = inst.getSender(ie)
    com.Links = inst.getLinks(ie)


    return nil
}


func (inst*p9d1242fce2_controllers_ShortLinkController) getSender(ie application.InjectionExt)pd1a916a20.Responder{
    return ie.GetComponent("#alias-d1a916a203352fd5d33eabc36896b42e-Responder").(pd1a916a20.Responder)
}


func (inst*p9d1242fce2_controllers_ShortLinkController) getLinks(ie application.InjectionExt)p8e5339560.Service{
    return ie.GetComponent("#alias-8e5339560ba50986c3246c299d38a6ab-Service").(p8e5339560.Service)
}


