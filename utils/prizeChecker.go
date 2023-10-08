package utils

import (
	"q6-golang/models"

	"github.com/shopspring/decimal"
)

func CheckPrizesTradicionalFirstPrize(players []models.Player, results models.GameTypeResult, prize decimal.Decimal) models.Winner {
	drawing := results.DrawingResults
	var winners []models.Player
	for _, p := range players {
		numberOfMatches := GetNumberOfMatches(p.Quini6Ticket.SelectedNumbers, drawing)
		prizeType := CheckMatchesTradicional(numberOfMatches)
		if prizeType == models.PTT_FirstPrize {
			winners = append(winners, p)
		}
	}
	return models.GetWinners(results.GameType, models.PTT_FirstPrize, 6, winners, prize)
}

func CheckPrizesTradicionalSecondPrize(players []models.Player, results models.GameTypeResult, prize decimal.Decimal) models.Winner {
	drawing := results.DrawingResults
	var winners []models.Player
	for _, p := range players {
		numberOfMatches := GetNumberOfMatches(p.Quini6Ticket.SelectedNumbers, drawing)
		prizeType := CheckMatchesTradicional(numberOfMatches)
		if prizeType == models.PTT_SecondPrize {
			winners = append(winners, p)
		}
	}
	return models.GetWinners(results.GameType, models.PTT_SecondPrize, 5, winners, prize)
}

func CheckPrizesTradicionalThirdPrize(players []models.Player, results models.GameTypeResult, prize decimal.Decimal) models.Winner {
	drawing := results.DrawingResults
	var winners []models.Player
	for _, p := range players {
		numberOfMatches := GetNumberOfMatches(p.Quini6Ticket.SelectedNumbers, drawing)
		prizeType := CheckMatchesTradicional(numberOfMatches)
		if prizeType == models.PTT_ThirdPrize {
			winners = append(winners, p)
		}
	}
	return models.GetWinners(results.GameType, models.PTT_ThirdPrize, 4, winners, prize)
}

func CheckPrizesSegundaFirstPrize(players []models.Player, results models.GameTypeResult, prize decimal.Decimal) models.Winner {
	drawing := results.DrawingResults
	var winners []models.Player
	for _, p := range players {
		numberOfMatches := GetNumberOfMatches(p.Quini6Ticket.SelectedNumbers, drawing)
		prizeType := CheckMatchesSegunda(numberOfMatches)
		if prizeType == models.PTS_FirstPrize {
			winners = append(winners, p)
		}
	}
	return models.GetWinners(results.GameType, models.PTS_FirstPrize, 6, winners, prize)
}

func CheckPrizesSegundaSecondPrize(players []models.Player, results models.GameTypeResult, prize decimal.Decimal) models.Winner {
	drawing := results.DrawingResults
	var winners []models.Player
	for _, p := range players {
		numberOfMatches := GetNumberOfMatches(p.Quini6Ticket.SelectedNumbers, drawing)
		prizeType := CheckMatchesSegunda(numberOfMatches)
		if prizeType == models.PTS_SecondPrize {
			winners = append(winners, p)
		}
	}
	return models.GetWinners(results.GameType, models.PTS_SecondPrize, 5, winners, prize)
}

func CheckPrizesSegundaThirdPrize(players []models.Player, results models.GameTypeResult, prize decimal.Decimal) models.Winner {
	drawing := results.DrawingResults
	var winners []models.Player
	for _, p := range players {
		numberOfMatches := GetNumberOfMatches(p.Quini6Ticket.SelectedNumbers, drawing)
		prizeType := CheckMatchesSegunda(numberOfMatches)
		if prizeType == models.PTS_ThirdPrize {
			winners = append(winners, p)
		}
	}
	return models.GetWinners(results.GameType, models.PTS_ThirdPrize, 4, winners, prize)
}

func CheckPrizesRevanchaPrize(players []models.Player, results models.GameTypeResult, prize decimal.Decimal) models.Winner {
	drawing := results.DrawingResults
	var winners []models.Player
	for _, p := range players {
		numberOfMatches := GetNumberOfMatches(p.Quini6Ticket.SelectedNumbers, drawing)
		prizeType := CheckMatchesRevancha(numberOfMatches)
		if prizeType == models.PTR_Prize {
			winners = append(winners, p)
		}
	}
	return models.GetWinners(results.GameType, models.PTR_Prize, 6, winners, prize)
}

