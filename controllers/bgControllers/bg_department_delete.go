package bgControllers

import (
	"github.com/astaxie/beego"
	"strconv"
	"GiantTech/models"
	"GiantTech/controllers/tools"
)

type BgDepartmentDeleteController struct {
	beego.Controller
}

func (this *BgDepartmentDeleteController) Prepare() {
	s := this.StartSession()
	username = s.Get("login")
	beego.Informational(username)
	if username == nil {
		this.Ctx.Redirect(302, "/login")
	}
}

func (this *BgDepartmentDeleteController) Get() {
	p := this.Ctx.Request.FormValue("id")
	id, _ := strconv.Atoi(p)
	beego.Informational(id)
	user, _ := models.GetTUsersByName(username.(string))
	this.Data["User"] = user
	department,_ := models.GetTDepartmentById(id)
	if err := models.DeleteTDepartment(id); err != nil {
		beego.Error(err)
		msg := department.DepartmentName + "失败！"
		tools.AddLog(user.UserRealName, tools.Delete, tools.Department, msg)
	}else {
		msg := department.DepartmentName + "成功！"
		tools.AddLog(user.UserRealName, tools.Delete, tools.Department, msg)
	}
	this.Ctx.Redirect(302, "/department/?page=1")
}