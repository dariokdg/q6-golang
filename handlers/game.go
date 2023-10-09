package handlers

import (
	"q6-golang/models"
	"q6-golang/utils"
)

func ExecuteGame(playerList []models.Player) models.Quini6 {
	utils.PrintProgramStartup()

	chPlayers := make(chan models.Players, 1)
	chSales := make(chan models.Sales, 1)
	chDrawings := make(chan models.Q6Results, 1)

	go utils.CalculateTotalPlayers(chPlayers, playerList)
	go utils.CalculateTotalSales(chSales, playerList)
	go utils.ExecuteGames(chDrawings, playerList)

	players := <-chPlayers
	sales := <-chSales
	drawings := <-chDrawings

	utils.PrintTotalPlayers(players)
	utils.PrintTotalSales(sales)
	utils.PrintDrawingResults(drawings)

	prizes := models.GetPrizes(sales.TSales, sales.TRSales, sales.TRSSSales)
	utils.PrintPrizes(prizes)

	winners := utils.CalculateWinners(playerList, drawings, prizes)
	utils.PrintWinners(drawings, prizes, winners)

	return models.Quini6{Players: len(playerList), TotalSales: sales, TotalPlayers: players, Prizes: prizes, Results: drawings, Winners: winners}
}
