package core

import (
	e "q6-golang/enumerators"

	"github.com/shopspring/decimal"
)

type Ticket struct {
	SelectedNumbers []int               `json:"selectedNumbers"`
	Games           e.GameParticipation `json:"gameParticipation"`
	Cost            decimal.Decimal     `json:"cost"`
}

func GetTicket(SelectedNumbers []int, Games e.GameParticipation) Ticket {
	return Ticket{SelectedNumbers, Games, GetTicketCost(Games)}
}

func GetTicketCost(Games e.GameParticipation) decimal.Decimal {
	switch Games {
	case e.GP_TradicionalOnly:
		return decimal.NewFromInt(50)
	case e.GP_TradicionalAndRevancha:
		return decimal.NewFromInt(60)
	case e.GP_TradicionalAndRevanchaAndSiempreSale:
		return decimal.NewFromInt(80)
	default:
		return decimal.NewFromInt(50)
	}
}
