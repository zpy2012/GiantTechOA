package bgControllers

import (
	"github.com/astaxie/beego"
	"GiantTech/controllers/tools"
	"fmt"
	"GiantTech/models"
)

type BgLogoutController struct {
	beego.Controller
}

func (c *BgLogoutController) Get() {
	s := c.StartSession()
	username = s.Get("login")
	loginUser, err := models.GetTUsersByName(username.(string))
	if err != nil {
		beego.Error(err)
		fmt.Fprint(c.Ctx.ResponseWriter, err)
	}
	s.Delete("login")
	tools.AddLog(loginUser.UserRealName, tools.Logout, tools.System, "成功！")
	c.Ctx.Redirect(302, "/login")
}
