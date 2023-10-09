package utils

import (
	"q6-golang/models"

	"github.com/shopspring/decimal"
)

func CalculateTotalSales(players []models.Player) models.Sales {
	totalTradicionalSales := decimal.NewFromInt(0)
	totalTradicionalPlayers := 0
	totalRevanchaSales := decimal.NewFromInt(0)
	totalRevanchaPlayers := 0
	totalSiempreSaleSales := decimal.NewFromInt(0)
	totalSiempreSalePlayers := 0
	for _, player := range players {
		gs := models.CheckPlayerSpends(player)
		totalTradicionalSales = totalTradicionalSales.Add(gs.TradicionalSpends)
		totalRevanchaSales = totalRevanchaSales.Add(gs.RevanchaSpends)
		totalSiempreSaleSales = totalSiempreSaleSales.Add(gs.SiempreSaleSpends)
		if gs.SiempreSaleSpends.GreaterThan(decimal.NewFromInt(0)) {
			totalSiempreSalePlayers++
		} else if gs.RevanchaSpends.GreaterThan(decimal.NewFromInt(0)) {
			totalRevanchaPlayers++
		} else {
			totalTradicionalPlayers++
		}
	}
	return models.Sales{TradicionalSales: totalTradicionalSales, TradicionalPlayers: totalTradicionalPlayers, RevanchaSales: totalRevanchaSales, RevanchaPlayers: totalRevanchaPlayers, SiempreSaleSales: totalSiempreSaleSales, SiempreSalePlayers: totalSiempreSalePlayers}
}
