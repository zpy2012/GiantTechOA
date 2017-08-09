package tools

import (
	"errors"
	"fmt"
	"golang.org/x/crypto/scrypt"
	"math/rand"
	"net"
	"os"
	"reflect"
	"strconv"
	"time"
	"github.com/tealeg/xlsx"
	"github.com/astaxie/beego"
	"math"
	"GiantTech/models"
	"strings"
)

var CurrentPath string = getCurrentPath()

type OperationType int

const (
	Login OperationType = iota
	Logout
	Addsome
	Edit
	Delete
	Pass
)

type OperationTarget int

const (
	System OperationTarget = iota
	Project
	Department
	Person
	File
)

//检测是否存在该路径
func IsDirExist(path string) bool {
	p, err := os.Stat(path)
	if err != nil {
		return os.IsExist(err)
	} else {
		return p.IsDir()
	}
}

//错误检查
func CheckError(err error) {
	if err != nil {
		beego.Error(err)
		return
	}
}

//***
//[]byte转string
//**
func byteString(p []byte) string {
	for i := 0; i < len(p); i++ {
		if p[i] == 0 {
			return string(p[0:i])
		}
	}
	return string(p)
}

//***
//form匹配struct
//**
//用map填充结构
func FillStruct(data map[string][]string, obj interface{}) error {
	for k, v := range data {
		fmt.Println(v[0])
		err := SetField(obj, k, v[0])
		if err != nil {
			return err
		}
	}
	return nil
}

//用map的值替换结构的值
func SetField(obj interface{}, name string, value interface{}) error {
	structValue := reflect.ValueOf(obj).Elem()        //结构体属性值
	structFieldValue := structValue.FieldByName(name) //结构体单个属性值
	if !structFieldValue.IsValid() {
		return fmt.Errorf("No such field: %s in obj", name)
	}
	if !structFieldValue.CanSet() {
		return fmt.Errorf("Cannot set %s field value", name)
	}
	structFieldType := structFieldValue.Type() //结构体的类型
	val := reflect.ValueOf(value)              //map值的反射值
	var err error
	if structFieldType != val.Type() {
		val, err = TypeConversion(fmt.Sprintf("%v", value), structFieldValue.Type().Name()) //类型转换
		if err != nil {
			return err
		}
	}
	structFieldValue.Set(val)
	return nil
}

//类型转换
func TypeConversion(value string, ntype string) (reflect.Value, error) {
	if ntype == "string" {
		return reflect.ValueOf(value), nil
	} else if ntype == "time.Time" {
		t, err := time.ParseInLocation("2006-01-02 15:04:05", value, time.Local)
		return reflect.ValueOf(t), err
	} else if ntype == "Time" {
		t, err := time.ParseInLocation("2006-01-02 15:04:05", value, time.Local)
		return reflect.ValueOf(t), err
	} else if ntype == "int" {
		i, err := strconv.Atoi(value)
		return reflect.ValueOf(i), err
	} else if ntype == "int8" {
		i, err := strconv.ParseInt(value, 10, 64)
		return reflect.ValueOf(int8(i)), err
	} else if ntype == "int32" {
		i, err := strconv.ParseInt(value, 10, 64)
		return reflect.ValueOf(int64(i)), err
	} else if ntype == "int64" {
		i, err := strconv.ParseInt(value, 10, 64)
		return reflect.ValueOf(i), err
	} else if ntype == "float32" {
		i, err := strconv.ParseFloat(value, 64)
		return reflect.ValueOf(float32(i)), err
	} else if ntype == "float64" {
		i, err := strconv.ParseFloat(value, 64)
		return reflect.ValueOf(i), err
	}
	//else if .......增加其他一些类型的转换
	return reflect.ValueOf(value), errors.New("未知的类型：" + ntype)
}

