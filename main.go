package main

import (
	"log"
	"websocket-lottery/infrastructure"
)

func main() {
	server := infrastructure.NewWebSocketServer()

	log.Println("啟動 WebSocket 伺服器...")
	err := server.Start(":8080")
	if err != nil {
		log.Fatal("伺服器啟動失敗:", err)
	}
}
