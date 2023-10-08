package models

import "github.com/shopspring/decimal"

type Winner struct {
	GameType             GameType        `json:"gameType"`
	PrizeType            PrizeType       `json:"prizeType"`
	NumberOfMatches      int             `json:"numberOfMatches"`
	PrizeWinnerList      []Player        `json:"prizeWinnerList"`
	PrizeAmountPerWinner decimal.Decimal `json:"prizeAmountPerWinner"`
}

func GetWinners(gameType GameType, prizeType PrizeType, numberOfMatches int, winners []Player, prizeAmountTotal decimal.Decimal) Winner {
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
	return Winner{gameType, prizeType, numberOfMatches, updatedPlayers, prizeAmountPerWinner}
}
