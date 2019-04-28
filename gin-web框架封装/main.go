package main

import (
	"fmt"
	"hytx_manager/models"
	"hytx_manager/pkg/gredis"
	"hytx_manager/pkg/logging"
	"hytx_manager/pkg/setting"
	"hytx_manager/routers"
	"net/http"
)

func main()  {
	setting.Setup()
	models.Setup()
	logging.Setup()
	gredis.Setup()

	routersInit := routers.InitRouter()
	readTimeout := setting.ServerSetting.ReadTimeout
	writeTimeout := setting.ServerSetting.WriteTimeout
	endPoint := fmt.Sprintf(":%d", setting.ServerSetting.HttpPort)
	maxHeaderBytes := 1 << 20

	server := &http.Server{
		Addr:           endPoint,
		Handler:        routersInit,
		ReadTimeout:    readTimeout,
		WriteTimeout:   writeTimeout,
		MaxHeaderBytes: maxHeaderBytes,
	}

	server.ListenAndServe()
	return

}