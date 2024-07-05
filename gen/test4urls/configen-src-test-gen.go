package test4urls
import (
    p162238ebb "github.com/starter-go/urls/src/test/golang/unit"
     "github.com/starter-go/application"
)

// type p162238ebb.DemoUnit in package:github.com/starter-go/urls/src/test/golang/unit
//
// id:com-162238ebbccd7ba7-unit-DemoUnit
// class:class-0dc072ed44b3563882bff4e657a52e62-Units
// alias:
// scope:singleton
//
type p162238ebbc_unit_DemoUnit struct {
}

func (inst* p162238ebbc_unit_DemoUnit) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-162238ebbccd7ba7-unit-DemoUnit"
	r.Classes = "class-0dc072ed44b3563882bff4e657a52e62-Units"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p162238ebbc_unit_DemoUnit) new() any {
    return &p162238ebb.DemoUnit{}
}

func (inst* p162238ebbc_unit_DemoUnit) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p162238ebb.DemoUnit)
	nop(ie, com)

	


    return nil
}


