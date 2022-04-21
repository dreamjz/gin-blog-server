package main

import (
	"gin-blog-server/dao"
	"gin-blog-server/global"
	"gin-blog-server/initialize"
	"log"
)

func main() {
	// 初始化 Viper，读取配置文件
	global.AppViper = initialize.InitViper()

	// 初始化 Gorm
	initialize.InitGorm()
	// 手动关闭连接
	sqlDB := dao.SqlDB()
	defer sqlDB.Close()

	err := initialize.Run()
	if err != nil {
		log.Fatal("Listen and serve error: ", err.Error())
	}
}
