package models

import "github.com/shopspring/decimal"

type GameSpends struct {
	TradicionalSpends decimal.Decimal `json:"tradicionalSpends"`
	RevanchaSpends    decimal.Decimal `json:"revanchaSpends"`
	SiempreSaleSpends decimal.Decimal `json:"siempreSaleSpends"`
}

func GetTotalSpends(totalSpends decimal.Decimal, games GameParticipation) GameSpends {
	var TS decimal.Decimal
	var RS decimal.Decimal
	var SSS decimal.Decimal
	zero := decimal.NewFromInt(0)
	four := decimal.NewFromInt(4)
	five := decimal.NewFromInt(5)
	six := decimal.NewFromInt(6)
	eight := decimal.NewFromInt(8)
	if games == GP_TradicionalOnly {
		TS = totalSpends
		RS = zero
		SSS = zero
	} else if games == GP_TradicionalAndRevancha {
		TS = totalSpends.Div(six).Mul(five)
		RS = totalSpends.Div(six)
		SSS = zero
	} else if games == GP_TradicionalAndRevanchaAndSiempreSale {
		TS = totalSpends.Div(eight).Mul(five)
		RS = totalSpends.Div(eight)
		SSS = totalSpends.Div(four)
	} else {
		TS = zero
		RS = zero
		SSS = zero
	}
	return GameSpends{TS, RS, SSS}
}

func CheckPlayerSpends(p Player) GameSpends {
	return GetTotalSpends(p.MoneySpent, p.Quini6Ticket.Games)
}
