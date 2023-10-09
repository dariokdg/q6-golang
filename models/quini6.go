package models

type Quini6 struct {
	Players      int       `json:"totalPlayers"`
	TotalSales   Sales     `json:"sales"`
	TotalPlayers Players   `json:"players"`
	Prizes       Prizes    `json:"prizes"`
	Results      Q6Results `json:"results"`
	Winners      []Winners `json:"winners"`
}
