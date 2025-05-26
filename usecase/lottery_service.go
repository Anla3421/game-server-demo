package usecase

import (
	"websocket-lottery/domain"
)

type LotteryService struct{}

func (s *LotteryService) ProcessLottery(number int) domain.LotteryMessage {
	result := domain.LotteryMessage{
		Number: number,
	}

	switch {
	case number == 7:
		result.Result = "特獎"
		result.Prize = 10000
	case number%2 == 0:
		result.Result = "中獎"
		result.Prize = 50
	default:
		result.Result = "未中獎"
		result.Prize = 0
	}

	return result
}
