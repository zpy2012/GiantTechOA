package bgControllers

import (
	"github.com/astaxie/beego"
	"strconv"
	"GiantTech/models"
	"GiantTech/controllers/tools"
)

type BgProjectFileDeleteController struct {
	beego.Controller
}

func (this *BgProjectFileDeleteController) Prepare() {
	s := this.StartSession()
	username = s.Get("login")
	beego.Informational(username)
	if username == nil {
		this.Ctx.Redirect(302, "/login")
	}
}

func (this *BgProjectFileDeleteController) Get() {
	s := this.StartSession()
	projectId := s.Get("ProjectId")
	user, _ := models.GetTUsersByName(username.(string))
	id := this.Ctx.Request.FormValue("id")
	fileId, _ := strconv.Atoi(id)
	file, _ := models.GetTProjectFileById(fileId)
	if err := tools.DeleteFileWithPath(file.FilePath); err == nil {
		if err := models.DeleteTProjectFile(fileId); err == nil {
			tools.AddLog(user.UserRealName, tools.Delete, tools.File, file.FileName + "成功")
		}else {
			beego.Error(err)
		}
	}else {
		beego.Error(err)
	}
	this.Ctx.Redirect(302, "/project/file/?page=1&id="+projectId.(string))
}