package models

import (
	"github.com/shopspring/decimal"
)

type Sales struct {
	TSales    decimal.Decimal `json:"tradicionalOnly"`
	TRSales   decimal.Decimal `json:"tradicionalAndRevancha"`
	TRSSSales decimal.Decimal `json:"tradicionalRevanchaAndSiempreSale"`
}

func GetTotalGameSales(sales Sales) decimal.Decimal {
	return decimal.Sum(sales.TSales, sales.TRSales, sales.TRSSSales)
}

type Players struct {
	TPlayers    int `json:"tradicionalOnly"`
	TRPlayers   int `json:"tradicionalAndRevancha"`
	TRSSPlayers int `json:"tradicionalRevanchaAndSiempreSale"`
}

func GetTotalGamePlayers(players Players) int {
	return players.TPlayers + players.TRPlayers + players.TRSSPlayers
}

type Tickets struct {
	TTickets    int `json:"tradicionalOnly"`
	TRTickets   int `json:"tradicionalAndRevancha"`
	TRSSTickets int `json:"tradicionalRevanchaAndSiempreSale"`
}

func GetTotalGameTickets(tickets Tickets) int {
	return tickets.TTickets + tickets.TRTickets + tickets.TRSSTickets
}
