package models

type Matches int64
type GameType string
type PrizeType string
type GameParticipation string

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
	GT_Tradicional GameType = "Tradicional"
	GT_Segunda     GameType = "Segunda"
	GT_Revancha    GameType = "Revancha"
	GT_SiempreSale GameType = "Siempre Sale"
	GT_PozoExtra   GameType = "Pozo Extra"
)

const (
	PTT_NoPrize     PrizeType = "Tradicional: No Prize"
	PTT_FirstPrize  PrizeType = "Tradicional: First Prize"
	PTT_SecondPrize PrizeType = "Tradicional: Second Prize"
	PTT_ThirdPrize  PrizeType = "Tradicional: Third Prize"
	PTS_NoPrize     PrizeType = "Segunda: No Prize"
	PTS_FirstPrize  PrizeType = "Segunda: First Prize"
	PTS_SecondPrize PrizeType = "Segunda: Second Prize"
	PTS_ThirdPrize  PrizeType = "Segunda: Third Prize"
	PTR_NoPrize     PrizeType = "Revancha: No Prize"
	PTR_Prize       PrizeType = "Revancha: Main Prize"
	PTSS_NoPrize    PrizeType = "Siempre Sale: No Prize"
	PTSS_Prize      PrizeType = "Siempre Sale: Main Prize"
	PTPE_NoPrize    PrizeType = "Pozo Extra: No Prize"
	PTPE_Prize      PrizeType = "Pozo Extra: Main Prize"
)

const (
	GP_TradicionalOnly                      GameParticipation = "Tradicional"
	GP_TradicionalAndRevancha               GameParticipation = "Tradicional, Revancha"
	GP_TradicionalAndRevanchaAndSiempreSale GameParticipation = "Tradicional, Revancha, Siempre Sale"
)
