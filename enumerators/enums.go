package enumerators

type GameType int64
type Matches int64
type PrizeTypeTradicional int64
type PrizeTypeSegunda int64
type PrizeTypeRevancha int64
type PrizeTypeSiempreSale int64
type PrizeTypePozoExtra int64
type GameParticipation int64

const (
	GT_Tradicional GameType = iota
	GT_Segunda
	GT_Revancha
	GT_SiempreSale
	GT_PozoExtra
)

const (
	M_NoMatches Matches = iota
	M_OneMatch
	M_TwoMatches
	M_ThreeMatches
	M_FourMatches
	M_FiveMatches
	M_SixMatches
)

const (
	PTT_NoPrize PrizeTypeTradicional = iota
	PTT_FirstPrize
	PTT_SecondPrize
	PTT_ThirdPrize
)

const (
	PTS_NoPrize PrizeTypeSegunda = iota
	PTS_FirstPrize
	PTS_SecondPrize
	PTS_ThirdPrize
)

const (
	PTR_NoPrize PrizeTypeRevancha = iota
	PTR_Prize
)

const (
	PTSS_NoPrize PrizeTypeSiempreSale = iota
	PTSS_PotentialWinnerSixMatches
	PTSS_PotentialWinnerFiveMatches
	PTSS_PotentialWinnerFourMatches
	PTSS_PotentialWinnerThreeMatches
	PTSS_PotentialWinnerTwoMatches
	PTSS_PotentialWinnerOneMatch
)

const (
	PTPE_NoPrize PrizeTypePozoExtra = iota
	PTPE_Prize
)

const (
	GP_TradicionalOnly GameParticipation = iota
	GP_TradicionalAndRevancha
	GP_TradicionalAndRevanchaAndSiempreSale
)
