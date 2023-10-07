package core

import "github.com/shopspring/decimal"

type Winner struct {
	GameType             GameType        `json:"gameType"`
	PrizeWinnerList      []Player        `json:"prizeWinnerList"`
	PrizeAmountPerWinner decimal.Decimal `json:"prizeAmountPerWinner"`
}

func GetWinner(gameType GameType, winners []Player, prizeAmountTotal decimal.Decimal) Winner {
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
