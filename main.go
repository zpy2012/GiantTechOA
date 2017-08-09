package main

import (
	_ "GiantTech/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"GiantTech/controllers/tools"
)

func main() {
	//注册数据库
	orm.RegisterDataBase("default", "mysql", "zpy:12345@tcp(127.0.0.1:3306)/GiantTech")

	beego.SetStaticPath("/department/static", "static")
	beego.SetStaticPath("/department/edit/static", "static")
	beego.SetStaticPath("/department/add/static", "static")
	beego.SetStaticPath("/user/static", "static")
	beego.SetStaticPath("/user/edit/static", "static")
	beego.SetStaticPath("/user/add/static", "static")
	beego.SetStaticPath("/project/static", "static")
	beego.SetStaticPath("/project/edit/static", "static")
	beego.SetStaticPath("/project/add/static", "static")
	beego.SetStaticPath("/project/file/static", "static")
	beego.SetStaticPath("/project/info/static", "static")
	beego.SetStaticPath("/project/upload/static", "static")
	beego.SetLogger("file", `{"filename":"logs/test.log"}`)
	beego.AddFuncMap("date", tools.FormatDate)
	beego.AddFuncMap("add", tools.Add)
	beego.AddFuncMap("departmentName", tools.SearchDepartment)
	beego.AddFuncMap("projectStatus", tools.CheckProjectStatus)
	beego.AddFuncMap("projectType", tools.CheckProjectType)
	beego.AddFuncMap("projectOwner", tools.CheckProjectOwner)
	beego.AddFuncMap("projectSource", tools.CheckProjectSource)
	beego.AddFuncMap("userLevel", tools.SearchUserLevel)
	beego.AddFuncMap("projectOperation", tools.CheckProjectOperation)
	beego.AddFuncMap("projectOperationStatus", tools.CheckProjectOperationStatus)
	beego.Run()
}

