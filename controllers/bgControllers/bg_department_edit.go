package bgControllers

import (
	"github.com/astaxie/beego"
	"strconv"
	"GiantTech/models"
	"GiantTech/controllers/tools"
)

type BgDepartmentEditController struct {
	beego.Controller
}

func (this *BgDepartmentEditController) Prepare() {
	s := this.StartSession()
	username = s.Get("login")
	beego.Informational(username)
	if username == nil {
		this.Ctx.Redirect(302, "/login")
	}
}

func (this *BgDepartmentEditController) Get() {
	p := this.Ctx.Request.FormValue("id")
	id, _ := strconv.Atoi(p)
	user, _ := models.GetTUsersByName(username.(string))
	this.Data["User"] = user
	department, _ := models.GetTDepartmentById(id)
	this.Data["id"] = department.Id
	this.Data["name"] = department.DepartmentName
	this.TplName = "bgview/departmentedit.html"
}

func (this *BgDepartmentEditController) Post() {
	user, _ := models.GetTUsersByName(username.(string))
	p := this.Ctx.Request.FormValue("id")
	beego.Informational(p)
	id, _ := strconv.Atoi(p)
	name := this.Ctx.Request.FormValue("departmentName")
	department, _ := models.GetTDepartmentById(id)
	department.Id = id
	department.DepartmentName = name
	if err := models.UpdateTDepartmentById(department); err != nil {
		msg := department.DepartmentName + "失败！"
		tools.AddLog(user.UserRealName, tools.Edit, tools.Department, msg)
		var errResult models.ErrorResult
		errResult.Msg = "更新失败，请重试！"
		errResult.Code = 201
		this.Data["json"] = errResult
		this.ServeJSON()
	}else {
		msg := department.DepartmentName + "失败！"
		tools.AddLog(user.UserRealName, tools.Edit, tools.Department, msg)
		var Result models.Result
		Result.Msg = "更新成功！"
		Result.Code = 200
		Result.JumpUrl = "/department/?page=1"
		this.Data["json"] = Result
		this.ServeJSON()
	}
}