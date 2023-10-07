package core

import (
	e "q6-golang/enumerators"

	"github.com/shopspring/decimal"
)

type Winner struct {
	GameType             e.GameType
	PrizeWinnerList      []Player
	PrizeAmountPerWinner decimal.Decimal
}

func GetWinner(gameType e.GameType, winners []Player, prizeAmountTotal decimal.Decimal) Winner {
	var prizeAmountPerWinner decimal.Decimal
	numberOfWinners := len(winners)
	numberOfWinners_d := decimal.NewFromInt(int64(numberOfWinners))
	if numberOfWinners == 0 {
		prizeAmountPerWinner = decimal.NewFromInt(0)
	} else {
		prizeAmountPerWinner = prizeAmountTotal.Div(numberOfWinners_d)
	}
	return Winner{gameType, winners, prizeAmountPerWinner}
}
