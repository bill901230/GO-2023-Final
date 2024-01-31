package main

import (
	"net/http"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {

	// 初始化Gin
	r := gin.Default()

	// 设置静态文件夹
	r.Static("/static", "./static")

	// 设置HTML模板文件夹
	r.LoadHTMLGlob("../templates/*")

	// 定义首页路由
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "interface.html", nil)
	})

	// 启动Web服务
	fmt.Println("\n\nPlease click the following link:\nhttp://localhost:8083\n\n")
	r.Run(":8083")
}
