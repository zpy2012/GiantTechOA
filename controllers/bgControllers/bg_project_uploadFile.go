package bgControllers

import (
	"github.com/astaxie/beego"
	"GiantTech/models"
	"fmt"
	"GiantTech/controllers/tools"
	"strconv"
	"strings"
)

type BgProjectUploadFileController struct {
	beego.Controller
}

func (this *BgProjectUploadFileController) Prepare() {
	s := this.StartSession()
	username = s.Get("login")
	beego.Informational(username)
	if username == nil {
		this.Ctx.Redirect(302, "/login")
	}
}

func (this *BgProjectUploadFileController) Post() {
	s := this.StartSession()
	projectId := s.Get("ProjectId")
	id, _ := strconv.Atoi(projectId.(string))
	project, _ := models.GetTProjectsById(id)
	user, _ := models.GetTUsersByName(username.(string))
	path := beego.AppConfig.String("projectfilepath") + project.ProjectName
	if file, handler, err := this.Ctx.Request.FormFile("file"); err != nil {
		beego.Error(err)
		fmt.Fprint(this.Ctx.ResponseWriter, err)
	}else {
		n := strings.Split(handler.Filename, ".")
		var name string
		for i := 0; i < len(n)-1; i++ {
			name = name + n[i]
		}
		query := make(map[string]string)
		query["FileProjectID"] = projectId.(string)
		query["FileName.contains"] = name
		files, _ := models.GetAllTProjectFile(query, nil, nil, nil, 0 ,0)
		fileName, filePath, e := tools.SaveFile(file, handler, int64(len(files)), path)
		if e != nil {
			beego.Error(err)
			fmt.Fprint(this.Ctx.ResponseWriter, err)
		}else {
			var fileModel models.TProjectFile
			fileModel.FileName = fileName
			fileModel.FilePath = filePath
			fileModel.FileOwner = user.UserRealName
			fileModel.FileProjectID = id
			fileModel.FileCreatedTime = tools.TimeNow()
			i, _ := models.AddTProjectFile(&fileModel)
			tools.AddLog(user.UserRealName, tools.Addsome, tools.File, fileName + "成功")
			beego.Informational(fileName, "上传成功", filePath)
			var uploadResult models.UploadResult
			uploadResult.Status = true
			uploadResult.Data.Id = strconv.FormatInt(i, 10)
			uploadResult.Message = "操作成功"
			this.Data["json"] = uploadResult
			this.ServeJSON()
		}
	}
}

func (this *BgProjectUploadFileController) Get() {
	s := this.StartSession()
	user, _ := models.GetTUsersByName(username.(string))
	this.Data["User"] = user
	i := s.Get("ProjectId")
	d := i.(string)
	id, _ := strconv.Atoi(d)
	project, _ := models.GetTProjectsById(id)
	this.Data["ProjectName"] = project.ProjectName
	this.Data["ProjectId"] = id
	this.TplName = "bgview/projectfileupload.html"
}