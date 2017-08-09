package bgControllers

import (
	"github.com/astaxie/beego"
	"GiantTech/models"
	"strconv"
	"GiantTech/controllers/tools"
)

type BgUserAddController struct {
	beego.Controller
}

func (this *BgUserAddController) Prepare() {
	s := this.StartSession()
	username = s.Get("login")
	beego.Informational(username)
	if username == nil {
		this.Ctx.Redirect(302, "/login")
	}
}

func (this *BgUserAddController) Get() {
	user, _ := models.GetTUsersByName(username.(string))
	this.Data["User"] = user
	departments, _ := models.GetAllTDepartment(nil,nil,nil,nil,0, 0)
	this.Data["departments"] = departments
	this.TplName = "bgview/useradd.html"
}

func (this *BgUserAddController) Post() {
	name := this.Ctx.Request.FormValue("UserName")
	realName := this.Ctx.Request.FormValue("RealName")
	passWd := this.Ctx.Request.FormValue("PassWord")
	l := this.Ctx.Request.FormValue("level")
	level, _ := strconv.Atoi(l)
	a := this.Ctx.Request.FormValue("allProjects")
	allProjects, _ := strconv.Atoi(a)
	p := this.Ctx.Request.FormValue("price")
	price, _ := strconv.Atoi(p)
	d := this.Ctx.Request.FormValue("departmentID")
	departmentID, _ := strconv.Atoi(d)
	loginUser, _ := models.GetTUsersByName(username.(string))
	defaultHeadImage := beego.AppConfig.String("defaultheadimage")
	var user models.TUsers
	user.UserName = name
	user.UserRealName = realName
	user.UserPassword = tools.EncryptionPassWord(passWd, "zpy")
	user.UserLevel = level
	user.UserStatus = 0
	user.UserAllProjects = allProjects
	user.UserPrice = price
	user.UserDepartmentID = departmentID
	user.UserCreatedTime = tools.TimeNow()
	user.UserHeadImagePath = defaultHeadImage
	if _, err := models.AddTUsers(&user); err != nil {
		beego.Error(err)
		msg := user.UserName + "失败！"
		tools.AddLog(loginUser.UserRealName, tools.Addsome, tools.Person, msg)
		var errResult models.ErrorResult
		errResult.Msg = "添加失败，请重试！"
		errResult.Code = 201
		this.Data["json"] = errResult
		this.ServeJSON()
	}else {
		msg := user.UserName + "成功！"
		tools.AddLog(loginUser.UserRealName, tools.Addsome, tools.Person, msg)
		var Result models.Result
		Result.Msg = "添加成功！"
		Result.Code = 200
		Result.JumpUrl = "/user/?page=1"
		this.Data["json"] = Result
		this.ServeJSON()
	}
}