package bgControllers

import (
	"github.com/astaxie/beego"
	"strconv"
	"GiantTech/models"
)

type BgProjectInfoController struct {
	beego.Controller
}

func (this *BgProjectInfoController) Prepare() {
	s := this.StartSession()
	username = s.Get("login")
	beego.Informational(username)
	if username == nil {
		this.Ctx.Redirect(302, "/login")
	}
}

func (this *BgProjectInfoController) Get() {
	user, _ := models.GetTUsersByName(username.(string))
	this.Data["User"] = user
	id := this.Ctx.Request.FormValue("id")
	projectId, _ := strconv.Atoi(id)
	project, _ := models.GetTProjectsById(projectId)
	this.Data["Project"] = project
	this.TplName = "bgview/projectinfo.html"
}