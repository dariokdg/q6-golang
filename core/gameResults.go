package core

import (
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
	return GameTypeResult{GameType: GT_Tradicional, DrawingResults: tradicionalNumbers}
}

func executeSegunda(players []Player) GameTypeResult {
	drawing := GenerateDrawing()
	segundaNumbers := []int{drawing.FirstNumber, drawing.SecondNumber, drawing.ThirdNumber, drawing.FourthNumber, drawing.FifthNumber, drawing.SixthNumber}
	sort.Ints(segundaNumbers)
	return GameTypeResult{GameType: GT_Segunda, DrawingResults: segundaNumbers}
}

func executeRevancha(players []Player) GameTypeResult {
	drawing := GenerateDrawing()
	revanchaNumbers := []int{drawing.FirstNumber, drawing.SecondNumber, drawing.ThirdNumber, drawing.FourthNumber, drawing.FifthNumber, drawing.SixthNumber}
	sort.Ints(revanchaNumbers)
	return GameTypeResult{GameType: GT_Revancha, DrawingResults: revanchaNumbers}
}

func executeSiempreSale(players []Player) GameTypeResult {
	drawing := GenerateDrawing()
	siempreSaleNumbers := []int{drawing.FirstNumber, drawing.SecondNumber, drawing.ThirdNumber, drawing.FourthNumber, drawing.FifthNumber, drawing.SixthNumber}
	sort.Ints(siempreSaleNumbers)
	return GameTypeResult{GameType: GT_SiempreSale, DrawingResults: siempreSaleNumbers}
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
	return GameTypeResult{GameType: GT_PozoExtra, DrawingResults: pozoExtraNumbers}
}
