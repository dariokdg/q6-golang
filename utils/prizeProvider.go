package utils

import "q6-golang/models"

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
