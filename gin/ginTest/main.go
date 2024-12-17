package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	// 创建一个 Gin 路由引擎
	r := gin.Default()

	// 使用自定义中间件
	r.Use(MyMiddleware())

	// 路由定义
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello from middleware",
		})
	})

	r.Run(":8080")
}

func MyMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 在请求处理之前执行
		log.Println("Before request")

		c.Next() // 处理请求

		// 在请求处理之后执行
		log.Println("After request")
	}
}
