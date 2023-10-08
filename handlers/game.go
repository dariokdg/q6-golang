package handlers

import (
	"fmt"
	"q6-golang/models"
	"q6-golang/utils"

	"github.com/shopspring/decimal"
)

func ExecuteGame(players []models.Player) models.Game {
	printProgramStartup()

	sales := calculateTotalSales(players)
	utils.PrintTotalSales(sales)

	prizes := models.GetPrizes(sales.TotalTradicionalSales, sales.TotalRevanchaSales, sales.TotalSiempreSaleSales)
	utils.PrintPrizes(prizes)

	drawings := utils.ExecuteGames(players)
	utils.PrintDrawingResults(drawings)

	winners := calculateWinners(players, drawings, prizes)
	utils.PrintWinners(drawings, prizes, winners)

	return models.Game{TotalSales: sales, Prizes: prizes, Results: drawings, Winners: winners}
}

func printProgramStartup() {
	fmt.Println("Starting game...")
}

func calculateTotalSales(players []models.Player) models.TotalSales {
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
	return models.TotalSales{TotalTradicionalSales: totalTradicionalSales, TotalTradicionalPlayers: totalTradicionalPlayers, TotalRevanchaSales: totalRevanchaSales, TotalRevanchaPlayers: totalRevanchaPlayers, TotalSiempreSaleSales: totalSiempreSaleSales, TotalSiempreSalePlayers: totalSiempreSalePlayers}
}

func calculateWinners(players []models.Player, drawings models.GameResults, pG models.PrizeGenerator) []models.Winner {
	tFP_winners := utils.CheckPrizesTradicionalFirstPrize(players, drawings.GTRT, pG.TradicionalFirstPrize)
	tSP_winners := utils.CheckPrizesTradicionalSecondPrize(players, drawings.GTRT, pG.TradicionalSecondPrize)
	tTP_winners := utils.CheckPrizesTradicionalThirdPrize(players, drawings.GTRT, pG.TradicionalThirdPrize)

	sFP_winners := utils.CheckPrizesSegundaFirstPrize(players, drawings.GTRS, pG.SegundaFirstPrize)
	sSP_winners := utils.CheckPrizesSegundaSecondPrize(players, drawings.GTRS, pG.SegundaSecondPrize)
	sTP_winners := utils.CheckPrizesSegundaThirdPrize(players, drawings.GTRS, pG.SegundaThirdPrize)

	var validRevanchaPlayers []models.Player
	for _, p := range players {
		if p.Quini6Ticket.Games == models.GP_TradicionalAndRevancha || p.Quini6Ticket.Games == models.GP_TradicionalAndRevanchaAndSiempreSale {
			validRevanchaPlayers = append(validRevanchaPlayers, p)
		}
	}
	r_winners := utils.CheckPrizesRevanchaPrize(validRevanchaPlayers, drawings.GTRR, pG.RevanchaPrize)

	var validSiempreSalePlayers []models.Player
	for _, p := range players {
		if p.Quini6Ticket.Games == models.GP_TradicionalAndRevancha || p.Quini6Ticket.Games == models.GP_TradicionalAndRevanchaAndSiempreSale {
			validSiempreSalePlayers = append(validSiempreSalePlayers, p)
		}
	}
	sS_winners := utils.CheckPrizesSiempreSalePrize(validSiempreSalePlayers, drawings.GTRSS, pG.SiempreSalePrize)

	pE_winners := utils.CheckPrizesPozoExtraPrize(players, drawings.GTRPE, pG.PozoExtraPrize, drawings)

	return []models.Winner{tFP_winners, tSP_winners, tTP_winners, sFP_winners, sSP_winners, sTP_winners, r_winners, sS_winners, pE_winners}
}
