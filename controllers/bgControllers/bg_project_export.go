package bgControllers

import (
	"github.com/astaxie/beego"
	"GiantTech/models"
	"fmt"
	"GiantTech/controllers/tools"
)

type BgProjectExportController struct {
	beego.Controller
}

func (this *BgProjectExportController) Prepare() {
	s := this.StartSession()
	username = s.Get("login")
	beego.Informational(username)
	if username == nil {
		this.Ctx.Redirect(302, "/login")
	}
}

func (this *BgProjectExportController) Get() {
	user, _ := models.GetTUsersByName(username.(string))
	if user.UserAllProjects == 1 {
		path, err := tools.ExportProjectsExcel()
		if err == nil {
			fmt.Fprint(this.Ctx.ResponseWriter, path)
		}else {
			beego.Error(err)
			fmt.Fprint(this.Ctx.ResponseWriter, "失败")
		}
	}else {
		fmt.Fprint(this.Ctx.ResponseWriter, "失败")
	}
}