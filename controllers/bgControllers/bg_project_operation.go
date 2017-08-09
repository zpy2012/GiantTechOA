package bgControllers

import (
	"github.com/astaxie/beego"
	"strconv"
	"fmt"
	"GiantTech/models"
	"github.com/ying32/alidayu"
	"GiantTech/controllers/tools"
	"encoding/json"
)

type BgProjectOperationController struct {
	beego.Controller
}

func (this *BgProjectOperationController) Prepare() {
	s := this.StartSession()
	username = s.Get("login")
	beego.Informational(username)
	if username == nil {
		this.Ctx.Redirect(302, "/login")
	}
}

func (this *BgProjectOperationController) Get() {
	user, _ := models.GetTUsersByName(username.(string))
	this.Data["User"] = user
	p := this.Ctx.Request.FormValue("ProjectId")
	id, _ := strconv.Atoi(p)
	beego.Informational(id)
	project, _ := models.GetTProjectsById(id)
	if user.UserPrice == -1 {
		var out string
		switch project.ProjectStatus {
		case 0: out = "提醒报价"
		case 1: out = "签约"
		case 6: out = "提醒结项"
		}
		fmt.Fprint(this.Ctx.ResponseWriter, out)
	}else {
		var out string
		switch project.ProjectStatus {
		case 0: out = "报价"
		case 2: out = "开始项目"
		case 6: out = "结项"
		default: out = "变更状态"
		}
		fmt.Fprint(this.Ctx.ResponseWriter, out)
	}
}

