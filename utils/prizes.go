package utils

import (
	"q6-golang/models"

	"github.com/shopspring/decimal"
)

func CheckPrizesTradicionalFirstPrize(ch chan models.Winners, players []models.Player, results models.Result, prize decimal.Decimal) {
	drawing := results.DrawingResults
	var winners []models.Ticket
	for _, p := range players {
		for _, t := range p.Tickets {
			numberOfMatches := GetNumberOfMatches(t.Numbers, drawing)
			prizeType := CheckPrizeTradicional(numberOfMatches)
			if prizeType == models.PTT_FirstPrize {
				winners = append(winners, t)
			}
		}
	}
	ch <- GetWinners(results.GameType, models.PTT_FirstPrize, 6, players, winners, prize.Round(2))
}

func CheckPrizesTradicionalSecondPrize(ch chan models.Winners, players []models.Player, results models.Result, prize decimal.Decimal) {
	drawing := results.DrawingResults
	var winners []models.Ticket
	for _, p := range players {
		for _, t := range p.Tickets {
			numberOfMatches := GetNumberOfMatches(t.Numbers, drawing)
			prizeType := CheckPrizeTradicional(numberOfMatches)
			if prizeType == models.PTT_SecondPrize {
				winners = append(winners, t)
			}
		}
	}
	ch <- GetWinners(results.GameType, models.PTT_SecondPrize, 5, players, winners, prize.Round(2))
}

func CheckPrizesTradicionalThirdPrize(ch chan models.Winners, players []models.Player, results models.Result, prize decimal.Decimal) {
	drawing := results.DrawingResults
	var winners []models.Ticket
	for _, p := range players {
		for _, t := range p.Tickets {
			numberOfMatches := GetNumberOfMatches(t.Numbers, drawing)
			prizeType := CheckPrizeTradicional(numberOfMatches)
			if prizeType == models.PTT_ThirdPrize {
				winners = append(winners, t)
			}
		}
	}
	ch <- GetWinners(results.GameType, models.PTT_ThirdPrize, 4, players, winners, prize.Round(2))
}

func CheckPrizesSegundaFirstPrize(ch chan models.Winners, players []models.Player, results models.Result, prize decimal.Decimal) {
	drawing := results.DrawingResults
	var winners []models.Ticket
	for _, p := range players {
		for _, t := range p.Tickets {
			numberOfMatches := GetNumberOfMatches(t.Numbers, drawing)
			prizeType := CheckPrizeSegunda(numberOfMatches)
			if prizeType == models.PTS_FirstPrize {
				winners = append(winners, t)
			}
		}
	}
	ch <- GetWinners(results.GameType, models.PTS_FirstPrize, 6, players, winners, prize.Round(2))
}

func CheckPrizesSegundaSecondPrize(ch chan models.Winners, players []models.Player, results models.Result, prize decimal.Decimal) {
	drawing := results.DrawingResults
	var winners []models.Ticket
	for _, p := range players {
		for _, t := range p.Tickets {
			numberOfMatches := GetNumberOfMatches(t.Numbers, drawing)
			prizeType := CheckPrizeSegunda(numberOfMatches)
			if prizeType == models.PTS_SecondPrize {
				winners = append(winners, t)
			}
		}
	}
	ch <- GetWinners(results.GameType, models.PTS_SecondPrize, 5, players, winners, prize.Round(2))
}

func CheckPrizesSegundaThirdPrize(ch chan models.Winners, players []models.Player, results models.Result, prize decimal.Decimal) {
	drawing := results.DrawingResults
	var winners []models.Ticket
	for _, p := range players {
		for _, t := range p.Tickets {
			numberOfMatches := GetNumberOfMatches(t.Numbers, drawing)
			prizeType := CheckPrizeSegunda(numberOfMatches)
			if prizeType == models.PTS_ThirdPrize {
				winners = append(winners, t)
			}
		}
	}
	ch <- GetWinners(results.GameType, models.PTS_ThirdPrize, 4, players, winners, prize.Round(2))
}

func CheckPrizesRevanchaPrize(ch chan models.Winners, players []models.Player, results models.Result, prize decimal.Decimal) {
	drawing := results.DrawingResults
	var winners []models.Ticket
	for _, p := range players {
		for _, t := range p.Tickets {
			if t.Participation == models.GP_TradicionalAndRevancha || t.Participation == models.GP_TradicionalAndRevanchaAndSiempreSale {
				numberOfMatches := GetNumberOfMatches(t.Numbers, drawing)
				prizeType := CheckPrizeRevancha(numberOfMatches)
				if prizeType == models.PTR_Prize {
					winners = append(winners, t)
				}
			}
		}
	}
	ch <- GetWinners(results.GameType, models.PTR_Prize, 6, players, winners, prize.Round(2))
}

