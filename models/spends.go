package models

import "github.com/shopspring/decimal"

type Spends struct {
	TradicionalSpends decimal.Decimal `json:"tradicionalSpends"`
	RevanchaSpends    decimal.Decimal `json:"revanchaSpends"`
	SiempreSaleSpends decimal.Decimal `json:"siempreSaleSpends"`
}

func GetTotalSpends(totalSpends decimal.Decimal, tickets []Ticket) Spends {
	TS := decimal.NewFromInt(0)
	RS := decimal.NewFromInt(0)
	SSS := decimal.NewFromInt(0)

	four := decimal.NewFromInt(4)
	five := decimal.NewFromInt(5)
	six := decimal.NewFromInt(6)
	eight := decimal.NewFromInt(8)

	for _, ticket := range tickets {
		if ticket.Participation == GP_TradicionalOnly {
			TS = decimal.Sum(TS, ticket.Cost)
		} else if ticket.Participation == GP_TradicionalAndRevancha {
			TS = decimal.Sum(TS, ticket.Cost.Div(six).Mul(five))
			RS = decimal.Sum(RS, ticket.Cost.Div(six))
		} else if ticket.Participation == GP_TradicionalAndRevanchaAndSiempreSale {
			TS = decimal.Sum(TS, ticket.Cost.Div(eight).Mul(five))
			RS = decimal.Sum(RS, ticket.Cost.Div(eight))
			SSS = decimal.Sum(SSS, ticket.Cost.Div(four))
		}
	}

	return Spends{TS, RS, SSS}
}

func CheckPlayerSpends(p Player) Spends {
	return GetTotalSpends(p.MoneySpent, p.Tickets)
}
