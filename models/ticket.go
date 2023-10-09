package models

import "github.com/shopspring/decimal"

type Ticket struct {
	Numbers       []int             `json:"numbers"`
	Participation GameParticipation `json:"participation"`
	Cost          decimal.Decimal   `json:"cost"`
}

func GetTicket(SelectedNumbers []int, Games GameParticipation) Ticket {
	return Ticket{SelectedNumbers, Games, GetTicketCost(Games)}
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
