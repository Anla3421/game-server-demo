package infrastructure

import (
	"encoding/json"
	"log"
	"net/http"
	"websocket-lottery/domain"
	"websocket-lottery/usecase"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type WebSocketServer struct {
	lotteryService *usecase.LotteryService
}

func NewWebSocketServer() *WebSocketServer {
	return &WebSocketServer{
		lotteryService: &usecase.LotteryService{},
	}
}

func (s *WebSocketServer) HandleConnections(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()

	// 從 URL 中提取 token
	token := r.URL.Query().Get("token")

	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			break
		}

		var lotteryMsg domain.LotteryMessage
		err = json.Unmarshal(msg, &lotteryMsg)
		if err != nil {
			log.Println("解析錯誤:", err)
			continue
		}

		// 設置 Token
		lotteryMsg.Token = token

		result := s.lotteryService.ProcessLottery(lotteryMsg)

		err = conn.WriteJSON(result)
		if err != nil {
			log.Println("發送錯誤:", err)
			break
		}
	}
}

func (s *WebSocketServer) Start(port string) error {
	http.HandleFunc("/lottery", s.HandleConnections)
	log.Printf("WebSocket 伺服器啟動於 %s", port)
	return http.ListenAndServe(port, nil)
}
