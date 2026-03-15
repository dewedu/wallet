package main

import (
	"log"
	"wallet/internal/handler"
	"wallet/internal/service"

	"gorm.io/driver/mysql"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func main() {
	// 数据库连接配置
	dsn := "root:1234@tcp(localhost:3306)/walletdb?charset=utf8mb4&parseTime=True&loc=Local"

	// 初始化数据库连接
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// 设置 Gin
	r := gin.Default()

	walletService := service.NewWalletService(db)
	walletHandler := handler.NewWalletHandler(walletService)

	// 路由
	r.POST("/wallets", walletHandler.CreateWallet)      // 获取一个唯一的钱包 ID
	r.GET("/wallets/:id", walletHandler.GetWallet)      // 获取用户钱包余额
	r.POST("/wallets/transfer", walletHandler.Transfer) // 转账

	// 启动服务器
	log.Println("Wallet service starting on :8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Failed to run server:", err)
	}

}
