package bgControllers

import (
	"github.com/astaxie/beego"
	"strconv"
	"fmt"
	"GiantTech/models"
	"GiantTech/controllers/tools"
)

type BgProjectDeleteController struct {
	beego.Controller
}

func (this *BgProjectDeleteController) Prepare() {
	s := this.StartSession()
	username = s.Get("login")
	beego.Informational(username)
	if username == nil {
		this.Ctx.Redirect(302, "/login")
	}
}

func (this *BgProjectDeleteController) Post() {
	user, _ := models.GetTUsersByName(username.(string))
	this.Data["User"] = user
	p := this.Ctx.Request.FormValue("ProjectId")
	id, _ := strconv.Atoi(p)
	project, _ := models.GetTProjectsById(id)
	if project.ProjectStatus == -1 {
		if err := models.DeleteTProjects(id); err == nil {
			tools.AddLog(user.UserRealName, tools.Delete, tools.Project, project.ProjectName + "成功！")
			fmt.Fprint(this.Ctx.ResponseWriter, "成功")
		}else {
			beego.Error(err)
			fmt.Fprint(this.Ctx.ResponseWriter, "删除失败，请重试！")
		}
	}else {
		project.ProjectStatus = -1
		if err := models.UpdateTProjectsById(project); err == nil {
			tools.AddLog(user.UserRealName, tools.Pass, tools.Project, project.ProjectName + "成功！")
			fmt.Fprint(this.Ctx.ResponseWriter, "刷新")
		}else {
			beego.Error(err)
			fmt.Fprint(this.Ctx.ResponseWriter, "放弃失败，请重试！")
		}
	}
}