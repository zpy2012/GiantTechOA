package bgControllers

import (
	"github.com/astaxie/beego"
	"strconv"
	"GiantTech/models"
	"GiantTech/controllers/tools"
)

type BgUserDeleteController struct {
	beego.Controller
}

func (this *BgUserDeleteController) Prepare() {
	s := this.StartSession()
	username = s.Get("login")
	beego.Informational(username)
	if username == nil {
		this.Ctx.Redirect(302, "/login")
	}
}

func (this *BgUserDeleteController) Get() {
	p := this.Ctx.Request.FormValue("id")
	id, _ := strconv.Atoi(p)
	user, _ := models.GetTUsersByName(username.(string))
	this.Data["User"] = user
	delUser, _ := models.GetTUsersById(id)
	if err := models.DeleteTUsers(id); err != nil {
		beego.Error(err)
		msg := delUser.UserName + "失败！"
		tools.AddLog(user.UserRealName, tools.Delete, tools.Person, msg)
	}else {
		msg := delUser.UserName + "成功！"
		tools.AddLog(user.UserRealName, tools.Delete, tools.Person, msg)
	}
	this.Ctx.Redirect(302, "/user/?page=1")
}