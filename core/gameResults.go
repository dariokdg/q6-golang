package core

import (
	e "q6-golang/enumerators"
	"sort"
)

type GameResults struct {
	GTRT  GameTypeResult `json:"GTRT"`
	GTRS  GameTypeResult `json:"GTRS"`
	GTRR  GameTypeResult `json:"GTRR"`
	GTRSS GameTypeResult `json:"GTRSS"`
	GTRPE GameTypeResult `json:"GTRPE"`
}

func executeGames(players []Player) GameResults {
	t := executeTradicional(players)
	s := executeSegunda(players)
	r := executeRevancha(players)
	ss := executeSiempreSale(players)
	pe := executePozoExtra(players, t.DrawingResults, s.DrawingResults, r.DrawingResults)
	return GameResults{t, s, r, ss, pe}
}

func executeTradicional(players []Player) GameTypeResult {
	drawing := GenerateDrawing()
	//drawing := DrawingResult{FirstNumber: 1, SecondNumber: 3, ThirdNumber: 5, FourthNumber: 17, FifthNumber: 21, SixthNumber: 25}
	tradicionalNumbers := []int{drawing.FirstNumber, drawing.SecondNumber, drawing.ThirdNumber, drawing.FourthNumber, drawing.FifthNumber, drawing.SixthNumber}
	sort.Ints(tradicionalNumbers)
	return GameTypeResult{Players: players, GameType: e.GT_Tradicional, DrawingResults: tradicionalNumbers}
}

func executeSegunda(players []Player) GameTypeResult {
	drawing := GenerateDrawing()
	//drawing := DrawingResult{FirstNumber: 2, SecondNumber: 7, ThirdNumber: 11, FourthNumber: 13, FifthNumber: 23, SixthNumber: 32}
	segundaNumbers := []int{drawing.FirstNumber, drawing.SecondNumber, drawing.ThirdNumber, drawing.FourthNumber, drawing.FifthNumber, drawing.SixthNumber}
	sort.Ints(segundaNumbers)
	return GameTypeResult{Players: players, GameType: e.GT_Segunda, DrawingResults: segundaNumbers}
}

func executeRevancha(players []Player) GameTypeResult {
	drawing := GenerateDrawing()
	//drawing := DrawingResult{FirstNumber: 2, SecondNumber: 7, ThirdNumber: 11, FourthNumber: 13, FifthNumber: 23, SixthNumber: 32}
	revanchaNumbers := []int{drawing.FirstNumber, drawing.SecondNumber, drawing.ThirdNumber, drawing.FourthNumber, drawing.FifthNumber, drawing.SixthNumber}
	sort.Ints(revanchaNumbers)
	var validPlayers []Player
	for _, p := range players {
		if p.Quini6Ticket.Games == e.GP_TradicionalAndRevancha || p.Quini6Ticket.Games == e.GP_TradicionalAndRevanchaAndSiempreSale {
			validPlayers = append(validPlayers, p)
		}
	}
	return GameTypeResult{Players: validPlayers, GameType: e.GT_Revancha, DrawingResults: revanchaNumbers}
}

func executeSiempreSale(players []Player) GameTypeResult {
	drawing := GenerateDrawing()
	siempreSaleNumbers := []int{drawing.FirstNumber, drawing.SecondNumber, drawing.ThirdNumber, drawing.FourthNumber, drawing.FifthNumber, drawing.SixthNumber}
	sort.Ints(siempreSaleNumbers)
	var validPlayers []Player
	for _, p := range players {
		if p.Quini6Ticket.Games == e.GP_TradicionalAndRevanchaAndSiempreSale {
			validPlayers = append(validPlayers, p)
		}
	}
	return GameTypeResult{Players: validPlayers, GameType: e.GT_SiempreSale, DrawingResults: siempreSaleNumbers}
}

func executePozoExtra(players []Player, tNumbers []int, sNumbers []int, rNumbers []int) GameTypeResult {
	var tempList []int
	tempList = append(tempList, tNumbers...)
	tempList = append(tempList, sNumbers...)
	tempList = append(tempList, rNumbers...)
	seen := make(map[int]bool)
	pozoExtraNumbers := []int{}
	for _, val := range tempList {
		if _, ok := seen[val]; !ok {
			seen[val] = true
			pozoExtraNumbers = append(pozoExtraNumbers, val)
		}
	}
	sort.Ints(pozoExtraNumbers)
	return GameTypeResult{Players: players, GameType: e.GT_PozoExtra, DrawingResults: pozoExtraNumbers}
}
