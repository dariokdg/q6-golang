package core

import (
	e "q6-golang/enumerators"
)

func CheckMatches(numberOfMatches int) e.Matches {
	switch numberOfMatches {
	case 6:
		return e.M_SixMatches
	case 5:
		return e.M_FiveMatches
	case 4:
		return e.M_FourMatches
	case 3:
		return e.M_ThreeMatches
	case 2:
		return e.M_TwoMatches
	case 1:
		return e.M_OneMatch
	case 0:
		return e.M_NoMatches
	default:
		return e.M_NoMatches
	}
}

func CheckMatchesTradicional(numberOfMatches int) e.PrizeTypeTradicional {
	tradicionalMatches := CheckMatches(numberOfMatches)
	switch tradicionalMatches {
	case e.M_SixMatches:
		return e.PTT_FirstPrize
	case e.M_FiveMatches:
		return e.PTT_SecondPrize
	case e.M_FourMatches:
		return e.PTT_ThirdPrize
	case e.M_ThreeMatches:
		return e.PTT_NoPrize
	case e.M_TwoMatches:
		return e.PTT_NoPrize
	case e.M_OneMatch:
		return e.PTT_NoPrize
	case e.M_NoMatches:
		return e.PTT_NoPrize
	default:
		return e.PTT_NoPrize
	}
}

func CheckMatchesSegunda(numberOfMatches int) e.PrizeTypeSegunda {
	segundaMatches := CheckMatches(numberOfMatches)
	switch segundaMatches {
	case e.M_SixMatches:
		return e.PTS_FirstPrize
	case e.M_FiveMatches:
		return e.PTS_SecondPrize
	case e.M_FourMatches:
		return e.PTS_ThirdPrize
	case e.M_ThreeMatches:
		return e.PTS_NoPrize
	case e.M_TwoMatches:
		return e.PTS_NoPrize
	case e.M_OneMatch:
		return e.PTS_NoPrize
	case e.M_NoMatches:
		return e.PTS_NoPrize
	default:
		return e.PTS_NoPrize
	}
}

func CheckMatchesRevancha(numberOfMatches int) e.PrizeTypeRevancha {
	revanchaMatches := CheckMatches(numberOfMatches)
	switch revanchaMatches {
	case e.M_SixMatches:
		return e.PTR_Prize
	case e.M_FiveMatches:
		return e.PTR_NoPrize
	case e.M_FourMatches:
		return e.PTR_NoPrize
	case e.M_ThreeMatches:
		return e.PTR_NoPrize
	case e.M_TwoMatches:
		return e.PTR_NoPrize
	case e.M_OneMatch:
		return e.PTR_NoPrize
	case e.M_NoMatches:
		return e.PTR_NoPrize
	default:
		return e.PTR_NoPrize
	}
}

func CheckMatchesSiempreSale(numberOfMatches int) e.PrizeTypeSiempreSale {
	siempreSaleMatches := CheckMatches(numberOfMatches)
	switch siempreSaleMatches {
	case e.M_SixMatches:
		return e.PTSS_PotentialWinnerSixMatches
	case e.M_FiveMatches:
		return e.PTSS_PotentialWinnerFiveMatches
	case e.M_FourMatches:
		return e.PTSS_PotentialWinnerFourMatches
	case e.M_ThreeMatches:
		return e.PTSS_PotentialWinnerThreeMatches
	case e.M_TwoMatches:
		return e.PTSS_PotentialWinnerTwoMatches
	case e.M_OneMatch:
		return e.PTSS_PotentialWinnerOneMatch
	case e.M_NoMatches:
		return e.PTSS_NoPrize
	default:
		return e.PTSS_NoPrize
	}
}

func CheckMatchesPozoExtra(numberOfMatches int) e.PrizeTypePozoExtra {
	pozoExtraMatches := CheckMatches(numberOfMatches)
	switch pozoExtraMatches {
	case e.M_SixMatches:
		return e.PTPE_Prize
	case e.M_FiveMatches:
		return e.PTPE_NoPrize
	case e.M_FourMatches:
		return e.PTPE_NoPrize
	case e.M_ThreeMatches:
		return e.PTPE_NoPrize
	case e.M_TwoMatches:
		return e.PTPE_NoPrize
	case e.M_OneMatch:
		return e.PTPE_NoPrize
	case e.M_NoMatches:
		return e.PTPE_NoPrize
	default:
		return e.PTPE_NoPrize
	}
}
