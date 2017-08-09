package bgControllers

import (
	"github.com/astaxie/beego"
	"strconv"
	"GiantTech/models"
	"GiantTech/controllers/tools"
)

type BgDepartmentController struct {
	beego.Controller
}

func (this *BgDepartmentController) Prepare() {
	s := this.StartSession()
	username = s.Get("login")
	beego.Informational(username)
	if username == nil {
		this.Ctx.Redirect(302, "/login")
	}
}

func (this *BgDepartmentController) Get() {
	p := this.Ctx.Request.FormValue("page")
	page, _ := strconv.Atoi(p)
	offset := (page-1)*prepage
	user, _ := models.GetTUsersByName(username.(string))
	this.Data["User"] = user
	var order, sortBy  []string
	order = append(order, "desc")
	sortBy = append(sortBy, "DepartmentCreatedTime")
	i, _ := models.GetAllTDepartment(nil, nil, sortBy, order, 0, 0)
	if departments, err := models.GetAllTDepartment(nil, nil, sortBy, order, int64(offset), int64(prepage)); err == nil {
		res := tools.Paginator(page, prepage, int64(len(i)))
		this.Data["paginator"] = res
		this.Data["departments"] = departments
	}else {
		res := tools.Paginator(page, prepage, 0)
		this.Data["paginator"] = res
		beego.Error(err)
	}
	this.TplName = "bgview/department.html"
}
