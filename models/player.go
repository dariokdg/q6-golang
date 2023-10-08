package models

import (
	"github.com/shopspring/decimal"
)

type Player struct {
	Name         string          `json:"playerName"`
	Age          int             `json:"playerAge"`
	City         string          `json:"playerCity"`
	Quini6Ticket Ticket          `json:"playerTicket"`
	PrizeMoney   decimal.Decimal `json:"playerPrizeMoney"`
	MoneySpent   decimal.Decimal `json:"playerMoneySpent"`
}

func GetPlayer(name string, age int, city string, q6Ticket Ticket) Player {
	return Player{name, age, city, q6Ticket, decimal.NewFromInt(0), q6Ticket.Cost}
}
