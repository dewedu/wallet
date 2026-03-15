package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// 设置 Gin
	r := gin.Default()

	// 路由
	r.POST("/wallets")
	r.GET("/wallets/:id")
	r.POST("/wallets/transfer")

	// 启动服务器
	log.Println("Wallet service starting on :8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Failed to run server:", err)
	}

}
