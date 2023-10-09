package utils

import (
	"q6-golang/models"

	"github.com/shopspring/decimal"
)

func CheckPrizesTradicionalFirstPrize(ch chan models.Winners, players []models.Player, results models.Result, prize decimal.Decimal) {
	drawing := results.DrawingResults
	var winners []models.Player
	for _, p := range players {
		numberOfMatches := GetNumberOfMatches(p.Quini6Ticket.Numbers, drawing)
		prizeType := CheckPrizeTradicional(numberOfMatches)
		if prizeType == models.PTT_FirstPrize {
			winners = append(winners, p)
		}
	}
	ch <- models.GetWinners(results.GameType, models.PTT_FirstPrize, 6, winners, prize.Round(2))
}

func CheckPrizesTradicionalSecondPrize(ch chan models.Winners, players []models.Player, results models.Result, prize decimal.Decimal) {
	drawing := results.DrawingResults
	var winners []models.Player
	for _, p := range players {
		numberOfMatches := GetNumberOfMatches(p.Quini6Ticket.Numbers, drawing)
		prizeType := CheckPrizeTradicional(numberOfMatches)
		if prizeType == models.PTT_SecondPrize {
			winners = append(winners, p)
		}
	}
	ch <- models.GetWinners(results.GameType, models.PTT_SecondPrize, 5, winners, prize.Round(2))
}

func CheckPrizesTradicionalThirdPrize(ch chan models.Winners, players []models.Player, results models.Result, prize decimal.Decimal) {
	drawing := results.DrawingResults
	var winners []models.Player
	for _, p := range players {
		numberOfMatches := GetNumberOfMatches(p.Quini6Ticket.Numbers, drawing)
		prizeType := CheckPrizeTradicional(numberOfMatches)
		if prizeType == models.PTT_ThirdPrize {
			winners = append(winners, p)
		}
	}
	ch <- models.GetWinners(results.GameType, models.PTT_ThirdPrize, 4, winners, prize.Round(2))
}

func CheckPrizesSegundaFirstPrize(ch chan models.Winners, players []models.Player, results models.Result, prize decimal.Decimal) {
	drawing := results.DrawingResults
	var winners []models.Player
	for _, p := range players {
		numberOfMatches := GetNumberOfMatches(p.Quini6Ticket.Numbers, drawing)
		prizeType := CheckPrizeSegunda(numberOfMatches)
		if prizeType == models.PTS_FirstPrize {
			winners = append(winners, p)
		}
	}
	ch <- models.GetWinners(results.GameType, models.PTS_FirstPrize, 6, winners, prize.Round(2))
}

func CheckPrizesSegundaSecondPrize(ch chan models.Winners, players []models.Player, results models.Result, prize decimal.Decimal) {
	drawing := results.DrawingResults
	var winners []models.Player
	for _, p := range players {
		numberOfMatches := GetNumberOfMatches(p.Quini6Ticket.Numbers, drawing)
		prizeType := CheckPrizeSegunda(numberOfMatches)
		if prizeType == models.PTS_SecondPrize {
			winners = append(winners, p)
		}
	}
	ch <- models.GetWinners(results.GameType, models.PTS_SecondPrize, 5, winners, prize.Round(2))
}

func CheckPrizesSegundaThirdPrize(ch chan models.Winners, players []models.Player, results models.Result, prize decimal.Decimal) {
	drawing := results.DrawingResults
	var winners []models.Player
	for _, p := range players {
		numberOfMatches := GetNumberOfMatches(p.Quini6Ticket.Numbers, drawing)
		prizeType := CheckPrizeSegunda(numberOfMatches)
		if prizeType == models.PTS_ThirdPrize {
			winners = append(winners, p)
		}
	}
	ch <- models.GetWinners(results.GameType, models.PTS_ThirdPrize, 4, winners, prize.Round(2))
}

