package bgControllers

import (
	"github.com/astaxie/beego"
	"GiantTech/models"
	"strconv"
	"GiantTech/controllers/tools"
	"encoding/json"
	"github.com/ying32/alidayu"
)

type BgProjectAddController struct {
	beego.Controller
}

func (this *BgProjectAddController) Prepare() {
	s := this.StartSession()
	username = s.Get("login")
	beego.Informational(username)
	if username == nil {
		this.Ctx.Redirect(302, "/login")
	}
}

func (this *BgProjectAddController) Get() {
	user, _ := models.GetTUsersByName(username.(string))
	this.Data["User"] = user
	this.TplName = "bgview/projectadd.html"
}

func (this *BgProjectAddController) Post() {
	name := this.Ctx.Request.FormValue("projectName")
	description := this.Ctx.Request.FormValue("projectDescription")
	s := this.Ctx.Request.FormValue("projectSource")
	source, _ := strconv.Atoi(s)
	expectedAmount := this.Ctx.Request.FormValue("projectExpectedAmount")
	expectedDuration := this.Ctx.Request.FormValue("projectExpectedDuration")
	dealAmount := this.Ctx.Request.FormValue("projectDealAmount")
	dealDuration := this.Ctx.Request.FormValue("projectDealDuration")
	ps := this.Ctx.Request.FormValue("projectStatus")
	status, _ := strconv.Atoi(ps)
	pt := this.Ctx.Request.FormValue("projectType")
	projectType, _:= strconv.Atoi(pt)
	user, _ := models.GetTUsersByName(username.(string))
	var project models.TProjects
	project.ProjectName = name
	project.ProjectDescription = description
	project.ProjectSource = source
	project.ProjectExpectedAmount = expectedAmount
	project.ProjectExpectedDuration = expectedDuration
	project.ProjectDealAmount = dealAmount
	project.ProjectDealDuration = dealDuration
	project.ProjectStatus = status
	project.ProjectType = projectType
	project.ProjectFile = 0
	project.ProjectWonerID = user.Id
	project.ProjectCreatedTime = tools.TimeNow()
	if _, err := models.AddTProjects(&project); err != nil {
		beego.Error(err)
		msg := project.ProjectName + "失败！"
		tools.AddLog(user.UserRealName, tools.Addsome, tools.Project, msg)
		var errResult models.ErrorResult
		errResult.Msg = "添加失败，请重试！"
		errResult.Code = 201
		this.Data["json"] = errResult
		this.ServeJSON()
	}else {
		query := make(map[string]string)
		query["UserAllProjects"] = "1"
		users,_ := models.GetAllTUsers(query, nil, nil, nil, 0, 0)
		var notice models.NewProjectNotice
		notice.Name = user.UserRealName
		notice.ProjectName = project.ProjectName
		p, _ := json.Marshal(notice)
		param := string(p[:])
		sms_id := beego.AppConfig.String("newproject_sms_id")
		for _, v := range users{
			_, resp, _ := alidayu.SendSMS(v.(models.TUsers).UserPhoneNumber, sign, sms_id, param, appKey, appSecret)
			beego.Informational(resp)
		}
		msg := project.ProjectName + "成功！"
		tools.AddLog(user.UserRealName, tools.Addsome, tools.Project, msg)
		var Result models.Result
		Result.Msg = "添加成功！"
		Result.Code = 200
		Result.JumpUrl = "/project/?page=1"
		this.Data["json"] = Result
		this.ServeJSON()
	}
}