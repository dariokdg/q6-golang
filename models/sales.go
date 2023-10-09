package models

import (
	"github.com/shopspring/decimal"
)

type Sales struct {
	TradicionalSales   decimal.Decimal `json:"tradicionalSales"`
	TradicionalPlayers int             `json:"tradicionalPlayers"`
	RevanchaSales      decimal.Decimal `json:"revanchaSales"`
	RevanchaPlayers    int             `json:"revanchaPlayers"`
	SiempreSaleSales   decimal.Decimal `json:"siempreSaleSales"`
	SiempreSalePlayers int             `json:"siempreSalePlayers"`
}

func GetTotalGameSales(totals Sales) decimal.Decimal {
	return decimal.Sum(totals.TradicionalSales, totals.RevanchaSales, totals.SiempreSaleSales)
}

func GetTotalGamePlayers(totals Sales) int {
	return totals.TradicionalPlayers + totals.RevanchaPlayers + totals.SiempreSalePlayers
}
