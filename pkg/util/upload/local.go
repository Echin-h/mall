package upload

import (
	"fmt"
	conf "gin-mall/conf/sql"
	"gin-mall/pkg/util/log"
	"io/ioutil"
	"mime/multipart"
	"os"
	"strconv"
)

// UploadToLocalStatic uploads to local file
// 1. Get the base path
// 2. Check if the directory exists
//    - if not exist , create a new cirectory
// 3. Create the path for the file
// 4. Read the file
// 5. Write the file
// 6. Return the path

func ProductUploadToLocalStatic(file multipart.File, bossId uint, productName string) (filePath string, err error) {
	bId := strconv.Itoa(int(bossId))
	basePath := "." + conf.Config.PhotoPath.AvatarPath + "user" + bId + "/"
	if !DirExistOrNot(basePath) {
		CreateDir(basePath)
	}
	avatarPath := fmt.Sprintf("%s%s.jpg", basePath, productName)
	f, err := ioutil.ReadAll(file)
	if err != nil {
		log.LogrusObj.Error(err)
		return "", err
	}
	err = ioutil.WriteFile(avatarPath, f, 0666)
	if err != nil {
		log.LogrusObj.Error(err)
		return "", err
	}
	return fmt.Sprintf("user%s/%s.jpg", bId, productName), err
}

func DirExistOrNot(basePath string) bool {
	stat, err := os.Stat(basePath)
	if err != nil {
		log.LogrusObj.Infoln(err)
		return false
	}
	return stat.IsDir()
}

func CreateDir(basePath string) bool {
	err := os.MkdirAll(basePath, os.ModePerm)
	if err != nil {
		return false
	}
	return true
}

func AvatarUploadToLocalStatic(file multipart.File, uid uint, userName string) (string, error) {
	bId := strconv.Itoa(int(uid))
	basePath := "." + conf.Config.PhotoPath.AvatarPath + "user" + bId + "/"
	if !DirExistOrNot(basePath) {
		CreateDir(basePath)
	}
	avatarPath := fmt.Sprintf("%s%s.jpg", basePath, userName)
	conten, err := ioutil.ReadAll(file)
	if err != nil {
		log.LogrusObj.Error(err)
		return "", err
	}
	err = ioutil.WriteFile(avatarPath, conten, 0666)
	if err != nil {
		log.LogrusObj.Error(err)
		return "", err
	}
	return fmt.Sprintf("user%s/%s.jpg", bId, userName), err
}
