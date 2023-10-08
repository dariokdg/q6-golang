package models

type Game struct {
	TotalSales TotalSales     `json:"sales"`
	Prizes     PrizeGenerator `json:"prizes"`
	Results    GameResults    `json:"results"`
	Winners    []Winner       `json:"winners"`
}
