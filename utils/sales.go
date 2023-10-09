package utils

import (
	"q6-golang/models"

	"github.com/shopspring/decimal"
)

func CalculateTotalSales(ch chan models.Sales, players []models.Player) {
	totalTradicionalSales := decimal.NewFromInt(0)
	totalRevanchaSales := decimal.NewFromInt(0)
	totalSiempreSaleSales := decimal.NewFromInt(0)
	for _, player := range players {
		gs := models.CheckPlayerSpends(player)
		totalTradicionalSales = totalTradicionalSales.Add(gs.TradicionalSpends)
		totalRevanchaSales = totalRevanchaSales.Add(gs.RevanchaSpends)
		totalSiempreSaleSales = totalSiempreSaleSales.Add(gs.SiempreSaleSpends)
	}
	ch <- models.Sales{TSales: totalTradicionalSales, TRSales: totalRevanchaSales, TRSSSales: totalSiempreSaleSales}
}
