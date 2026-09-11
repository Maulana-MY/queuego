package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"

	"queuego-backend/config"
	"queuego-backend/routes"
)

func main() {

	// ========================================
	// KONEKSI DATABASE
	// ========================================

	config.ConnectDatabase()

	// ========================================
	// ROUTER
	// ========================================

	router := gin.Default()

	// ========================================
	// TEST API
	// ========================================

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"success": true,
			"message": "QueueGo API berhasil berjalan!",
		})
	})

	// ========================================
	// ROUTES
	// ========================================

	routes.SetupRoutes(router)

	// ========================================
	// SERVER
	// ========================================

	fmt.Println("========================================")
	fmt.Println("      QUEUEGO REST API")
	fmt.Println("========================================")
	fmt.Println("Server berjalan di:")
	fmt.Println("http://localhost:8080")
	fmt.Println("========================================")

	if err := router.Run(":8080"); err != nil {
		log.Fatal("Gagal menjalankan server:", err)
	}
}