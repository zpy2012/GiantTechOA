package bgControllers

import (
	"github.com/astaxie/beego"
	"GiantTech/models"
)

type BgUserInfoController struct {
	beego.Controller
}

func (this *BgUserInfoController) Prepare() {
	s := this.StartSession()
	username = s.Get("login")
	beego.Informational(username)
	if username == nil {
		this.Ctx.Redirect(302, "/login")
	}
}

func (this *BgUserInfoController) Get() {
	user, _ := models.GetTUsersByName(username.(string))
	this.Data["User"] = user
	this.TplName = "bgview/userinfo.html"
}