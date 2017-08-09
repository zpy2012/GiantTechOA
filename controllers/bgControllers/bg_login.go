package bgControllers

import (
	"github.com/astaxie/beego"
	"fmt"
	"strings"
	"GiantTech/models"
	"GiantTech/controllers/tools"
)

type BgLoginController struct {
	beego.Controller
}

var prepage, _ = beego.AppConfig.Int("prepage")
var username interface{}

func (c *BgLoginController) Get() {
	SetUP()
	c.TplName = "bgview/login.html"
}

func (c *BgLoginController) Post ()  {
	c.ParseForm(c.Ctx.Input.RequestBody)
	username := c.Ctx.Request.FormValue("username")
	password := c.Ctx.Request.FormValue("password")
	password = tools.EncryptionPassWord(password, "zpy")
	loginUser, err := models.GetTUsersByName(username)
	if err != nil {
		beego.Error(err)
		fmt.Fprint(c.Ctx.ResponseWriter, err)
	}
	if loginUser == nil {
		tools.AddLog(loginUser.UserRealName, tools.Login, tools.System, "失败：用户不存在！")
		fmt.Fprint(c.Ctx.ResponseWriter, "用户不存在")
	}else if !strings.EqualFold(loginUser.UserPassword, password) {
		tools.AddLog(loginUser.UserRealName, tools.Login, tools.System, "失败：密码错误！")
		fmt.Fprint(c.Ctx.ResponseWriter, "密码错误")
	}else if strings.EqualFold(loginUser.UserPhoneNumber, "") {
		tools.AddLog(loginUser.UserRealName, tools.Login, tools.System, "失败：需要绑定手机号！")
		fmt.Fprint(c.Ctx.ResponseWriter, "绑定手机")
	}else {
		s := c.StartSession()
		s.Set("login", loginUser.UserName)
		tools.AddLog(loginUser.UserRealName, tools.Login, tools.System, "登陆成功！")
		//c.Ctx.Redirect(302, "/admin")
		fmt.Fprint(c.Ctx.ResponseWriter, "登陆成功")
	}
}