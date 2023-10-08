package models

type GameResults struct {
	GTRT  GameTypeResult `json:"gameTypeResultTradicional"`
	GTRS  GameTypeResult `json:"gameTypeResultSegunda"`
	GTRR  GameTypeResult `json:"gameTypeResultRevancha"`
	GTRSS GameTypeResult `json:"gameTypeResultSiempreSale"`
	GTRPE GameTypeResult `json:"gameTypeResultPozoExtra"`
}
