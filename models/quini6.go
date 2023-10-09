package models

type Quini6 struct {
	TotalSales Sales     `json:"sales"`
	Prizes     Prizes    `json:"prizes"`
	Results    Q6Results `json:"results"`
	Winners    []Winners `json:"winners"`
}
