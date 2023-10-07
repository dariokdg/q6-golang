package core

import (
	"q6-golang/enumerators"

	"github.com/shopspring/decimal"
)

func CheckPrizesTradicionalFirstPrize(results GameTypeResult, prize decimal.Decimal) Winner {
	drawing := results.DrawingResults
	players := results.Players
	var winners []Player
	for _, p := range players {
		numberOfMatches := GetNumberOfMatches(p.Quini6Ticket.SelectedNumbers, drawing)
		prizeType := CheckMatchesTradicional(numberOfMatches)
		if prizeType == enumerators.PTT_FirstPrize {
			winners = append(winners, p)
		}
	}
	return GetWinner(results.GameType, winners, prize)
}

func CheckPrizesTradicionalSecondPrize(results GameTypeResult, prize decimal.Decimal) Winner {
	drawing := results.DrawingResults
	players := results.Players
	var winners []Player
	for _, p := range players {
		numberOfMatches := GetNumberOfMatches(p.Quini6Ticket.SelectedNumbers, drawing)
		prizeType := CheckMatchesTradicional(numberOfMatches)
		if prizeType == enumerators.PTT_SecondPrize {
			winners = append(winners, p)
		}
	}
	return GetWinner(results.GameType, winners, prize)
}

func CheckPrizesTradicionalThirdPrize(results GameTypeResult, prize decimal.Decimal) Winner {
	drawing := results.DrawingResults
	players := results.Players
	var winners []Player
	for _, p := range players {
		numberOfMatches := GetNumberOfMatches(p.Quini6Ticket.SelectedNumbers, drawing)
		prizeType := CheckMatchesTradicional(numberOfMatches)
		if prizeType == enumerators.PTT_ThirdPrize {
			winners = append(winners, p)
		}
	}
	return GetWinner(results.GameType, winners, prize)
}

func CheckPrizesSegundaFirstPrize(results GameTypeResult, prize decimal.Decimal) Winner {
	drawing := results.DrawingResults
	players := results.Players
	var winners []Player
	for _, p := range players {
		numberOfMatches := GetNumberOfMatches(p.Quini6Ticket.SelectedNumbers, drawing)
		prizeType := CheckMatchesSegunda(numberOfMatches)
		if prizeType == enumerators.PTS_FirstPrize {
			winners = append(winners, p)
		}
	}
	return GetWinner(results.GameType, winners, prize)
}

func CheckPrizesSegundaSecondPrize(results GameTypeResult, prize decimal.Decimal) Winner {
	drawing := results.DrawingResults
	players := results.Players
	var winners []Player
	for _, p := range players {
		numberOfMatches := GetNumberOfMatches(p.Quini6Ticket.SelectedNumbers, drawing)
		prizeType := CheckMatchesSegunda(numberOfMatches)
		if prizeType == enumerators.PTS_SecondPrize {
			winners = append(winners, p)
		}
	}
	return GetWinner(results.GameType, winners, prize)
}

func CheckPrizesSegundaThirdPrize(results GameTypeResult, prize decimal.Decimal) Winner {
	drawing := results.DrawingResults
	players := results.Players
	var winners []Player
	for _, p := range players {
		numberOfMatches := GetNumberOfMatches(p.Quini6Ticket.SelectedNumbers, drawing)
		prizeType := CheckMatchesSegunda(numberOfMatches)
		if prizeType == enumerators.PTS_ThirdPrize {
			winners = append(winners, p)
		}
	}
	return GetWinner(results.GameType, winners, prize)
}

func CheckPrizesRevanchaPrize(results GameTypeResult, prize decimal.Decimal) Winner {
	drawing := results.DrawingResults
	players := results.Players
	var winners []Player
	for _, p := range players {
		numberOfMatches := GetNumberOfMatches(p.Quini6Ticket.SelectedNumbers, drawing)
		prizeType := CheckMatchesRevancha(numberOfMatches)
		if prizeType == enumerators.PTR_Prize {
			winners = append(winners, p)
		}
	}
	return GetWinner(results.GameType, winners, prize)
}

func CheckPrizesSiempreSalePrize(results GameTypeResult, prize decimal.Decimal) Winner {
	drawing := results.DrawingResults
	players := results.Players
	var sixMatches []Player
	var fiveMatches []Player
	var fourMatches []Player
	var threeMatches []Player
	var twoMatches []Player
	var oneMatch []Player
	var winners []Player
	for _, p := range players {
		numberOfMatches := GetNumberOfMatches(p.Quini6Ticket.SelectedNumbers, drawing)
		prizeType := CheckMatchesSiempreSale(numberOfMatches)
		if prizeType == enumerators.PTSS_PotentialWinnerSixMatches {
			sixMatches = append(winners, p)
		} else if prizeType == enumerators.PTSS_PotentialWinnerFiveMatches {
			fiveMatches = append(winners, p)
		} else if prizeType == enumerators.PTSS_PotentialWinnerFourMatches {
			fourMatches = append(winners, p)
		} else if prizeType == enumerators.PTSS_PotentialWinnerThreeMatches {
			threeMatches = append(winners, p)
		} else if prizeType == enumerators.PTSS_PotentialWinnerTwoMatches {
			twoMatches = append(winners, p)
		} else if prizeType == enumerators.PTSS_PotentialWinnerOneMatch {
			oneMatch = append(winners, p)
		}
	}
	if len(sixMatches) > 0 {
		winners = append(winners, sixMatches...)
	} else if len(fiveMatches) > 0 {
		winners = append(winners, fiveMatches...)
	} else if len(fourMatches) > 0 {
		winners = append(winners, fourMatches...)
	} else if len(threeMatches) > 0 {
		winners = append(winners, threeMatches...)
	} else if len(twoMatches) > 0 {
		winners = append(winners, twoMatches...)
	} else if len(oneMatch) > 0 {
		winners = append(winners, oneMatch...)
	}

	return GetWinner(results.GameType, winners, prize)
}

func CheckPrizesPozoExtraPrize(results GameTypeResult, prize decimal.Decimal, drawings GameResults) Winner {

	//to participate this player must NOT have 6 matches in any of the three main results

	drawing := results.DrawingResults
	players := results.Players
	var winners []Player
	for _, p := range players {
		if p.Quini6Ticket.Games == enumerators.GP_TradicionalOnly {
			//check if users' 6 numbers are identical to tradicional's drawing and segunda's drawing
			if GetNumberOfMatches(p.Quini6Ticket.SelectedNumbers, drawings.GTRT.DrawingResults) == 6 ||
				GetNumberOfMatches(p.Quini6Ticket.SelectedNumbers, drawings.GTRS.DrawingResults) == 6 {
				//this user can't play
				continue
			}
		} else if p.Quini6Ticket.Games == enumerators.GP_TradicionalAndRevancha || p.Quini6Ticket.Games == enumerators.GP_TradicionalAndRevanchaAndSiempreSale {
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
		if prizeType == enumerators.PTPE_Prize {
			winners = append(winners, p)
		}
	}
	return GetWinner(results.GameType, winners, prize)
}
