package core

import (
	"fmt"

	"github.com/shopspring/decimal"
)

type Q6Game struct {
	players    []Player
	totalSales TotalSales
	prizes     PrizeGenerator
	results    GameResults
	winners    []Winner
}

func ExecuteGame(players []Player) Q6Game {
	printProgramStartup()

	tS := calculateTotalSales(players)
	printTotalSales(tS)

	pG := getPrizes(tS.TotalTradicionalSales, tS.TotalRevanchaSales, tS.TotalSiempreSaleSales)
	printPrizes(pG)

	drawings := executeGames(players)
	printDrawingResults(drawings)

	winners := calculateWinners(drawings, pG)
	printWinners(drawings, pG, winners)

	return Q6Game{players, tS, pG, drawings, winners}
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

func calculateWinners(drawings GameResults, pG PrizeGenerator) []Winner {
	tFP_winners := CheckPrizesTradicionalFirstPrize(drawings.GTRT, pG.TradicionalFirstPrize)
	tSP_winners := CheckPrizesTradicionalSecondPrize(drawings.GTRT, pG.TradicionalSecondPrize)
	tTP_winners := CheckPrizesTradicionalThirdPrize(drawings.GTRT, pG.TradicionalThirdPrize)

	sFP_winners := CheckPrizesSegundaFirstPrize(drawings.GTRS, pG.SegundaFirstPrize)
	sSP_winners := CheckPrizesSegundaSecondPrize(drawings.GTRS, pG.SegundaSecondPrize)
	sTP_winners := CheckPrizesSegundaThirdPrize(drawings.GTRS, pG.SegundaThirdPrize)

	r_winners := CheckPrizesRevanchaPrize(drawings.GTRR, pG.RevanchaPrize)

	sS_winners := CheckPrizesSiempreSalePrize(drawings.GTRSS, pG.SiempreSalePrize)

	pE_winners := CheckPrizesPozoExtraPrize(drawings.GTRPE, pG.PozoExtraPrize, drawings)

	return []Winner{tFP_winners, tSP_winners, tTP_winners, sFP_winners, sSP_winners, sTP_winners, r_winners, sS_winners, pE_winners}
}
