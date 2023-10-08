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

func CheckMatchesTradicional(numberOfMatches int) PrizeType {
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

func CheckMatchesSegunda(numberOfMatches int) PrizeType {
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

func CheckMatchesRevancha(numberOfMatches int) PrizeType {
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

func CheckMatchesSiempreSale(numberOfMatches int) int {
	siempreSaleMatches := CheckMatches(numberOfMatches)
	switch siempreSaleMatches {
	case M_SixMatches:
		return 6
	case M_FiveMatches:
		return 5
	case M_FourMatches:
		return 4
	case M_ThreeMatches:
		return 3
	case M_TwoMatches:
		return 2
	case M_OneMatch:
		return 1
	case M_NoMatches:
		return 0
	default:
		return 0
	}
}

func CheckMatchesPozoExtra(numberOfMatches int) PrizeType {
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
