package env

import (
	"../common"
	"../models"
	"log"
)

func InitEnv() error {

	// 初始化日志
	err := common.InitLogger()
	if err != nil {
		log.Fatal("初始化日志失败，err:", err)
		return err
	}

	// 初始化postgres数据库
	err = models.InitDB()
	if err != nil {
		log.Fatal("初始化数据库失败，err:", err)
		return err
	}
	return nil
}