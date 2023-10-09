package handlers

import (
	"q6-golang/models"
	"q6-golang/utils"
)

func ExecuteGame(players []models.Player) models.Game {
	utils.PrintProgramStartup()

	sales := utils.CalculateTotalSales(players)
	utils.PrintTotalSales(sales)

	prizes := models.GetPrizes(sales.TradicionalSales, sales.RevanchaSales, sales.SiempreSaleSales)
	utils.PrintPrizes(prizes)

	drawings := utils.ExecuteGames(players)
	utils.PrintDrawingResults(drawings)

	winners := utils.CalculateWinners(players, drawings, prizes)
	utils.PrintWinners(drawings, prizes, winners)

	return models.Game{TotalSales: sales, Prizes: prizes, Results: drawings, Winners: winners}
}
