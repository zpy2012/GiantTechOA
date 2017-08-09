package tools

import (
	"os"
	"strings"
	"mime/multipart"
	"io"
	"fmt"
	"strconv"
	"github.com/astaxie/beego"
)

//保存文件
func SaveFile(file multipart.File, handler *multipart.FileHeader, num int64, filePath string) (name string, path string, err error)  {
	if err != nil {
		return
	}
	defer file.Close()
	if !IsDirExist(filePath) {
		os.MkdirAll(filePath, 0777)
	}
	var fileName string
	n := strings.Split(handler.Filename, ".")
	if num > 0 {
		for i := 0; i < len(n)-1; i++ {
			fileName = fileName + n[i]
		}
		fileName = fileName + "(" + strconv.FormatInt(num, 10) + ")." + n[len(n)-1]
	}else {
		fileName = handler.Filename
	}
	dir := filePath + "/" + fileName
	f, err := os.Create(dir)
	CheckError(err)
	defer f.Close()
	a, err := io.Copy(f, file)
	fmt.Println(a)
	return fileName, dir, err
}

//保存文件
func SaveUserHeadImage(file multipart.File, handler *multipart.FileHeader, id int64, filePath string) (path string, err error)  {
	if err != nil {
		return
	}
	defer file.Close()
	if !IsDirExist(filePath) {
		os.MkdirAll(filePath, 0777)
	}
	var fileName string
	n := strings.Split(handler.Filename, ".")
	t := strings.Replace(FormatDate(TimeNow())," ", "_", 1)
	t = strings.Replace(FormatDate(TimeNow()),":", "_", 2)
	fileName = strconv.FormatInt(id, 10) + t + "." + n[len(n)-1]
	dir := filePath + fileName
	f, err := os.Create(dir)
	CheckError(err)
	defer f.Close()
	a, err := io.Copy(f, file)
	beego.Informational(a)
	return dir, err
}

//删除图片
func DeleteImageWithPath(path string) (err error) {
	p := strings.Split(path, "/")
	filePath := "."+"/"+p[len(p)-3]+"/"+p[len(p)-2]+"/"+p[len(p)-1]
	if err := os.Remove(filePath);err !=nil {
		return nil
	}else {
		CheckError(err)
		return err
	}
}

//删除文件
func DeleteFileWithPath(path string) (err error) {
	p := strings.Split(path, "/")
	filePath := "."+"/"+p[len(p)-4] + "/"+p[len(p)-3]+"/"+p[len(p)-2]+"/"+p[len(p)-1]
	if err := os.Remove(filePath);err !=nil {
		return nil
	}else {
		CheckError(err)
		return err
	}
}