func CheckPrizesRevanchaPrize(ch chan models.Winners, players []models.Player, results models.Result, prize decimal.Decimal) {
	drawing := results.DrawingResults
	var winners []models.Player
	for _, p := range players {
		if p.Quini6Ticket.Participation == models.GP_TradicionalAndRevancha || p.Quini6Ticket.Participation == models.GP_TradicionalAndRevanchaAndSiempreSale {
			numberOfMatches := GetNumberOfMatches(p.Quini6Ticket.Numbers, drawing)
			prizeType := CheckPrizeRevancha(numberOfMatches)
			if prizeType == models.PTR_Prize {
				winners = append(winners, p)
			}
		}
	}
	ch <- models.GetWinners(results.GameType, models.PTR_Prize, 6, winners, prize.Round(2))
}

func CheckPrizesSiempreSalePrize(ch chan models.Winners, players []models.Player, results models.Result, prize decimal.Decimal) {
	drawing := results.DrawingResults
	var sixMatches []models.Player
	var fiveMatches []models.Player
	var fourMatches []models.Player
	var threeMatches []models.Player
	var twoMatches []models.Player
	var oneMatch []models.Player
	var winners []models.Player
	skip := 0
	for _, p := range players {
		if p.Quini6Ticket.Participation == models.GP_TradicionalAndRevancha || p.Quini6Ticket.Participation == models.GP_TradicionalAndRevanchaAndSiempreSale {
			numberOfMatches := GetNumberOfMatches(p.Quini6Ticket.Numbers, drawing)
			if numberOfMatches == 6 {
				sixMatches = append(winners, p)
				skip = 6
			} else if numberOfMatches == 5 && skip <= 5 {
				fiveMatches = append(winners, p)
				skip = 5
			} else if numberOfMatches == 4 && skip <= 4 {
				fourMatches = append(winners, p)
				skip = 4
			} else if numberOfMatches == 3 && skip <= 3 {
				threeMatches = append(winners, p)
				skip = 3
			} else if numberOfMatches == 2 && skip <= 2 {
				twoMatches = append(winners, p)
				skip = 2
			} else if numberOfMatches == 1 && skip <= 1 {
				oneMatch = append(winners, p)
				skip = 1
			}
		}
	}
	if len(sixMatches) > 0 {
		ch <- models.GetWinners(results.GameType, models.PTSS_Prize, 6, sixMatches, prize.Round(2))
	} else if len(fiveMatches) > 0 {
		ch <- models.GetWinners(results.GameType, models.PTSS_Prize, 5, fiveMatches, prize.Round(2))
	} else if len(fourMatches) > 0 {
		ch <- models.GetWinners(results.GameType, models.PTSS_Prize, 4, fourMatches, prize.Round(2))
	} else if len(threeMatches) > 0 {
		ch <- models.GetWinners(results.GameType, models.PTSS_Prize, 3, threeMatches, prize.Round(2))
	} else if len(twoMatches) > 0 {
		ch <- models.GetWinners(results.GameType, models.PTSS_Prize, 2, twoMatches, prize.Round(2))
	} else {
		ch <- models.GetWinners(results.GameType, models.PTSS_Prize, 1, oneMatch, prize.Round(2))
	}
}

func CheckPrizesPozoExtraPrize(ch chan models.Winners, players []models.Player, results models.Result, prize decimal.Decimal, drawings models.Q6Results) {

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
		prizeType := CheckPrizePozoExtra(numberOfMatches)
		if prizeType == models.PTPE_Prize {
			winners = append(winners, p)
		}
	}
	ch <- models.GetWinners(results.GameType, models.PTPE_Prize, 6, winners, prize.Round(2))
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

func CheckPrizeTradicional(numberOfMatches int) models.PrizeType {
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

func CheckPrizeSegunda(numberOfMatches int) models.PrizeType {
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

func CheckPrizeRevancha(numberOfMatches int) models.PrizeType {
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

func CheckPrizePozoExtra(numberOfMatches int) models.PrizeType {
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
