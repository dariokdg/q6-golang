package utils

import (
	"q6-golang/models"

	"github.com/shopspring/decimal"
)

func CalculateTotalPlayers(ch chan models.Players, players []models.Player) {
	totalTradicionalPlayers := 0
	totalRevanchaPlayers := 0
	totalSiempreSalePlayers := 0
	for _, player := range players {
		gs := models.CheckPlayerSpends(player)
		if gs.SiempreSaleSpends.GreaterThan(decimal.NewFromInt(0)) {
			totalSiempreSalePlayers++
		} else if gs.RevanchaSpends.GreaterThan(decimal.NewFromInt(0)) {
			totalRevanchaPlayers++
		} else {
			totalTradicionalPlayers++
		}
	}
	ch <- models.Players{TPlayers: totalTradicionalPlayers, TRPlayers: totalRevanchaPlayers, TRSSPlayers: totalSiempreSalePlayers}
}
