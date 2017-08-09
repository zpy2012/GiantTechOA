package bgControllers

import (
	"github.com/astaxie/beego"
)

type BgIndexController struct {
	beego.Controller
}



func (this *BgIndexController) Prepare() {
	s := this.StartSession()
	username = s.Get("login")
	beego.Informational(username)
	if username == nil {
		this.Ctx.Redirect(302, "/login")
	}
}

func (this *BgIndexController) Get() {
	this.Ctx.Redirect(302, "/project/?page=1&projectname=all")
}