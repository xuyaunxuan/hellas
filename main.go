package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"hellas/common/setting"
	"hellas/models"
	"hellas/routers"
	"log"
	"net/http"
)

func init() {
	//setting.init()
	models.Setup()
	//logging.Setup()
	//gredis.Setup()
	//util.Setup()
}

func main() {
	gin.SetMode(setting.RunMode)

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
