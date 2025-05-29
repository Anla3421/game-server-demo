package main

import (
	"game-server-demo/infrastructure"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// WebSocket 服務器
	wsServer := infrastructure.NewWebSocketServer()

	go func() {
		log.Println("啟動 WebSocket 服務器...")
		err := wsServer.Start(":5000")
		if err != nil {
			log.Fatal("WebSocket 服務器啟動失敗", err)
		}
	}()

	// Gin HTTP 服務器
	r := gin.Default()

	r.GET("/helloworld", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello, World!",
		})
	})

	log.Println("啟動 HTTP 服務器...")
	err := r.Run(":5001")
	if err != nil {
		log.Fatal("HTTP 服務器啟動失敗", err)
	}
}
