package main

import (
	"fmt"
	"gin-blog-server/global"
	"gin-blog-server/initialize"
	"log"
)

func main() {
	global.AppViper = initialize.InitViper()
	fmt.Printf("%+v", global.AppCfg)
	err := initialize.Run()
	if err != nil {
		log.Fatal("Listen and serve error: ", err.Error())
	}
}
