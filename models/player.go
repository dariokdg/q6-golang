package models

import (
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Player struct {
	ID         uuid.UUID       `json:"UUID"`
	Name       string          `json:"name"`
	Age        int             `json:"age"`
	City       string          `json:"city"`
	Tickets    []Ticket        `json:"tickets"`
	PrizeMoney decimal.Decimal `json:"prizeMoney"`
	MoneySpent decimal.Decimal `json:"moneySpent"`
}

func GetPlayer(name string, age int, city string) Player {
	/*
		moneySpent := decimal.NewFromInt(0)
		for _, ticket := range tickets {
			moneySpent = decimal.Sum(moneySpent, ticket.Cost)
		}
	*/
	return Player{uuid.New(), name, age, city, []Ticket{}, decimal.NewFromInt(0), decimal.NewFromInt(0)}
}

func AssignTicketToPlayer(player *Player, ticket *Ticket) {
	player.Tickets = append(player.Tickets, *ticket)
	player.MoneySpent = decimal.Sum(player.MoneySpent, ticket.Cost)
}

func AssignTicketsToPlayer(player Player, tickets []Ticket) {
	for _, ticket := range tickets {
		AssignTicketToPlayer(&player, &ticket)
	}
}
