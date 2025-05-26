package domain

type LotteryMessage struct {
	RoomID string `json:"room_id"`
	Number int    `json:"number"`
	Result string `json:"result"`
	Prize  int    `json:"prize"`
}
