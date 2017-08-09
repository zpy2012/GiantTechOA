package bgControllers

import (
	"github.com/astaxie/beego"
	"GiantTech/models"
	"GiantTech/controllers/tools"
)

type BgDepartmentAddController struct {
	beego.Controller
}

func (this *BgDepartmentAddController) Prepare() {
	s := this.StartSession()
	username = s.Get("login")
	beego.Informational(username)
	if username == nil {
		this.Ctx.Redirect(302, "/login")
	}
}

func (this *BgDepartmentAddController) Get() {
	user, _ := models.GetTUsersByName(username.(string))
	this.Data["User"] = user
	this.TplName = "bgview/departmentadd.html"
}

func (this *BgDepartmentAddController) Post() {
	user, _ := models.GetTUsersByName(username.(string))
	name := this.Ctx.Request.FormValue("departmentName")
	var department models.TDepartment
	department.DepartmentName = name
	department.DepartmentParentID = -1
	department.DepartmentCreatedTime = tools.TimeNow()
	if _, err := models.AddTDepartment(&department); err != nil {
		beego.Error(err)
		msg := department.DepartmentName + "失败！"
		tools.AddLog(user.UserRealName, tools.Addsome, tools.Department, msg)
		var errResult models.ErrorResult
		errResult.Msg = "添加失败，请重试！"
		errResult.Code = 201
		this.Data["json"] = errResult
		this.ServeJSON()
	}else {
		msg := department.DepartmentName + "成功！"
		tools.AddLog(user.UserRealName, tools.Addsome, tools.Department, msg)
		var Result models.Result
		Result.Msg = "添加成功！"
		Result.Code = 200
		Result.JumpUrl = "/department/?page=1"
		this.Data["json"] = Result
		this.ServeJSON()
	}
}