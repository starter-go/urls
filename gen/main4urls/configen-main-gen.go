package main4urls

import "github.com/starter-go/application"

func nop(a ... any) {    
}

func registerComponents(cr application.ComponentRegistry) error {
    ac:=&autoRegistrar{}
    ac.init(cr)
    return ac.addAll()
}

type comFactory interface {
    register(cr application.ComponentRegistry) error
}

type autoRegistrar struct {
    cr application.ComponentRegistry
}

func (inst *autoRegistrar) init(cr application.ComponentRegistry) {
	inst.cr = cr
}

func (inst *autoRegistrar) register(factory comFactory) error {
	return factory.register(inst.cr)
}

func (inst*autoRegistrar) addAll() error {

    
    inst.register(&p2f326772a6_ilinks_ShortLinkServiceImpl{})
    inst.register(&p9d1242fce2_controllers_ShortLinkController{})
    inst.register(&pb9ef5d19a5_idao_ShortLinkDaoImpl{})
    inst.register(&pf97a3dad43_database_ThisGroup{})


    return nil
}
