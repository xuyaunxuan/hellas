package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"hellas/common/setting"
	"hellas/models"
	"hellas/routers"
	"log"
	"net/http"
	"os"
)

func init() {
	// 环境变量读取
	setting.Init()
	// 数据库连接
	models.Setup()
}

func main() {
	gin.SetMode(setting.RunMode)
	file, _ := os.Create("access.log")
	gin.DefaultWriter = file

	routersInit := routers.InitRouter()
	readTimeout := setting.ReadTimeout
	writeTimeout := setting.WriteTimeout
	endPoint := fmt.Sprintf(":%d", setting.HTTPPort)
	maxHeaderBytes := 1 << 20

	server := &http.Server{
		Addr:           endPoint,
		Handler:        routersInit,
		ReadTimeout:    readTimeout,
		WriteTimeout:   writeTimeout,
		MaxHeaderBytes: maxHeaderBytes,
	}

	log.Printf("[info] start http server listening %s", endPoint)

	server.ListenAndServe()
}
