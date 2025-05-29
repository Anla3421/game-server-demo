package usecase

import "game-server-demo/domain"

type LotteryService struct{}

func (s *LotteryService) ProcessLottery(msg domain.LotteryMessage) domain.LotteryMessage {
	result := msg

	switch {
	case msg.Number == 7:
		result.Result = "特獎"
		result.Prize = 10000
	case msg.Number%2 == 0:
		result.Result = "中獎"
		result.Prize = 50
	default:
		result.Result = "未中獎"
		result.Prize = 0
	}

	return result
}
