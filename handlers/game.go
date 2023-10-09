package handlers

import (
	"q6-golang/models"
	"q6-golang/utils"
)

func ExecuteGame(playerList []models.Player) models.Quini6 {
	utils.PrintProgramStartup()

	sales := utils.CalculateTotalSales(playerList)
	utils.PrintTotalSales(sales)

	players := utils.CalculateTotalPlayers(playerList)
	utils.PrintTotalPlayers(players)

	prizes := models.GetPrizes(sales.TSales, sales.TRSales, sales.TRSSSales)
	utils.PrintPrizes(prizes)

	drawings := utils.ExecuteGames(playerList)
	utils.PrintDrawingResults(drawings)

	winners := utils.CalculateWinners(playerList, drawings, prizes)
	utils.PrintWinners(drawings, prizes, winners)

	return models.Quini6{Players: len(playerList), TotalSales: sales, TotalPlayers: players, Prizes: prizes, Results: drawings, Winners: winners}
}
