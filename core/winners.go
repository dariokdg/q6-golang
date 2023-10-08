package core

import "github.com/shopspring/decimal"

type Winner struct {
	GameType             GameType        `json:"gameType"`
	PrizeType            PrizeType       `json:"prizeType"`
	PrizeWinnerList      []Player        `json:"prizeWinnerList"`
	PrizeAmountPerWinner decimal.Decimal `json:"prizeAmountPerWinner"`
}

func GetWinners(gameType GameType, prizeType PrizeType, winners []Player, prizeAmountTotal decimal.Decimal) Winner {
	var prizeAmountPerWinner decimal.Decimal
	numberOfWinners := len(winners)
	numberOfWinners_d := decimal.NewFromInt(int64(numberOfWinners))
	if numberOfWinners == 0 {
		prizeAmountPerWinner = decimal.NewFromInt(0)
	} else {
		prizeAmountPerWinner = prizeAmountTotal.Div(numberOfWinners_d)
	}
	var updatedPlayers []Player
	for _, updatedPlayer := range winners {
		updatedPlayer.PrizeMoney = decimal.Sum(updatedPlayer.PrizeMoney, prizeAmountPerWinner)
		updatedPlayers = append(updatedPlayers, updatedPlayer)
	}
	return Winner{gameType, prizeType, updatedPlayers, prizeAmountPerWinner}
}
