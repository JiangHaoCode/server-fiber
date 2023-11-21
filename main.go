package main

import (
	"server/core"
	"server/global"
	"server/initialize"

	utilsInit "server/utils"

	"go.uber.org/zap"
)

//go:generate go env -w GO111MODULE=on
//go:generate go env -w GOPROXY=https://goproxy.cn,direct
//go:generate go mod tidy
//go:generate go mod download

// @title Swagger Example API
// @version 0.0.1
// @description This is a sample Server pets
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name x-token
// @BasePath /
func main() {
	global.VIP = core.Viper() // 初始化Viper 配置
	global.LOG = core.Zap()   // 初始化zap日志库
	// global.Logger = core.InitLogger() // 初始化 log 让log标准输出
	zap.ReplaceGlobals(global.LOG) // 部署到全局

	db, err := initialize.Gorm() // gorm连接数据库
	if err == nil {
		global.DB = db
	} else {
		global.LOG.Error(err.Error() + ": 数据库链接失败")
	}
	// initialize.Timer() //定时清除数据库数据
	// conn, err := global.Timer.FindCron("ClearDB")
	initialize.Tasks() //定时 执行任务
	// initialize.DBList()
	utilsInit.TransInit("zh")
	if global.DB != nil {
		// initialize.RegisterTables(global.DB) // 初始化表
		// 程序结束前关闭数据库链接
		db, _ := global.DB.DB()
		defer db.Close()
	}
	core.RunWindowsServer()
}
