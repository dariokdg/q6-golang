package core

import (
	"fmt"

	"github.com/shopspring/decimal"
)

type Q6Game struct {
	TotalSales TotalSales     `json:"sales"`
	Prizes     PrizeGenerator `json:"prizes"`
	Results    GameResults    `json:"results"`
	Winners    []Winner       `json:"winners"`
}

func ExecuteGame(players []Player) Q6Game {
	printProgramStartup()

	tS := calculateTotalSales(players)
	printTotalSales(tS)

	pG := getPrizes(tS.TotalTradicionalSales, tS.TotalRevanchaSales, tS.TotalSiempreSaleSales)
	printPrizes(pG)

	drawings := executeGames(players)
	printDrawingResults(drawings)

	winners := calculateWinners(players, drawings, pG)
	printWinners(drawings, pG, winners)

	return Q6Game{tS, pG, drawings, winners}
}

func printProgramStartup() {
	fmt.Println("Starting game...")
}

func calculateTotalSales(players []Player) TotalSales {
	totalTradicionalSales := decimal.NewFromInt(0)
	totalTradicionalPlayers := 0
	totalRevanchaSales := decimal.NewFromInt(0)
	totalRevanchaPlayers := 0
	totalSiempreSaleSales := decimal.NewFromInt(0)
	totalSiempreSalePlayers := 0
	for _, player := range players {
		gs := CheckPlayerSpends(player)
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
	return TotalSales{totalTradicionalSales, totalTradicionalPlayers, totalRevanchaSales, totalRevanchaPlayers, totalSiempreSaleSales, totalSiempreSalePlayers}
}

func calculateWinners(players []Player, drawings GameResults, pG PrizeGenerator) []Winner {
	tFP_winners := CheckPrizesTradicionalFirstPrize(players, drawings.GTRT, pG.TradicionalFirstPrize)
	tSP_winners := CheckPrizesTradicionalSecondPrize(players, drawings.GTRT, pG.TradicionalSecondPrize)
	tTP_winners := CheckPrizesTradicionalThirdPrize(players, drawings.GTRT, pG.TradicionalThirdPrize)

	sFP_winners := CheckPrizesSegundaFirstPrize(players, drawings.GTRS, pG.SegundaFirstPrize)
	sSP_winners := CheckPrizesSegundaSecondPrize(players, drawings.GTRS, pG.SegundaSecondPrize)
	sTP_winners := CheckPrizesSegundaThirdPrize(players, drawings.GTRS, pG.SegundaThirdPrize)

	var validRevanchaPlayers []Player
	for _, p := range players {
		if p.Quini6Ticket.Games == GP_TradicionalAndRevancha || p.Quini6Ticket.Games == GP_TradicionalAndRevanchaAndSiempreSale {
			validRevanchaPlayers = append(validRevanchaPlayers, p)
		}
	}
	r_winners := CheckPrizesRevanchaPrize(validRevanchaPlayers, drawings.GTRR, pG.RevanchaPrize)

	var validSiempreSalePlayers []Player
	for _, p := range players {
		if p.Quini6Ticket.Games == GP_TradicionalAndRevancha || p.Quini6Ticket.Games == GP_TradicionalAndRevanchaAndSiempreSale {
			validSiempreSalePlayers = append(validSiempreSalePlayers, p)
		}
	}
	sS_winners := CheckPrizesSiempreSalePrize(validSiempreSalePlayers, drawings.GTRSS, pG.SiempreSalePrize)

	pE_winners := CheckPrizesPozoExtraPrize(players, drawings.GTRPE, pG.PozoExtraPrize, drawings)

	return []Winner{tFP_winners, tSP_winners, tTP_winners, sFP_winners, sSP_winners, sTP_winners, r_winners, sS_winners, pE_winners}
}