//获取当前ip地址
func GetIPAddress() string {

	addrs, err := net.InterfaceAddrs()

	if err != nil {
		CheckError(err)
		return ""
	}

	for _, address := range addrs {
		// 检查ip地址判断是否回环地址
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return ""
}

//获取当前项目路径
func getCurrentPath() string {
	file, _ := os.Getwd()
	return file
}

//加密字符串
func EncryptionPassWord(pwd, salt string) string {
	password, _ := scrypt.Key([]byte(pwd), []byte(salt), 16384, 8, 1, 32)
	return fmt.Sprintf("%x", password)
}

//获取当前时间
func TimeNow() time.Time {
	return time.Now().UTC().Add(time.Hour * 8)
}

//生成随机数
func GenerateRandomNumber(start int, end int) int {
	if end > start {
		//随机数生成器，加入时间戳保证每次生成的随机数不一样
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		//生成随机数
		num := r.Intn((end - start)) + start

		return num
	} else if start == end {
		return start
	} else {
		return 0
	}
}

//根据struct导出excel
func ExportExcelWithData(data []interface{}) (string, error) {
	file := xlsx.NewFile()
	sheet, err := file.AddSheet("Sheet1")
	if err != nil {
		return "", err
	}
	row := sheet.AddRow()
	row.SetHeight(30)
	cell := row.AddCell()
	cell.Value = "抽奖码"
	cell = row.AddCell()
	cell.Value = "奖项"
	cell = row.AddCell()
	cell.Value = "省份"
	cell = row.AddCell()
	cell.Value = "城市"
	cell = row.AddCell()
	cell.Value = "姓名"
	cell = row.AddCell()
	cell.Value = "电话"
	cell = row.AddCell()
	cell.Value = "地址"
	cell = row.AddCell()
	cell.Value = "邮编"
	cell = row.AddCell()
	cell.Value = "领奖状态"
	cell = row.AddCell()
	cell.Value = "中奖时间"

	//for _, v := range data {
	//	value := v.(models.TJackpot520)
	//	row := sheet.AddRow()
	//	cell := row.AddCell()
	//	cell.Value = value.RedeemCode
	//	cell = row.AddCell()
	//	cell.Value = value.Award
	//	cell = row.AddCell()
	//	cell.Value = value.Province
	//	cell = row.AddCell()
	//	cell.Value = value.Area
	//	cell = row.AddCell()
	//	cell.Value = value.UserName
	//	cell = row.AddCell()
	//	cell.Value = value.Telephone
	//	cell = row.AddCell()
	//	cell.Value = value.TrugstoreAddress
	//	cell = row.AddCell()
	//	cell.Value = value.ZipCode
	//	cell = row.AddCell()
	//	if value.ReceiveStatus == 1 {
	//		cell.Value = "已领奖"
	//	} else {
	//		cell.Value = "未领奖"
	//	}
	//	cell = row.AddCell()
	//	cell.Value = value.UpdateTime
	//
	//}
	if !IsDirExist(beego.AppConfig.String("exportpath")) {
		os.MkdirAll(beego.AppConfig.String("exportpath"), 0777)
	}
	path := beego.AppConfig.String("exportpath") + beego.AppConfig.String("exportname")
	err = file.Save(path)
	if err != nil {
		return "", err
	} else {
		return path, err
	}
}

//分页方法，根据传递过来的页数，每页数，总数，返回分页的内容 7个页数 前 1，2，3，4，5 后 的格式返回,小于5页返回具体页数
func Paginator(page, prepage int, nums int64) map[string]interface{} {

	//var firstpage int //前一页地址
	//var lastpage int  //后一页地址
	//根据nums总数，和prepage每页数量 生成分页总数
	totalpages := int(math.Ceil(float64(nums) / float64(prepage))) //page总数
	if page > totalpages {
		page = totalpages
	}
	if page <= 0 {
		page = 1
	}
	var pages []int
	switch {
	case page >= totalpages-5 && totalpages > 5: //最后5页
		start := totalpages - 5 + 1
		//firstpage = page - 1
		//lastpage = int(math.Min(float64(totalpages), float64(page+1)))
		pages = make([]int, 5)
		for i, _ := range pages {
			pages[i] = start + i
		}
	case page >= 3 && totalpages > 5:
		start := page - 3 + 1
		pages = make([]int, 5)
		//firstpage = page - 3
		for i, _ := range pages {
			pages[i] = start + i
		}
		//firstpage = page - 1
		//lastpage = page + 1
	default:
		pages = make([]int, int(math.Min(5, float64(totalpages))))
		for i, _ := range pages {
			pages[i] = i + 1
		}
		//firstpage = int(math.Max(float64(1), float64(page-1)))
		//lastpage = page + 1
		//fmt.Println(pages)
	}
	paginatorMap := make(map[string]interface{})
	paginatorMap["pages"] = pages
	paginatorMap["totalpages"] = totalpages
	paginatorMap["firstpage"] = 1
	paginatorMap["lastpage"] = totalpages
	paginatorMap["currpage"] = page
	return paginatorMap
}

//格式化时间模板函数
func FormatDate(in time.Time)(out string){
	out = in.Format(beego.AppConfig.String("time"))
	if strings.EqualFold(out, "0001-01-01 00:00:00") {
		out = "无"
	}
	return
}

//查询部门模板函数
func SearchDepartment(in int)(out string){
	department, err := models.GetTDepartmentById(in)
	if err != nil {
		beego.Error(err)
		out = "请重新分配部门"
	}else {
		out = department.DepartmentName
	}
	return
}

//查询身份模板函数
func SearchUserLevel(in int)(out string){
	switch in {
	case 0: out = "管理员"
	case 1: out = "普通用户"
	}
	return
}

//两数相加模板函数
func Add(a int, b int)(out int){
	out = a + b
	return
}

//判断项目状态模板函数
func CheckProjectStatus(in int)(out string){
	switch in {
	case -1: out = "已放弃"
	case 0: out = "待报价"
	case 1: out = "谈判中"
	case 2: out = "已签订"
	case 3: out = "设计中"
	case 4: out = "开发中"
	case 5: out = "测试中"
	case 6: out = "已完成"
	case 7: out = "已结项"
	}
	return
}

//判断项目操作模板函数
func CheckProjectOperation(userId int, pStatus int)(out string){
	user, _ := models.GetTUsersById(userId)
	if user.UserPrice == -1 {
		switch pStatus {
		case -1: out = "已结束"
		case 0: out = "提醒报价"
		case 1: out = "签约"
		case 7: out = "已结束"
		case 6: out = "提醒结项"
		default: out = "进行中"
		}
	}else {
		switch pStatus {
		case -1: out = "已结束"
		case 1: out = "谈判中"
		case 0: out = "报价"
		case 2: out = "开始项目"
		case 6: out = "结项"
		case 7: out = "已结项"
		default: out = "变更状态"
		}
	}
	return
}

//判断项目操作模板函数
func CheckProjectOperationStatus(userId int, pStatus int)(out string){
	user, _ := models.GetTUsersById(userId)
	if user.UserPrice == -1 {
		switch pStatus {
		case -1: out = "disabled"
		case 2: out = "disabled"
		case 3: out = "disabled"
		case 4: out = "disabled"
		case 5: out = "disabled"
		case 7: out = "disabled"
		}
	}else {
		switch pStatus {
		case -1: out = "disabled"
		case 1: out = "disabled"
		case 7: out = "disabled"
		}
	}
	return
}

//判断项目来源模板函数
func CheckProjectSource(in int)(out string){
	switch in {
	case 0: out = "线上"
	case 1: out = "线下"
	}
	return
}

//判断项目类型模板函数
func CheckProjectType(in int)(out string){
	switch in {
	case 0: out = "设计类"
	case 1: out = "开发类"
	case 2: out = "策划类"
	}
	return
}

//获取项目创建者模板函数
func CheckProjectOwner(id int)(out string){
	user, _ := models.GetTUsersById(id)
	out = user.UserRealName
	return
}

//添加日志
func AddLog(name string, optype OperationType, target OperationTarget, msg string){
	var log models.TLogs
	log.LogUserRealName = name
	switch optype {
	case Login: log.LogOperationType = 0
	case Logout: log.LogOperationType = -1
	case Addsome: log.LogOperationType = 1
	case Edit: log.LogOperationType = 2
	case Delete: log.LogOperationType = 3
	case Pass: log.LogOperationType = 4
	}
	switch target {
	case System: log.LogOperationTarget = 0
	case Project: log.LogOperationTarget = 1
	case Department: log.LogOperationTarget = 2
	case Person: log.LogOperationTarget = 3
	case File: log.LogOperationTarget = 4
	}
	log.LogOperationMsg = msg
	log.LogOperationTime = TimeNow()
	if _, err := models.AddTLogs(&log); err != nil {
		CheckError(err)
	}
}

//生成随机验证码
func RandCode() string  {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	num := beego.AppConfig.String("code")
	num1, _ := strconv.ParseInt(num, 10, 64)
	if num1 == 4 {
		vcode := fmt.Sprintf("%06v", r.Int31n(10000))
		return vcode
	}else {
		vcode := fmt.Sprintf("%06v", r.Int31n(1000000))
		return vcode
	}
}

//项目导出excel
//根据struct导出excel
func ExportProjectsExcel() (string, error) {
	var order, sortby []string
	order = append(order, "desc")
	sortby = append(sortby, "ProjectCreatedTime")
	projects, err := models.GetAllTProjects(nil, nil, sortby, order, 0, 0)
	beego.Informational(err)
	file := xlsx.NewFile()
	sheet, err := file.AddSheet("项目表")
	if err != nil {
		beego.Error(err)
		return "", err
	}
	row := sheet.AddRow()
	row.SetHeight(30)
	cell := row.AddCell()
	cell.Value = "项目名称"
	cell = row.AddCell()
	cell.Value = "项目介绍"
	cell = row.AddCell()
	cell.Value = "项目来源"
	cell = row.AddCell()
	cell.Value = "项目报价"
	cell = row.AddCell()
	cell.Value = "项目预估工期"
	cell = row.AddCell()
	cell.Value = "项目成交金额"
	cell = row.AddCell()
	cell.Value = "项目成交工期"
	cell = row.AddCell()
	cell.Value = "项目状态"
	cell = row.AddCell()
	cell.Value = "项目类型"
	cell = row.AddCell()
	cell.Value = "项目创建人"
	cell = row.AddCell()
	cell.Value = "项目创建时间"
	cell = row.AddCell()
	cell.Value = "项目开始时间"
	cell = row.AddCell()
	cell.Value = "项目完成时间"
	cell = row.AddCell()
	cell.Value = "项目结项时间"

	for _, v := range projects {
		value := v.(models.TProjects)
		row := sheet.AddRow()
		row.SetHeight(50)
		cell := row.AddCell()
		cell.Value = value.ProjectName
		cell = row.AddCell()
		cell.Value = value.ProjectDescription
		cell = row.AddCell()
		switch value.ProjectSource {
		case 0: cell.Value = "线下"
		case 1: cell.Value = "线上"
		}
		cell = row.AddCell()
		cell.Value = value.ProjectExpectedAmount
		cell = row.AddCell()
		cell.Value = value.ProjectExpectedDuration
		cell = row.AddCell()
		cell.Value = value.ProjectDealAmount
		cell = row.AddCell()
		cell.Value = value.ProjectDealDuration
		cell = row.AddCell()
		switch value.ProjectStatus {
		case -1: cell.Value = "已放弃"
		case 0: cell.Value = "待报价"
		case 1: cell.Value = "谈判中"
		case 2: cell.Value = "已签订"
		case 3: cell.Value = "设计中"
		case 4: cell.Value = "开发中"
		case 5: cell.Value = "测试中"
		case 6: cell.Value = "已完成"
		case 7: cell.Value = "已结项"
		}
		cell = row.AddCell()
		switch value.ProjectType {
		case 0: cell.Value = "设计类"
		case 1: cell.Value = "开发类"
		case 2: cell.Value = "运营类"
		}
		cell = row.AddCell()
		owner, _ := models.GetTUsersById(value.ProjectWonerID)
		cell.Value = owner.UserRealName
		cell = row.AddCell()
		cell.Value = FormatDate(value.ProjectCreatedTime)
		cell = row.AddCell()
		cell.Value = FormatDate(value.ProjectStartTime)
		cell = row.AddCell()
		cell.Value = FormatDate(value.ProjectCompleteTime)
		cell = row.AddCell()
		cell.Value = FormatDate(value.ProjectEndTime)

	}
	if !IsDirExist(beego.AppConfig.String("exportpath")) {
		os.MkdirAll(beego.AppConfig.String("exportpath"), 0777)
	}
	path := beego.AppConfig.String("exportpath") + beego.AppConfig.String("exportname")
	err = file.Save(path)
	if err != nil {
		beego.Error(err)
		return "", err
	} else {
		return path, err
	}
}