func CheckPrizesSiempreSalePrize(players []models.Player, results models.GameTypeResult, prize decimal.Decimal) models.Winner {
	drawing := results.DrawingResults
	var sixMatches []models.Player
	var fiveMatches []models.Player
	var fourMatches []models.Player
	var threeMatches []models.Player
	var twoMatches []models.Player
	var oneMatch []models.Player
	var winners []models.Player
	for _, p := range players {
		numberOfMatches := GetNumberOfMatches(p.Quini6Ticket.SelectedNumbers, drawing)
		prizeType := CheckMatchesSiempreSale(numberOfMatches)
		if prizeType == 6 {
			sixMatches = append(winners, p)
		} else if prizeType == 5 {
			fiveMatches = append(winners, p)
		} else if prizeType == 4 {
			fourMatches = append(winners, p)
		} else if prizeType == 3 {
			threeMatches = append(winners, p)
		} else if prizeType == 2 {
			twoMatches = append(winners, p)
		} else if prizeType == 1 {
			oneMatch = append(winners, p)
		}
	}
	if len(sixMatches) > 0 {
		winners = append(winners, sixMatches...)
		return models.GetWinners(results.GameType, models.PTSS_Prize, 6, winners, prize)
	} else if len(fiveMatches) > 0 {
		winners = append(winners, fiveMatches...)
		return models.GetWinners(results.GameType, models.PTSS_Prize, 5, winners, prize)
	} else if len(fourMatches) > 0 {
		winners = append(winners, fourMatches...)
		return models.GetWinners(results.GameType, models.PTSS_Prize, 4, winners, prize)
	} else if len(threeMatches) > 0 {
		winners = append(winners, threeMatches...)
		return models.GetWinners(results.GameType, models.PTSS_Prize, 3, winners, prize)
	} else if len(twoMatches) > 0 {
		winners = append(winners, twoMatches...)
		return models.GetWinners(results.GameType, models.PTSS_Prize, 2, winners, prize)
	} else {
		winners = append(winners, oneMatch...)
		return models.GetWinners(results.GameType, models.PTSS_Prize, 1, winners, prize)
	}

}

func CheckPrizesPozoExtraPrize(players []models.Player, results models.GameTypeResult, prize decimal.Decimal, drawings models.GameResults) models.Winner {

	//to participate this player must NOT have 6 matches in any of the three main results

	drawing := results.DrawingResults
	var winners []models.Player
	for _, p := range players {
		if p.Quini6Ticket.Games == models.GP_TradicionalOnly {
			//check if users' 6 numbers are identical to tradicional's drawing and segunda's drawing
			if GetNumberOfMatches(p.Quini6Ticket.SelectedNumbers, drawings.GTRT.DrawingResults) == 6 ||
				GetNumberOfMatches(p.Quini6Ticket.SelectedNumbers, drawings.GTRS.DrawingResults) == 6 {
				//this user can't play
				continue
			}
		} else if p.Quini6Ticket.Games == models.GP_TradicionalAndRevancha || p.Quini6Ticket.Games == models.GP_TradicionalAndRevanchaAndSiempreSale {
			//check if users' 6 numbers are identical to tradicional's drawing, segunda's drawing and revancha's drawing
			if GetNumberOfMatches(p.Quini6Ticket.SelectedNumbers, drawings.GTRT.DrawingResults) == 6 ||
				GetNumberOfMatches(p.Quini6Ticket.SelectedNumbers, drawings.GTRS.DrawingResults) == 6 ||
				GetNumberOfMatches(p.Quini6Ticket.SelectedNumbers, drawings.GTRR.DrawingResults) == 6 {
				//this user can't play
				continue
			}
		}
		numberOfMatches := GetNumberOfMatches(p.Quini6Ticket.SelectedNumbers, drawing)
		prizeType := CheckMatchesPozoExtra(numberOfMatches)
		if prizeType == models.PTPE_Prize {
			winners = append(winners, p)
		}
	}
	return models.GetWinners(results.GameType, models.PTPE_Prize, 6, winners, prize)
}
