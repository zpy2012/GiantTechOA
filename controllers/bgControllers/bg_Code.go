package bgControllers

import (
	"github.com/astaxie/beego"
	"fmt"
	"github.com/ying32/alidayu"
	"GiantTech/models"
	"GiantTech/controllers/tools"
	"time"
	"strings"
)

type BgCodeController struct {
	beego.Controller
}

var appKey = beego.AppConfig.String("appkey")
var appSecret = beego.AppConfig.String("appsecret")
var sign = beego.AppConfig.String("sign")

func (this *BgCodeController) Get() {
	phoneNumber := this.Ctx.Request.FormValue("phoneNumber")
	codeModel, _ := models.GetTCodeByPhoneNumber(phoneNumber)
	if codeModel != nil {
		models.DeleteTCode(codeModel.Id)
	}
	code := tools.RandCode()
	param := fmt.Sprintf(`{"code":%s}`,code)
	sms_id := beego.AppConfig.String("code_sms_id")
	success, resp, _ := alidayu.SendSMS(phoneNumber, sign, sms_id, param, appKey, appSecret)
	beego.Informational(resp)
	if success {
		t := tools.TimeNow()
		var m models.TCode
		m.CodePhoneNumber = phoneNumber
		m.Code = code
		m.CodeCreatedTime = t
		m.CodeEndTime = t.Add(time.Minute*10)
		if _, err := models.AddTCode(&m); err != nil {
			beego.Error(err)
		}
		fmt.Fprint(this.Ctx.ResponseWriter, "获取成功")
	}else {
		fmt.Fprint(this.Ctx.ResponseWriter, "获取失败")
	}
}

func (this *BgCodeController) Post() {
	phoneNumber := this.Ctx.Request.FormValue("phoneNumber")
	code := this.Ctx.Request.FormValue("code")
	userName := this.Ctx.Request.FormValue("userName")
	codeType:= this.Ctx.Request.FormValue("type")
	codeModel, err := models.GetTCodeByPhoneNumber(phoneNumber)
	if err !=nil {
		beego.Warning(err)
		fmt.Fprint(this.Ctx.ResponseWriter, "请先获取验证码！")
		return
	}
	t := tools.TimeNow()
	if t.After(codeModel.CodeEndTime) {
		err := models.DeleteTCode(codeModel.Id)
		if err == nil {
			fmt.Fprint(this.Ctx.ResponseWriter, "验证码过期，请重新获取！")
		}else {
			beego.Error(err)
		}
	}else {
		if strings.EqualFold(codeModel.Code, code) {
			switch codeType {
			case "new":
				{
					user, _ := models.GetTUsersByName(userName)
					user.UserPhoneNumber = phoneNumber
					models.UpdateTUsersById(user)
					fmt.Fprint(this.Ctx.ResponseWriter, "验证成功")
				}
			case "checkPhone":
				{
					user, _ := models.GetTUsersByName(username.(string))
					user.UserStatus = 1
					models.UpdateTUsersById(user)
					fmt.Fprint(this.Ctx.ResponseWriter, "验证成功")
				}
			case "updatePhone":
				{
					user, _ := models.GetTUsersByName(username.(string))
					if user.UserStatus == 1 {
						user.UserPhoneNumber = phoneNumber
						user.UserStatus = 0
						if err := models.UpdateTUsersById(user);err != nil {
							beego.Error(err)
							fmt.Fprint(this.Ctx.ResponseWriter, "更新失败")
						}else {
							fmt.Fprint(this.Ctx.ResponseWriter, "修改成功")
						}

					}else {
						fmt.Fprint(this.Ctx.ResponseWriter, "更新失败,请先验证手机号码！")
					}
				}
				
			}
		}else {
			fmt.Fprint(this.Ctx.ResponseWriter, "验证码错误！")
		}
	}
}