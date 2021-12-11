package main

import (
	"gin-blog-server/dao"
	"gin-blog-server/global"
	"gin-blog-server/initialize"
	"log"
)

func main() {
	global.AppViper = initialize.InitViper()

	initialize.InitGorm()
	sqlDB := dao.SqlDB()
	defer sqlDB.Close()

	err := initialize.Run()
	if err != nil {
		log.Fatal("Listen and serve error: ", err.Error())
	}
}
