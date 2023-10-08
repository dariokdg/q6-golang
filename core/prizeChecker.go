package core

import "github.com/shopspring/decimal"

func CheckPrizesTradicionalFirstPrize(players []Player, results GameTypeResult, prize decimal.Decimal) Winner {
	drawing := results.DrawingResults
	var winners []Player
	for _, p := range players {
		numberOfMatches := GetNumberOfMatches(p.Quini6Ticket.SelectedNumbers, drawing)
		prizeType := CheckMatchesTradicional(numberOfMatches)
		if prizeType == PTT_FirstPrize {
			winners = append(winners, p)
		}
	}
	return GetWinners(results.GameType, PTT_FirstPrize, winners, prize)
}

func CheckPrizesTradicionalSecondPrize(players []Player, results GameTypeResult, prize decimal.Decimal) Winner {
	drawing := results.DrawingResults
	var winners []Player
	for _, p := range players {
		numberOfMatches := GetNumberOfMatches(p.Quini6Ticket.SelectedNumbers, drawing)
		prizeType := CheckMatchesTradicional(numberOfMatches)
		if prizeType == PTT_SecondPrize {
			winners = append(winners, p)
		}
	}
	return GetWinners(results.GameType, PTT_SecondPrize, winners, prize)
}

func CheckPrizesTradicionalThirdPrize(players []Player, results GameTypeResult, prize decimal.Decimal) Winner {
	drawing := results.DrawingResults
	var winners []Player
	for _, p := range players {
		numberOfMatches := GetNumberOfMatches(p.Quini6Ticket.SelectedNumbers, drawing)
		prizeType := CheckMatchesTradicional(numberOfMatches)
		if prizeType == PTT_ThirdPrize {
			winners = append(winners, p)
		}
	}
	return GetWinners(results.GameType, PTT_ThirdPrize, winners, prize)
}

func CheckPrizesSegundaFirstPrize(players []Player, results GameTypeResult, prize decimal.Decimal) Winner {
	drawing := results.DrawingResults
	var winners []Player
	for _, p := range players {
		numberOfMatches := GetNumberOfMatches(p.Quini6Ticket.SelectedNumbers, drawing)
		prizeType := CheckMatchesSegunda(numberOfMatches)
		if prizeType == PTS_FirstPrize {
			winners = append(winners, p)
		}
	}
	return GetWinners(results.GameType, PTS_FirstPrize, winners, prize)
}

func CheckPrizesSegundaSecondPrize(players []Player, results GameTypeResult, prize decimal.Decimal) Winner {
	drawing := results.DrawingResults
	var winners []Player
	for _, p := range players {
		numberOfMatches := GetNumberOfMatches(p.Quini6Ticket.SelectedNumbers, drawing)
		prizeType := CheckMatchesSegunda(numberOfMatches)
		if prizeType == PTS_SecondPrize {
			winners = append(winners, p)
		}
	}
	return GetWinners(results.GameType, PTS_SecondPrize, winners, prize)
}

func CheckPrizesSegundaThirdPrize(players []Player, results GameTypeResult, prize decimal.Decimal) Winner {
	drawing := results.DrawingResults
	var winners []Player
	for _, p := range players {
		numberOfMatches := GetNumberOfMatches(p.Quini6Ticket.SelectedNumbers, drawing)
		prizeType := CheckMatchesSegunda(numberOfMatches)
		if prizeType == PTS_ThirdPrize {
			winners = append(winners, p)
		}
	}
	return GetWinners(results.GameType, PTS_ThirdPrize, winners, prize)
}

func CheckPrizesRevanchaPrize(players []Player, results GameTypeResult, prize decimal.Decimal) Winner {
	drawing := results.DrawingResults
	var winners []Player
	for _, p := range players {
		numberOfMatches := GetNumberOfMatches(p.Quini6Ticket.SelectedNumbers, drawing)
		prizeType := CheckMatchesRevancha(numberOfMatches)
		if prizeType == PTR_Prize {
			winners = append(winners, p)
		}
	}
	return GetWinners(results.GameType, PTR_Prize, winners, prize)
}

func CheckPrizesSiempreSalePrize(players []Player, results GameTypeResult, prize decimal.Decimal) Winner {
	drawing := results.DrawingResults
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

	return GetWinners(results.GameType, PTSS_Prize, winners, prize)
}

func CheckPrizesPozoExtraPrize(players []Player, results GameTypeResult, prize decimal.Decimal, drawings GameResults) Winner {

	//to participate this player must NOT have 6 matches in any of the three main results

	drawing := results.DrawingResults
	var winners []Player
	for _, p := range players {
		if p.Quini6Ticket.Games == GP_TradicionalOnly {
			//check if users' 6 numbers are identical to tradicional's drawing and segunda's drawing
			if GetNumberOfMatches(p.Quini6Ticket.SelectedNumbers, drawings.GTRT.DrawingResults) == 6 ||
				GetNumberOfMatches(p.Quini6Ticket.SelectedNumbers, drawings.GTRS.DrawingResults) == 6 {
				//this user can't play
				continue
			}
		} else if p.Quini6Ticket.Games == GP_TradicionalAndRevancha || p.Quini6Ticket.Games == GP_TradicionalAndRevanchaAndSiempreSale {
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
		if prizeType == PTPE_Prize {
			winners = append(winners, p)
		}
	}
	return GetWinners(results.GameType, PTPE_Prize, winners, prize)
}
