package common

import (
	"bytes"
	"github.com/astaxie/beego/logs"
	"io"
	"os"
)

// 判断文件/文件夹是否存在
func PathExist(path string)(bool, error) {
	_, err := os.Stat(path)
	if err != nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
// 创建文件夹
func CreateDir(dirName string) error {
	isExist, err := PathExist(dirName)
	if err != nil {
		logs.Error(err)
		return err
	}
	if !isExist {
		err := os.Mkdir(dirName, os.ModePerm)
		if err != nil {
			logs.Error("mkdir failed.err:", err)
			return err
		}
		logs.Info("mkdir success")
		return nil
	}
    logs.Info(dirName," has been existed")
	return nil
}
// 存储图片
func SaveFile(filePath string, data []byte) error {
	f, err := os.OpenFile(filePath, os.O_WRONLY | os.O_CREATE, 0666)
	if err != nil {
		logs.Error(err)
		return err
	}
	defer f.Close()
	_, err = io.Copy(f, bytes.NewReader(data))
	if err != nil {
		logs.Error(err)
		return err
	}
	return nil
}

// 删除文件
func DeleteFile(filePath string) error {
	err := os.Remove("filePath")
	if err != nil {
		logs.Error(filePath, " deleted failed, err:",err)
        return err
	}
	logs.Info(filePath, " delete successfully")
	return nil
}