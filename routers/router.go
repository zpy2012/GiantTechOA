package routers

import (
	"github.com/astaxie/beego"
	"GiantTech/controllers/bgControllers"
)

func init() {
	beego.Router("/", &bgControllers.BgIndexController{})
	beego.Router("/login", &bgControllers.BgLoginController{})
	beego.Router("/logout", &bgControllers.BgLogoutController{})
	beego.Router("/department", &bgControllers.BgDepartmentController{})
	beego.Router("/department/edit", &bgControllers.BgDepartmentEditController{})
	beego.Router("/department/add", &bgControllers.BgDepartmentAddController{})
	beego.Router("/department/delete", &bgControllers.BgDepartmentDeleteController{})
	beego.Router("/user", &bgControllers.BgUserController{})
	beego.Router("/user/info", &bgControllers.BgUserInfoController{})
	beego.Router("/user/add", &bgControllers.BgUserAddController{})
	beego.Router("/user/edit", &bgControllers.BgUserEditController{})
	beego.Router("/user/delete", &bgControllers.BgUserDeleteController{})
	beego.Router("/user/updatePwd", &bgControllers.BgUserUpdatePassWordController{})
	beego.Router("/user/updateimage", &bgControllers.BgUserUpdateImageController{})
	beego.Router("/project", &bgControllers.BgProjectController{})
	beego.Router("/project/add", &bgControllers.BgProjectAddController{})
	beego.Router("/project/edit", &bgControllers.BgProjectEditController{})
	beego.Router("/project/delete", &bgControllers.BgProjectDeleteController{})
	beego.Router("/project/operation", &bgControllers.BgProjectOperationController{})
	beego.Router("/project/file", &bgControllers.BgProjectFileController{})
	beego.Router("/project/info", &bgControllers.BgProjectInfoController{})
	beego.Router("/project/upload", &bgControllers.BgProjectUploadFileController{})
	beego.Router("/project/deleteFile", &bgControllers.BgProjectFileDeleteController{})
	beego.Router("/project/export", &bgControllers.BgProjectExportController{})
	beego.Router("/getCode", &bgControllers.BgCodeController{})
	beego.Router("/compareCode", &bgControllers.BgCodeController{})
}
