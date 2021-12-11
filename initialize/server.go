package initialize

import (
	"fmt"
	"gin-blog-server/global"
	"gin-blog-server/routers"
	"log"
	"syscall"

	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
)

func initRouter() *gin.Engine {
	router := gin.Default()

	publicGroup := router.Group("/")
	{
		routers.InitPublicRouter(publicGroup)
	}

	return router
}

func Run() error {
	router := initRouter()
	addr := fmt.Sprintf(":%d", global.AppCfg.Server.Port)
	server := endless.NewServer(addr, router)

	server.BeforeBegin = func(addr string) {
		log.Printf("Actual PID: %d,Addr: %s", syscall.Getpid(), addr)
	}

	srvCfg := global.AppCfg.Server
	server.ReadTimeout = srvCfg.ReadTimeout
	server.ReadHeaderTimeout = srvCfg.ReadTimeout
	server.WriteTimeout = srvCfg.WriteTimeout
	server.MaxHeaderBytes = 1 << 20

	return server.ListenAndServe()
}
