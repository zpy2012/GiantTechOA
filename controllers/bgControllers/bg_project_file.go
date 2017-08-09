package bgControllers

import (
	"github.com/astaxie/beego"
	"strconv"
	"GiantTech/models"
	"GiantTech/controllers/tools"
)

type BgProjectFileController struct {
	beego.Controller
}

func (this *BgProjectFileController) Prepare() {
	s := this.StartSession()
	username = s.Get("login")
	beego.Informational(username)
	if username == nil {
		this.Ctx.Redirect(302, "/login")
	}
}

func (this *BgProjectFileController) Get() {
	s := this.StartSession()
	p := this.Ctx.Request.FormValue("page")
	page, _ := strconv.Atoi(p)
	id := this.Ctx.Request.FormValue("id")
	s.Set("ProjectId", id)
	projectId, _ := strconv.Atoi(id)
	project, _ := models.GetTProjectsById(projectId)
	this.Data["ProjectName"] = project.ProjectName
	offset := (page-1)*prepage
	user, _ := models.GetTUsersByName(username.(string))
	this.Data["User"] = user
	query := make(map[string]string)
	query["FileProjectID"] = id
	var  order, sortBy []string
	order = append(order, "desc")
	sortBy = append(sortBy, "FileCreatedTime")
	i, _ := models.GetAllTProjectFile(query, nil, sortBy, order, 0, 0)
	if files, err := models.GetAllTProjectFile(query, nil, sortBy, order, int64(offset), int64(prepage)); err == nil {
		res := tools.Paginator(page, prepage, int64(len(i)))
		this.Data["paginator"] = res
		this.Data["files"] = files
		this.Data["ProjectId"] = id
	}else {
		res := tools.Paginator(page, prepage, 0)
		this.Data["paginator"] = res
		beego.Error(err)
	}
	this.TplName = "bgview/projectfile.html"
}