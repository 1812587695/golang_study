package main

// 实例一

import (
	"hytx_sync/pkg/setting"
	"hytx_sync/models"
	"hytx_sync/pkg/logging"
	"hytx_sync/pkg/gredis"
)

func init() {
	setting.Setup()
	models.Setup()
	logging.Setup()
	gredis.Setup()
}

func main() {

}

