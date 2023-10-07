package core

func CheckMatches(numberOfMatches int) Matches {
	switch numberOfMatches {
	case 6:
		return M_SixMatches
	case 5:
		return M_FiveMatches
	case 4:
		return M_FourMatches
	case 3:
		return M_ThreeMatches
	case 2:
		return M_TwoMatches
	case 1:
		return M_OneMatch
	case 0:
		return M_NoMatches
	default:
		return M_NoMatches
	}
}

func CheckMatchesTradicional(numberOfMatches int) PrizeTypeTradicional {
	tradicionalMatches := CheckMatches(numberOfMatches)
	switch tradicionalMatches {
	case M_SixMatches:
		return PTT_FirstPrize
	case M_FiveMatches:
		return PTT_SecondPrize
	case M_FourMatches:
		return PTT_ThirdPrize
	case M_ThreeMatches:
		return PTT_NoPrize
	case M_TwoMatches:
		return PTT_NoPrize
	case M_OneMatch:
		return PTT_NoPrize
	case M_NoMatches:
		return PTT_NoPrize
	default:
		return PTT_NoPrize
	}
}

func CheckMatchesSegunda(numberOfMatches int) PrizeTypeSegunda {
	segundaMatches := CheckMatches(numberOfMatches)
	switch segundaMatches {
	case M_SixMatches:
		return PTS_FirstPrize
	case M_FiveMatches:
		return PTS_SecondPrize
	case M_FourMatches:
		return PTS_ThirdPrize
	case M_ThreeMatches:
		return PTS_NoPrize
	case M_TwoMatches:
		return PTS_NoPrize
	case M_OneMatch:
		return PTS_NoPrize
	case M_NoMatches:
		return PTS_NoPrize
	default:
		return PTS_NoPrize
	}
}

func CheckMatchesRevancha(numberOfMatches int) PrizeTypeRevancha {
	revanchaMatches := CheckMatches(numberOfMatches)
	switch revanchaMatches {
	case M_SixMatches:
		return PTR_Prize
	case M_FiveMatches:
		return PTR_NoPrize
	case M_FourMatches:
		return PTR_NoPrize
	case M_ThreeMatches:
		return PTR_NoPrize
	case M_TwoMatches:
		return PTR_NoPrize
	case M_OneMatch:
		return PTR_NoPrize
	case M_NoMatches:
		return PTR_NoPrize
	default:
		return PTR_NoPrize
	}
}

func CheckMatchesSiempreSale(numberOfMatches int) PrizeTypeSiempreSale {
	siempreSaleMatches := CheckMatches(numberOfMatches)
	switch siempreSaleMatches {
	case M_SixMatches:
		return PTSS_PotentialWinnerSixMatches
	case M_FiveMatches:
		return PTSS_PotentialWinnerFiveMatches
	case M_FourMatches:
		return PTSS_PotentialWinnerFourMatches
	case M_ThreeMatches:
		return PTSS_PotentialWinnerThreeMatches
	case M_TwoMatches:
		return PTSS_PotentialWinnerTwoMatches
	case M_OneMatch:
		return PTSS_PotentialWinnerOneMatch
	case M_NoMatches:
		return PTSS_NoPrize
	default:
		return PTSS_NoPrize
	}
}

func CheckMatchesPozoExtra(numberOfMatches int) PrizeTypePozoExtra {
	pozoExtraMatches := CheckMatches(numberOfMatches)
	switch pozoExtraMatches {
	case M_SixMatches:
		return PTPE_Prize
	case M_FiveMatches:
		return PTPE_NoPrize
	case M_FourMatches:
		return PTPE_NoPrize
	case M_ThreeMatches:
		return PTPE_NoPrize
	case M_TwoMatches:
		return PTPE_NoPrize
	case M_OneMatch:
		return PTPE_NoPrize
	case M_NoMatches:
		return PTPE_NoPrize
	default:
		return PTPE_NoPrize
	}
}