func CheckPrizesSiempreSalePrize(ch chan models.Winners, players []models.Player, results models.Result, prize decimal.Decimal) {
	drawing := results.DrawingResults
	var sixMatches []models.Ticket
	var fiveMatches []models.Ticket
	var fourMatches []models.Ticket
	var threeMatches []models.Ticket
	var twoMatches []models.Ticket
	var oneMatch []models.Ticket
	var winners []models.Ticket
	skip := 0
	for _, p := range players {
		for _, t := range p.Tickets {
			if t.Participation == models.GP_TradicionalAndRevanchaAndSiempreSale {
				numberOfMatches := GetNumberOfMatches(t.Numbers, drawing)
				if numberOfMatches == 6 {
					sixMatches = append(winners, t)
					skip = 6
				} else if numberOfMatches == 5 && skip <= 5 {
					fiveMatches = append(winners, t)
					skip = 5
				} else if numberOfMatches == 4 && skip <= 4 {
					fourMatches = append(winners, t)
					skip = 4
				} else if numberOfMatches == 3 && skip <= 3 {
					threeMatches = append(winners, t)
					skip = 3
				} else if numberOfMatches == 2 && skip <= 2 {
					twoMatches = append(winners, t)
					skip = 2
				} else if numberOfMatches == 1 && skip <= 1 {
					oneMatch = append(winners, t)
					skip = 1
				}
			}
		}
	}
	if len(sixMatches) > 0 {
		ch <- GetWinners(results.GameType, models.PTSS_Prize, 6, players, sixMatches, prize.Round(2))
	} else if len(fiveMatches) > 0 {
		ch <- GetWinners(results.GameType, models.PTSS_Prize, 5, players, fiveMatches, prize.Round(2))
	} else if len(fourMatches) > 0 {
		ch <- GetWinners(results.GameType, models.PTSS_Prize, 4, players, fourMatches, prize.Round(2))
	} else if len(threeMatches) > 0 {
		ch <- GetWinners(results.GameType, models.PTSS_Prize, 3, players, threeMatches, prize.Round(2))
	} else if len(twoMatches) > 0 {
		ch <- GetWinners(results.GameType, models.PTSS_Prize, 2, players, twoMatches, prize.Round(2))
	} else if len(oneMatch) > 0 {
		ch <- GetWinners(results.GameType, models.PTSS_Prize, 1, players, oneMatch, prize.Round(2))
	} else {
		ch <- GetWinners(results.GameType, models.PTSS_Prize, 0, players, []models.Ticket{}, prize.Round(2))
	}
}

func CheckPrizesPozoExtraPrize(ch chan models.Winners, players []models.Player, results models.Result, prize decimal.Decimal, drawings models.Q6Results) {

	//to participate this player must NOT have 6 matches in any of the three main results

	drawing := results.DrawingResults
	var winners []models.Ticket
	for _, p := range players {
		for _, t := range p.Tickets {
			if t.Participation == models.GP_TradicionalOnly {
				//check if tickets' 6 numbers are identical to tradicional's drawing and segunda's drawing
				if GetNumberOfMatches(t.Numbers, drawings.GTRT.DrawingResults) == 6 ||
					GetNumberOfMatches(t.Numbers, drawings.GTRS.DrawingResults) == 6 {
					//this ticket can't play 'Pozo Extra'
					continue
				}
			} else if t.Participation == models.GP_TradicionalAndRevancha || t.Participation == models.GP_TradicionalAndRevanchaAndSiempreSale {
				//check if tickets' 6 numbers are identical to tradicional's drawing, segunda's drawing and revancha's drawing
				if GetNumberOfMatches(t.Numbers, drawings.GTRT.DrawingResults) == 6 ||
					GetNumberOfMatches(t.Numbers, drawings.GTRS.DrawingResults) == 6 ||
					GetNumberOfMatches(t.Numbers, drawings.GTRR.DrawingResults) == 6 {
					//this ticket can't play 'Pozo Extra'
					continue
				}
			}
			numberOfMatches := GetNumberOfMatches(t.Numbers, drawing)
			prizeType := CheckPrizePozoExtra(numberOfMatches)
			if prizeType == models.PTPE_Prize {
				winners = append(winners, t)
			}
		}
	}
	ch <- GetWinners(results.GameType, models.PTPE_Prize, 6, players, winners, prize.Round(2))
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
