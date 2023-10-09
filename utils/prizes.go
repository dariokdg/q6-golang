package utils

import (
	"q6-golang/models"

	"github.com/shopspring/decimal"
)

func CheckPrizesTradicionalFirstPrize(players []models.Player, results models.Result, prize decimal.Decimal) models.Winners {
	drawing := results.DrawingResults
	var winners []models.Player
	for _, p := range players {
		numberOfMatches := GetNumberOfMatches(p.Quini6Ticket.Numbers, drawing)
		prizeType := CheckMatchesTradicional(numberOfMatches)
		if prizeType == models.PTT_FirstPrize {
			winners = append(winners, p)
		}
	}
	return models.GetWinners(results.GameType, models.PTT_FirstPrize, 6, winners, prize.Round(2))
}

func CheckPrizesTradicionalSecondPrize(players []models.Player, results models.Result, prize decimal.Decimal) models.Winners {
	drawing := results.DrawingResults
	var winners []models.Player
	for _, p := range players {
		numberOfMatches := GetNumberOfMatches(p.Quini6Ticket.Numbers, drawing)
		prizeType := CheckMatchesTradicional(numberOfMatches)
		if prizeType == models.PTT_SecondPrize {
			winners = append(winners, p)
		}
	}
	return models.GetWinners(results.GameType, models.PTT_SecondPrize, 5, winners, prize.Round(2))
}

func CheckPrizesTradicionalThirdPrize(players []models.Player, results models.Result, prize decimal.Decimal) models.Winners {
	drawing := results.DrawingResults
	var winners []models.Player
	for _, p := range players {
		numberOfMatches := GetNumberOfMatches(p.Quini6Ticket.Numbers, drawing)
		prizeType := CheckMatchesTradicional(numberOfMatches)
		if prizeType == models.PTT_ThirdPrize {
			winners = append(winners, p)
		}
	}
	return models.GetWinners(results.GameType, models.PTT_ThirdPrize, 4, winners, prize.Round(2))
}

func CheckPrizesSegundaFirstPrize(players []models.Player, results models.Result, prize decimal.Decimal) models.Winners {
	drawing := results.DrawingResults
	var winners []models.Player
	for _, p := range players {
		numberOfMatches := GetNumberOfMatches(p.Quini6Ticket.Numbers, drawing)
		prizeType := CheckMatchesSegunda(numberOfMatches)
		if prizeType == models.PTS_FirstPrize {
			winners = append(winners, p)
		}
	}
	return models.GetWinners(results.GameType, models.PTS_FirstPrize, 6, winners, prize.Round(2))
}

func CheckPrizesSegundaSecondPrize(players []models.Player, results models.Result, prize decimal.Decimal) models.Winners {
	drawing := results.DrawingResults
	var winners []models.Player
	for _, p := range players {
		numberOfMatches := GetNumberOfMatches(p.Quini6Ticket.Numbers, drawing)
		prizeType := CheckMatchesSegunda(numberOfMatches)
		if prizeType == models.PTS_SecondPrize {
			winners = append(winners, p)
		}
	}
	return models.GetWinners(results.GameType, models.PTS_SecondPrize, 5, winners, prize.Round(2))
}

func CheckPrizesSegundaThirdPrize(players []models.Player, results models.Result, prize decimal.Decimal) models.Winners {
	drawing := results.DrawingResults
	var winners []models.Player
	for _, p := range players {
		numberOfMatches := GetNumberOfMatches(p.Quini6Ticket.Numbers, drawing)
		prizeType := CheckMatchesSegunda(numberOfMatches)
		if prizeType == models.PTS_ThirdPrize {
			winners = append(winners, p)
		}
	}
	return models.GetWinners(results.GameType, models.PTS_ThirdPrize, 4, winners, prize.Round(2))
}