func (this *BgProjectOperationController) Post() {
	user, _ := models.GetTUsersByName(username.(string))
	this.Data["User"] = user
	p := this.Ctx.Request.FormValue("ProjectId")
	id, _ := strconv.Atoi(p)
	project, _ := models.GetTProjectsById(id)
	optype := this.Ctx.Request.FormValue("type")
	beego.Informational(optype)
	switch optype {
	case "提醒报价":{
		query := make(map[string]string)
		switch project.ProjectType {
		case 0:{
			query["UserPrice"] = "0"
		}
		case 1:{
			query["UserPrice"] = "1"
		}
		case 2:{
			query["UserPrice"] = "2"
		}
		}
		users, _ := models.GetAllTUsers(query, nil ,nil, nil, 0 ,0)
		if len(users) >= 1 {
			projectOwner, _ := models.GetTUsersById(project.ProjectWonerID)
			phoneNumber := users[0].(models.TUsers).UserPhoneNumber
			var notice models.NoticeModel
			notice.Name = users[0].(models.TUsers).UserRealName
			notice.OwnerName = projectOwner.UserRealName
			notice.ProjectName = project.ProjectName
			p, _ := json.Marshal(notice)
			param := string(p[:])
			sms_id := beego.AppConfig.String("price_sms_id")
			success, resp, err := alidayu.SendSMS(phoneNumber, sign, sms_id, param, appKey, appSecret)
			beego.Informational(resp)
			if success {
				fmt.Fprint(this.Ctx.ResponseWriter, "提醒成功")
			}else {
				beego.Error(err)
				fmt.Fprint(this.Ctx.ResponseWriter, "提醒失败，请稍后重试")
			}
		}else {
			fmt.Fprint(this.Ctx.ResponseWriter, "尚未设置该分类的报价人员，请联系管理员")
		}
	}
	case "签约":{
		ProjectPrice := this.Ctx.Request.FormValue("price")
		ProjectTime := this.Ctx.Request.FormValue("time")
		project.ProjectDealAmount = ProjectPrice
		project.ProjectDealDuration = ProjectTime
		project.ProjectStatus = 2
		if err := models.UpdateTProjectsById(project); err == nil {
			fmt.Fprint(this.Ctx.ResponseWriter, "操作成功")
		}else {
			beego.Error(err)
			fmt.Fprint(this.Ctx.ResponseWriter, "操作失败，请稍后重试！")
		}
	}
	case "提醒结项":{
		query := make(map[string]string)
		switch project.ProjectType {
		case 0:{
			query["UserPrice"] = "0"
		}
		case 1:{
			query["UserPrice"] = "1"
		}
		case 2:{
			query["UserPrice"] = "2"
		}
		}
		users, _ := models.GetAllTUsers(query, nil ,nil, nil, 0 ,0)
		projectOwner, _ := models.GetTUsersById(project.ProjectWonerID)
		phoneNumber := users[0].(models.TUsers).UserPhoneNumber
		var notice models.NoticeModel
		notice.Name = users[0].(models.TUsers).UserRealName
		notice.OwnerName = projectOwner.UserRealName
		notice.ProjectName = project.ProjectName
		p, _ := json.Marshal(notice)
		param := string(p[:])
		sms_id := beego.AppConfig.String("end_sms_id")
		success, resp, err := alidayu.SendSMS(phoneNumber, sign, sms_id, param, appKey, appSecret)
		beego.Informational(resp)
		if success {
			fmt.Fprint(this.Ctx.ResponseWriter, "提醒成功")
		}else {
			beego.Error(err)
			fmt.Fprint(this.Ctx.ResponseWriter, "提醒失败，请稍后重试")
		}
	}
	case "报价":{
		ProjectPrice := this.Ctx.Request.FormValue("price")
		ProjectTime := this.Ctx.Request.FormValue("time")
		beego.Informational(ProjectPrice,ProjectTime)
		project.ProjectExpectedAmount = ProjectPrice
		project.ProjectExpectedDuration = ProjectTime
		project.ProjectStatus = 1
		if err := models.UpdateTProjectsById(project); err == nil {
			owner,_ := models.GetTUsersById(project.ProjectWonerID)
			phoneNumber := owner.UserPhoneNumber
			var notice models.PriceNoticeModel
			notice.Name = user.UserRealName
			notice.ProjectName = project.ProjectName
			p, _ := json.Marshal(notice)
			param := string(p[:])
			sms_id := beego.AppConfig.String("pricenotice_sms_id")
			_, resp, _ := alidayu.SendSMS(phoneNumber, sign, sms_id, param, appKey, appSecret)
			beego.Informational(resp)
			fmt.Fprint(this.Ctx.ResponseWriter, "操作成功")
		}else {
			beego.Error(err)
			fmt.Fprint(this.Ctx.ResponseWriter, "操作失败，请稍后重试！")
		}
	}
	case "开始项目":{
		project.ProjectStartTime = tools.TimeNow()
		project.ProjectStatus = 3
		if err := models.UpdateTProjectsById(project); err == nil {
			fmt.Fprint(this.Ctx.ResponseWriter, "操作成功")
		}else {
			beego.Error(err)
			fmt.Fprint(this.Ctx.ResponseWriter, "操作失败，请稍后重试！")
		}
	}
	case "结项":{
		project.ProjectEndTime = tools.TimeNow()
		project.ProjectStatus = 7
		if err := models.UpdateTProjectsById(project); err == nil {
			fmt.Fprint(this.Ctx.ResponseWriter, "操作成功")
		}else {
			beego.Error(err)
			fmt.Fprint(this.Ctx.ResponseWriter, "操作失败，请稍后重试！")
		}
	}
	case "变更状态":{
		ps := this.Ctx.Request.FormValue("status")
		ProjectStatus, _ := strconv.Atoi(ps)
		if ProjectStatus == 6 {
			project.ProjectCompleteTime = tools.TimeNow()
		}
		project.ProjectStatus = ProjectStatus
		if err := models.UpdateTProjectsById(project); err == nil {
			fmt.Fprint(this.Ctx.ResponseWriter, "操作成功")
		}else {
			beego.Error(err)
			fmt.Fprint(this.Ctx.ResponseWriter, "操作失败，请稍后重试！")
		}
	}
	}
}