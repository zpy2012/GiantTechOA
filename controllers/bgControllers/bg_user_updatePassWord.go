package bgControllers

import (
	"github.com/astaxie/beego"
	"fmt"
	"GiantTech/models"
	"GiantTech/controllers/tools"
)

type BgUserUpdatePassWordController struct {
	beego.Controller
}

func (this *BgUserUpdatePassWordController) Prepare() {
	s := this.StartSession()
	username = s.Get("login")
	beego.Informational(username)
	if username == nil {
		this.Ctx.Redirect(302, "/login")
	}
}

func (this *BgUserUpdatePassWordController) Post() {
	user, _ := models.GetTUsersByName(username.(string))
	password := this.Ctx.Request.FormValue("passWord")
	if user.UserStatus == 1 {
		user.UserPassword = tools.EncryptionPassWord(password, "zpy")
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