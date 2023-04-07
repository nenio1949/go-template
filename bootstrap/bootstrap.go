package bootstrap

import (
	"go-template/global"
	"go-template/models"
)

func Init() {
	// 初始化配置
	InitializeConfig()

	// 初始化日志
	global.App.Log = InitializeLog()
	global.App.Log.Info("log init success!")

	// 初始化数据库
	global.App.DB = models.InitializeDB()

	RunServer()
}
