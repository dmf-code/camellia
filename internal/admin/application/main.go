package main

import (
	"camellia/cmd"
	"camellia/internal/admin/application/bootstrap"
	"camellia/internal/admin/application/route"
	"camellia/tool/logObj"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"io"
	"log"
	"os"
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

	if len(os.Args) > 1 {

		if os.Args[1] == "Migration" {

			// 迁移数据
			bootstrap.InitTable()

			err = cmd.MigrationCmd.Execute()

			if err != nil {
				fmt.Println(err)
			}
		}

		return
	}

	// 初始化日志
	logObj.InitLogger()
	defer logObj.SugarLogger.Sync()

	logObj.SugarLogger.Infof("aaaaaaaaaaaa %d", 1)

	// 权限初始化
	permission.Init()

	r, err := route.SetupRouter()

	if err != nil {
		panic("初始化路由失败： " + err.Error())
	}

	// Listen and Server in 0.0.0.0:8080
	r.Run(":8888")
}
