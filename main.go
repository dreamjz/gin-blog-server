package main

import (
	"gin-blog-server/global"
	"gin-blog-server/initialize"
	"log"
)

func main() {
	global.AppViper = initialize.InitViper()
	err := initialize.Run()
	if err != nil {
		log.Fatal("Listen and serve error: ", err.Error())
	}
}
