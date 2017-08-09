package bgControllers

import (
	"github.com/astaxie/beego"
	"strconv"
	"GiantTech/models"
	"GiantTech/controllers/tools"
	"fmt"
)

type BgProjectEditController struct {
	beego.Controller
}

func (this *BgProjectEditController) Prepare() {
	s := this.StartSession()
	username = s.Get("login")
	beego.Informational(username)
	if username == nil {
		this.Ctx.Redirect(302, "/login")
	}
}

func (this *BgProjectEditController) Post() {
	user, _ := models.GetTUsersByName(username.(string))
	i := this.Ctx.Request.FormValue("ProjectId")
	projectId, _ := strconv.Atoi(i)
	editType := this.Ctx.Request.FormValue("Type")
	project, _ := models.GetTProjectsById(projectId)
	switch editType {
	case "newName":{
		newName := this.Ctx.Request.FormValue("Name")
		project.ProjectName = newName
	}
	case "newDescription":{
		description := this.Ctx.Request.FormValue("Description")
		project.ProjectDescription = description
	}
	case "newSource":{
		s := this.Ctx.Request.FormValue("Source")
		source, _ := strconv.Atoi(s)
		project.ProjectSource = source
	}
	case "newType":{
		pt := this.Ctx.Request.FormValue("ProjectType")
		projectType, _:= strconv.Atoi(pt)
		project.ProjectType = projectType
	}
	}
	if err := models.UpdateTProjectsById(project); err != nil {
		beego.Error(err)
		msg := project.ProjectName + "失败！"
		tools.AddLog(user.UserRealName, tools.Edit, tools.Project, msg)
		fmt.Fprint(this.Ctx.ResponseWriter, "修改失败，请重试！")
	}else {
		msg := project.ProjectName + "成功！"
		tools.AddLog(user.UserRealName, tools.Edit, tools.Project, msg)
		fmt.Fprint(this.Ctx.ResponseWriter, "修改成功！")
	}
}