package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/starter-go/libgin"
	"github.com/starter-go/urls/src/main/golang/lib/classes/links"
	"github.com/starter-go/urls/src/main/golang/lib/data/dxo"
	"github.com/starter-go/urls/src/main/golang/lib/web/dto"
	"github.com/starter-go/urls/src/main/golang/lib/web/vo"
)

// ShortLinkController ...
type ShortLinkController struct {

	//starter:component
	_as func(libgin.Controller) //starter:as(".")

	Sender libgin.Responder //starter:inject("#")
	Links  links.Service    //starter:inject("#")
}

func (inst *ShortLinkController) _impl() libgin.Controller {
	return inst
}

// Registration ...
func (inst *ShortLinkController) Registration() *libgin.ControllerRegistration {
	return &libgin.ControllerRegistration{Route: inst.route}
}

func (inst *ShortLinkController) route(rp libgin.RouterProxy) error {

	rp = rp.For("urls")

	rp.POST("", inst.handlePostInsert)
	rp.DELETE(":id", inst.handle)
	rp.GET(":id", inst.handleGetLinkInfo)

	rp.GET("/urls/:magic", inst.handleGetRedirectToRawLink)

	return nil
}

func (inst *ShortLinkController) handle(c *gin.Context) {
	req := &myShortLinkRequest{
		context:       c,
		controller:    inst,
		wantRequestID: false,
	}
	req.execute(req.doNOP)
}

func (inst *ShortLinkController) handlePostInsert(c *gin.Context) {
	req := &myShortLinkRequest{
		context:         c,
		controller:      inst,
		wantRequestBody: true,
	}
	req.execute(req.doInsert)
}

func (inst *ShortLinkController) handleGetLinkInfo(c *gin.Context) {
	req := &myShortLinkRequest{
		context:       c,
		controller:    inst,
		wantRequestID: true,
	}
	req.execute(req.doGetLinkInfo)
}

func (inst *ShortLinkController) handleGetRedirectToRawLink(c *gin.Context) {
	req := &myShortLinkRequest{
		context:          c,
		controller:       inst,
		wantRequestMagic: true,
	}
	err := req.open()
	if err == nil {
		err = req.doGetLink()
	}
	if err != nil {
		// send error
		req.send(err)
	} else {
		// send redirect
		item := req.body2.Items[0]
		location := item.Raw
		code := http.StatusTemporaryRedirect
		ctyp := "text/plain"
		text := "redirect to " + location
		c.Header("Location", location)
		c.Data(code, ctyp, []byte(text))
	}
}

////////////////////////////////////////////////////////////////////////////////

type myShortLinkRequest struct {
	context    *gin.Context
	controller *ShortLinkController

	wantRequestID    bool
	wantRequestMagic bool
	wantRequestBody  bool

	id    dxo.ShortLinkID
	magic dxo.ShortLinkMagic

	body1 vo.ShortLinks
	body2 vo.ShortLinks
}

func (inst *myShortLinkRequest) open() error {

	c := inst.context

	if inst.wantRequestID {
		idstr := c.Param("id")
		idnum, err := strconv.ParseInt(idstr, 10, 64)
		if err != nil {
			return err
		}
		inst.id = dxo.ShortLinkID(idnum)
	}

	if inst.wantRequestMagic {
		magic := c.Param("magic")
		// idnum, err := strconv.ParseInt(idstr, 10, 64)
		// if err != nil { 			return err 		}
		inst.magic = dxo.ShortLinkMagic(magic)
	}

	if inst.wantRequestBody {
		err := c.BindJSON(&inst.body1)
		if err != nil {
			return err
		}
	}

	return nil
}

func (inst *myShortLinkRequest) send(err error) {
	data := &inst.body2
	code := inst.body2.Status
	resp := new(libgin.Response)
	resp.Context = inst.context
	resp.Error = err
	resp.Data = data
	resp.Status = code
	inst.controller.Sender.Send(resp)
}

func (inst *myShortLinkRequest) execute(fn func() error) {
	err := inst.open()
	if err == nil {
		err = fn()
	}
	inst.send(err)
}

func (inst *myShortLinkRequest) doNOP() error {
	return nil
}

func (inst *myShortLinkRequest) doInsert() error {
	ctx := inst.context
	ser := inst.controller.Links
	o1 := inst.body1.Items[0]
	o2, err := ser.Insert(ctx, o1)
	if err != nil {
		return err
	}
	inst.body2.Items = []*dto.ShortLink{o2}
	return nil
}

// doGetLink 取短链接信息
func (inst *myShortLinkRequest) doGetLinkInfo() error {
	id := inst.id
	ctx := inst.context
	ser := inst.controller.Links
	o1, err := ser.Find(ctx, id)
	if err != nil {
		return err
	}
	inst.body2.Items = []*dto.ShortLink{o1}
	return nil
}

// doGetLink 取短链接本身
func (inst *myShortLinkRequest) doGetLink() error {
	magic := inst.magic
	c := inst.context
	ser := inst.controller.Links
	o1, err := ser.FindByMagic(c, magic)
	if err != nil {
		return err
	}
	inst.body2.Items = []*dto.ShortLink{o1}
	return nil
}
