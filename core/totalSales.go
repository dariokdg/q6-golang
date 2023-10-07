package core

import (
	"github.com/shopspring/decimal"
)

type TotalSales struct {
	TotalTradicionalSales   decimal.Decimal `json:"totalTradicionalSales"`
	TotalTradicionalPlayers int             `json:"totalTradicionalPlayers"`
	TotalRevanchaSales      decimal.Decimal `json:"totalRevanchaSales"`
	TotalRevanchaPlayers    int             `json:"totalRevanchaPlayers"`
	TotalSiempreSaleSales   decimal.Decimal `json:"totalSiempreSaleSales"`
	TotalSiempreSalePlayers int             `json:"totalSiempreSalePlayers"`
}

func GetTotalGameSales(totals TotalSales) decimal.Decimal {
	return decimal.Sum(totals.TotalTradicionalSales, totals.TotalRevanchaSales, totals.TotalSiempreSaleSales)
}

func GetTotalGamePlayers(totals TotalSales) int {
	return totals.TotalTradicionalPlayers + totals.TotalRevanchaPlayers + totals.TotalSiempreSalePlayers
}
