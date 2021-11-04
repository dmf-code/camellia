package main

import (
	"camellia/internal/admin/application/route"
	"camellia/internal/admin/cmd"
	"camellia/tool/database"
	"camellia/tool/logObj"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"io"
	"log"
	"os"
	"time"
)

func main() {
	// 配置日志
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	// 加载.env配置
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// 初始化日志
	logObj.InitLogger()
	defer logObj.SugarLogger.Sync()

	logObj.SugarLogger.Infof("aaaaaaaaaaaa %d", 1)
	
	// 初始化数据库
	database.InitMysql(&database.Config{
		Addr:         "",
		DSN:          "root:root@tcp(127.0.0.1:3306)/camellia?charset=utf8mb4&parseTime=True&loc=Local",
		ReadDSN:      nil,
		Active:       10,
		Idle:         5,
		IdleTimeout:  time.Duration(time.Minute),
		QueryTimeout: time.Duration(time.Minute),
		ExecTimeout:  time.Duration(time.Minute),
		TranTimeout:  time.Duration(time.Minute),
	})


	cmd.MigrationCmd.Execute()

	// 权限初始化
	//permission.Init()

	r, err := route.SetupRouter()

	if err != nil {
		panic("初始化路由失败： " + err.Error())
	}

	// Listen and Server in 0.0.0.0:8080
	r.Run(":8888")
}