func CheckPrizesRevanchaPrize(players []models.Player, results models.Result, prize decimal.Decimal) models.Winners {
	drawing := results.DrawingResults
	var winners []models.Player
	for _, p := range players {
		numberOfMatches := GetNumberOfMatches(p.Quini6Ticket.Numbers, drawing)
		prizeType := CheckMatchesRevancha(numberOfMatches)
		if prizeType == models.PTR_Prize {
			winners = append(winners, p)
		}
	}
	return models.GetWinners(results.GameType, models.PTR_Prize, 6, winners, prize.Round(2))
}

func CheckPrizesSiempreSalePrize(players []models.Player, results models.Result, prize decimal.Decimal) models.Winners {
	drawing := results.DrawingResults
	var sixMatches []models.Player
	var fiveMatches []models.Player
	var fourMatches []models.Player
	var threeMatches []models.Player
	var twoMatches []models.Player
	var oneMatch []models.Player
	var winners []models.Player
	for _, p := range players {
		numberOfMatches := GetNumberOfMatches(p.Quini6Ticket.Numbers, drawing)
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
		return models.GetWinners(results.GameType, models.PTSS_Prize, 6, winners, prize.Round(2))
	} else if len(fiveMatches) > 0 {
		winners = append(winners, fiveMatches...)
		return models.GetWinners(results.GameType, models.PTSS_Prize, 5, winners, prize.Round(2))
	} else if len(fourMatches) > 0 {
		winners = append(winners, fourMatches...)
		return models.GetWinners(results.GameType, models.PTSS_Prize, 4, winners, prize.Round(2))
	} else if len(threeMatches) > 0 {
		winners = append(winners, threeMatches...)
		return models.GetWinners(results.GameType, models.PTSS_Prize, 3, winners, prize.Round(2))
	} else if len(twoMatches) > 0 {
		winners = append(winners, twoMatches...)
		return models.GetWinners(results.GameType, models.PTSS_Prize, 2, winners, prize.Round(2))
	} else {
		winners = append(winners, oneMatch...)
		return models.GetWinners(results.GameType, models.PTSS_Prize, 1, winners, prize.Round(2))
	}

}

func CheckPrizesPozoExtraPrize(players []models.Player, results models.Result, prize decimal.Decimal, drawings models.Q6Results) models.Winners {

	//to participate this player must NOT have 6 matches in any of the three main results

	drawing := results.DrawingResults
	var winners []models.Player
	for _, p := range players {
		if p.Quini6Ticket.Participation == models.GP_TradicionalOnly {
			//check if users' 6 numbers are identical to tradicional's drawing and segunda's drawing
			if GetNumberOfMatches(p.Quini6Ticket.Numbers, drawings.GTRT.DrawingResults) == 6 ||
				GetNumberOfMatches(p.Quini6Ticket.Numbers, drawings.GTRS.DrawingResults) == 6 {
				//this user can't play
				continue
			}
		} else if p.Quini6Ticket.Participation == models.GP_TradicionalAndRevancha || p.Quini6Ticket.Participation == models.GP_TradicionalAndRevanchaAndSiempreSale {
			//check if users' 6 numbers are identical to tradicional's drawing, segunda's drawing and revancha's drawing
			if GetNumberOfMatches(p.Quini6Ticket.Numbers, drawings.GTRT.DrawingResults) == 6 ||
				GetNumberOfMatches(p.Quini6Ticket.Numbers, drawings.GTRS.DrawingResults) == 6 ||
				GetNumberOfMatches(p.Quini6Ticket.Numbers, drawings.GTRR.DrawingResults) == 6 {
				//this user can't play
				continue
			}
		}
		numberOfMatches := GetNumberOfMatches(p.Quini6Ticket.Numbers, drawing)
		prizeType := CheckMatchesPozoExtra(numberOfMatches)
		if prizeType == models.PTPE_Prize {
			winners = append(winners, p)
		}
	}
	return models.GetWinners(results.GameType, models.PTPE_Prize, 6, winners, prize.Round(2))
}

