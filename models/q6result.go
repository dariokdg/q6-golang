package models

type Q6Results struct {
	GTRT  Result `json:"tradicional"`
	GTRS  Result `json:"segunda"`
	GTRR  Result `json:"revancha"`
	GTRSS Result `json:"siempreSale"`
	GTRPE Result `json:"pozoExtra"`
}
