package models

type ErrorResult struct {
	Msg string
	Code int
}

type Result struct {
	Msg string
	Code int
	JumpUrl string
}

type AwadResult struct {
	Msg string
	Awd int
}

type Data struct {
	Id string	`json:"id"`
} 

type UploadResult struct {
//{"status":true,"data":{"id":"431cbf5cfe3e45c4abcea878723d7b89"},"message":"操作成功"}
	Status bool	`json:"status"`
	Data Data	`json:"data"`
	Message string	`json:"message"`
}

type NoticeModel struct {
	Name string		`json:"name"`
	OwnerName string	`json:"ownerName"`
	ProjectName string	`json:"projectName"`
}

type PriceNoticeModel struct {
	Name string		`json:"UserRealName"`
	ProjectName string	`json:"ProjectName"`
}

type NewProjectNotice struct {
	Name string		`json:"UserRealName"`
	ProjectName string	`json:"ProjectName"`
}

type NewUserNotice struct {
	Name string		`json:"UserRealName"`
	AdminName string	`json:"AmdinRealName"`
	UserName string		`json:"UserName"`
	PassWord string		`json:"PassWord"`
}