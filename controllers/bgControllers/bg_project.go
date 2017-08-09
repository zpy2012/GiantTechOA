package bgControllers

import (
	"github.com/astaxie/beego"
	"strconv"
	"GiantTech/models"
	"GiantTech/controllers/tools"
	"strings"
)

type BgProjectController struct {
	beego.Controller
}

func (this *BgProjectController) Prepare() {
	s := this.StartSession()
	username = s.Get("login")
	beego.Informational(username)
	if username == nil {
		this.Ctx.Redirect(302, "/login")
	}
}

func (this *BgProjectController) Get() {
	p := this.Ctx.Request.FormValue("page")
	page, _ := strconv.Atoi(p)
	searchName := this.Ctx.Request.FormValue("projectname")
	this.Data["searchName"] = searchName
	offset := (page-1)*prepage
	user, _ := models.GetTUsersByName(username.(string))
	this.Data["User"] = user
	var query = make(map[string]string)
	if user.UserAllProjects == 0 {
		query["ProjectWonerID"] = strconv.FormatInt(int64(user.Id), 10)
	}else {
		delete(query, "ProjectWonerID")
	}
	if strings.EqualFold(searchName, "all") {
		delete(query, "ProjectName.contains")
	}else {
		query["ProjectName.contains"] = searchName
	}
	var order, sortBy  []string
	order = append(order, "desc")
	sortBy = append(sortBy, "ProjectCreatedTime")
	i, _ := models.GetAllTProjects(query, nil, sortBy, order, 0, 0)
	if projects, err := models.GetAllTProjects(query, nil, sortBy, order, int64(offset), int64(prepage)); err == nil {
		res := tools.Paginator(page, prepage, int64(len(i)))
		this.Data["paginator"] = res
		this.Data["projects"] = projects
	}else {
		res := tools.Paginator(page, prepage, 0)
		this.Data["paginator"] = res
		beego.Error(err)
	}
	this.TplName = "bgview/project.html"
}
