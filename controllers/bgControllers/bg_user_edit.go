package bgControllers

import (
	"github.com/astaxie/beego"
	"strconv"
	"GiantTech/models"
	"GiantTech/controllers/tools"
)

type BgUserEditController struct {
	beego.Controller
}

func (this *BgUserEditController) Prepare() {
	s := this.StartSession()
	username = s.Get("login")
	beego.Informational(username)
	if username == nil {
		this.Ctx.Redirect(302, "/login")
	}
}

func (this *BgUserEditController) Get() {
	p := this.Ctx.Request.FormValue("id")
	id, _ := strconv.Atoi(p)
	user, _ := models.GetTUsersByName(username.(string))
	this.Data["User"] = user
	u, _ := models.GetTUsersById(id)
	departments, _ := models.GetAllTDepartment(nil,nil,nil,nil,0, 0)
	this.Data["departments"] = departments
	this.Data["id"] = u.Id
	this.Data["allProjects"] = u.UserAllProjects
	this.Data["price"] = u.UserPrice
	this.Data["departmentID"] = u.UserDepartmentID
	this.TplName = "bgview/useredit.html"
}

func (this *BgUserEditController) Post() {
	i := this.Ctx.Request.FormValue("id")
	id, _ := strconv.Atoi(i)
	a := this.Ctx.Request.FormValue("allProjects")
	allProjects, _ := strconv.Atoi(a)
	p := this.Ctx.Request.FormValue("price")
	price, _ := strconv.Atoi(p)
	d := this.Ctx.Request.FormValue("departmentID")
	departmentID, _ := strconv.Atoi(d)
	beego.Informational(id,allProjects,price,departmentID)
	loginUser, _ := models.GetTUsersByName(username.(string))
	user, _ := models.GetTUsersById(id)
	user.UserAllProjects = allProjects
	user.UserPrice = price
	user.UserDepartmentID = departmentID
	if err := models.UpdateTUsersById(user); err != nil {
		msg := user.UserName + "失败！"
		tools.AddLog(loginUser.UserRealName, tools.Edit, tools.Person, msg)
		var errResult models.ErrorResult
		errResult.Msg = "更新失败，请重试！"
		errResult.Code = 201
		this.Data["json"] = errResult
		this.ServeJSON()
	}else {
		msg := user.UserName + "成功！"
		tools.AddLog(loginUser.UserRealName, tools.Edit, tools.Person, msg)
		var Result models.Result
		Result.Msg = "更新成功！"
		Result.Code = 200
		Result.JumpUrl = "/user/?page=1"
		this.Data["json"] = Result
		this.ServeJSON()
	}
}