package models

import (
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Ticket struct {
	Numbers       []int             `json:"numbers"`
	Participation GameParticipation `json:"participation"`
	Cost          decimal.Decimal   `json:"cost"`
	PlayerID      uuid.UUID         `json:"playerUUID"`
}

func GetTicket(SelectedNumbers []int, Games GameParticipation, playerUUID uuid.UUID) Ticket {
	return Ticket{SelectedNumbers, Games, GetTicketCost(Games), playerUUID}
}

func GetTicketOwner(players []Player, ticket Ticket) Player {
	for _, player := range players {
		if player.ID == ticket.PlayerID {
			return player
		}
	}
	return Player{}
}

func GetTicketCost(Games GameParticipation) decimal.Decimal {
	switch Games {
	case GP_TradicionalOnly:
		return decimal.NewFromInt(50)
	case GP_TradicionalAndRevancha:
		return decimal.NewFromInt(60)
	case GP_TradicionalAndRevanchaAndSiempreSale:
		return decimal.NewFromInt(80)
	default:
		return decimal.NewFromInt(50)
	}
}
