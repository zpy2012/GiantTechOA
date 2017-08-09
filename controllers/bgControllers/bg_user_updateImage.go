package bgControllers

import (
	"github.com/astaxie/beego"
	"GiantTech/models"
	"strings"
	"GiantTech/controllers/tools"
	"fmt"
)

type BgUserUpdateImageController struct {
	beego.Controller
}

func (this *BgUserUpdateImageController) Prepare() {
	s := this.StartSession()
	username = s.Get("login")
	beego.Informational(username)
	if username == nil {
		this.Ctx.Redirect(302, "/login")
	}
}

func (this *BgUserUpdateImageController) Post() {
	user, _ := models.GetTUsersByName(username.(string))
	this.Data["User"] = user
	imagePath := beego.AppConfig.String("userimagepath")
	defaultHeadImage := beego.AppConfig.String("defaultheadimage")
	if !strings.EqualFold(user.UserHeadImagePath, defaultHeadImage) {
		tools.DeleteImageWithPath(user.UserHeadImagePath)
	}
	file, handler, err := this.Ctx.Request.FormFile("userFile")
	path, err := tools.SaveUserHeadImage(file, handler, int64(user.Id), imagePath)
	if err != nil {
		beego.Error(err)
		fmt.Fprint(this.Ctx.ResponseWriter, "失败，请重试！")
	}else {
		user.UserHeadImagePath = path
		err := models.UpdateTUsersById(user)
		if err != nil {
			beego.Error(err)
			fmt.Fprint(this.Ctx.ResponseWriter, "失败，请重试！")
		}else {
			fmt.Fprint(this.Ctx.ResponseWriter, "成功")
		}
	}
}