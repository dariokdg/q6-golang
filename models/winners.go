package models

import "github.com/shopspring/decimal"

type Winners struct {
	GameType             GameType        `json:"gameType"`
	PrizeType            PrizeType       `json:"prizeType"`
	NumberOfMatches      int             `json:"numberOfMatches"`
	PrizeWinnerList      []Ticket        `json:"prizeWinnerList"`
	PrizeAmountPerWinner decimal.Decimal `json:"prizeAmountPerWinner"`
}