func CheckMatches(numberOfMatches int) models.Matches {
	switch numberOfMatches {
	case 6:
		return models.M_SixMatches
	case 5:
		return models.M_FiveMatches
	case 4:
		return models.M_FourMatches
	case 3:
		return models.M_ThreeMatches
	case 2:
		return models.M_TwoMatches
	case 1:
		return models.M_OneMatch
	case 0:
		return models.M_NoMatches
	default:
		return models.M_NoMatches
	}
}

func CheckMatchesTradicional(numberOfMatches int) models.PrizeType {
	tradicionalMatches := CheckMatches(numberOfMatches)
	switch tradicionalMatches {
	case models.M_SixMatches:
		return models.PTT_FirstPrize
	case models.M_FiveMatches:
		return models.PTT_SecondPrize
	case models.M_FourMatches:
		return models.PTT_ThirdPrize
	case models.M_ThreeMatches:
		return models.PTT_NoPrize
	case models.M_TwoMatches:
		return models.PTT_NoPrize
	case models.M_OneMatch:
		return models.PTT_NoPrize
	case models.M_NoMatches:
		return models.PTT_NoPrize
	default:
		return models.PTT_NoPrize
	}
}

func CheckMatchesSegunda(numberOfMatches int) models.PrizeType {
	segundaMatches := CheckMatches(numberOfMatches)
	switch segundaMatches {
	case models.M_SixMatches:
		return models.PTS_FirstPrize
	case models.M_FiveMatches:
		return models.PTS_SecondPrize
	case models.M_FourMatches:
		return models.PTS_ThirdPrize
	case models.M_ThreeMatches:
		return models.PTS_NoPrize
	case models.M_TwoMatches:
		return models.PTS_NoPrize
	case models.M_OneMatch:
		return models.PTS_NoPrize
	case models.M_NoMatches:
		return models.PTS_NoPrize
	default:
		return models.PTS_NoPrize
	}
}

func CheckMatchesRevancha(numberOfMatches int) models.PrizeType {
	revanchaMatches := CheckMatches(numberOfMatches)
	switch revanchaMatches {
	case models.M_SixMatches:
		return models.PTR_Prize
	case models.M_FiveMatches:
		return models.PTR_NoPrize
	case models.M_FourMatches:
		return models.PTR_NoPrize
	case models.M_ThreeMatches:
		return models.PTR_NoPrize
	case models.M_TwoMatches:
		return models.PTR_NoPrize
	case models.M_OneMatch:
		return models.PTR_NoPrize
	case models.M_NoMatches:
		return models.PTR_NoPrize
	default:
		return models.PTR_NoPrize
	}
}

func CheckMatchesSiempreSale(numberOfMatches int) int {
	siempreSaleMatches := CheckMatches(numberOfMatches)
	switch siempreSaleMatches {
	case models.M_SixMatches:
		return 6
	case models.M_FiveMatches:
		return 5
	case models.M_FourMatches:
		return 4
	case models.M_ThreeMatches:
		return 3
	case models.M_TwoMatches:
		return 2
	case models.M_OneMatch:
		return 1
	case models.M_NoMatches:
		return 0
	default:
		return 0
	}
}

func CheckMatchesPozoExtra(numberOfMatches int) models.PrizeType {
	pozoExtraMatches := CheckMatches(numberOfMatches)
	switch pozoExtraMatches {
	case models.M_SixMatches:
		return models.PTPE_Prize
	case models.M_FiveMatches:
		return models.PTPE_NoPrize
	case models.M_FourMatches:
		return models.PTPE_NoPrize
	case models.M_ThreeMatches:
		return models.PTPE_NoPrize
	case models.M_TwoMatches:
		return models.PTPE_NoPrize
	case models.M_OneMatch:
		return models.PTPE_NoPrize
	case models.M_NoMatches:
		return models.PTPE_NoPrize
	default:
		return models.PTPE_NoPrize
	}
}
