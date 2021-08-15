package main

import (
	"camellia/route"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"io"
	"log"
	"os"
)

func main() {

	// 框架日志
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	// 加载.env配置
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	r, _ := route.SetupRouter()

	r.Run() // 监听并在 0.0.0.0:8080 上启动服务
}
