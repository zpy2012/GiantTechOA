package bgControllers

import (
	"github.com/astaxie/beego"
	"GiantTech/models"
	"GiantTech/controllers/tools"
)

type SetupController struct {
	beego.Controller
}

func checkAdminUser() error {
	query := make(map[string]string)
	query["UserLevel"] = "0"
	departments, err := models.GetAllTUsers(query, nil, nil, nil, 0, 0)
	if err != nil {
		beego.Error(err)
		return err
	}
	if len(departments) < 1 {
		defaultHeadImage := beego.AppConfig.String("defaultheadimage")
		var admin models.TUsers
		admin.UserName = "admin"
		admin.UserDepartmentID = -1
		admin.UserRealName = "赵鹏宇"
		admin.UserLevel = 0
		admin.UserPassword = tools.EncryptionPassWord("123456", "zpy")
		admin.UserCreatedTime = tools.TimeNow()
		admin.UserHeadImagePath = defaultHeadImage
		_, err := models.AddTUsers(&admin)
		if err != nil {
			beego.Error(err)
			return err
		}
	}
	return nil
}

func SetUP()  {
	err := checkAdminUser()
	if err != nil {
		beego.Error("检测管理员用户失败!", err)
	}
}
