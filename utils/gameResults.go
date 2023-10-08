package utils

import (
	"q6-golang/models"
	"sort"
)

func ExecuteGames(players []models.Player) models.GameResults {
	t := executeTradicional(players)
	s := executeSegunda(players)
	r := executeRevancha(players)
	ss := executeSiempreSale(players)
	pe := executePozoExtra(players, t.DrawingResults, s.DrawingResults, r.DrawingResults)
	return models.GameResults{GTRT: t, GTRS: s, GTRR: r, GTRSS: ss, GTRPE: pe}
}

func executeTradicional(players []models.Player) models.GameTypeResult {
	drawing := models.GenerateDrawing()
	//drawing := DrawingResult{FirstNumber: 1, SecondNumber: 3, ThirdNumber: 5, FourthNumber: 17, FifthNumber: 21, SixthNumber: 25}
	tradicionalNumbers := []int{drawing.FirstNumber, drawing.SecondNumber, drawing.ThirdNumber, drawing.FourthNumber, drawing.FifthNumber, drawing.SixthNumber}
	sort.Ints(tradicionalNumbers)
	return models.GameTypeResult{GameType: models.GT_Tradicional, DrawingResults: tradicionalNumbers}
}

func executeSegunda(players []models.Player) models.GameTypeResult {
	drawing := models.GenerateDrawing()
	segundaNumbers := []int{drawing.FirstNumber, drawing.SecondNumber, drawing.ThirdNumber, drawing.FourthNumber, drawing.FifthNumber, drawing.SixthNumber}
	sort.Ints(segundaNumbers)
	return models.GameTypeResult{GameType: models.GT_Segunda, DrawingResults: segundaNumbers}
}

func executeRevancha(players []models.Player) models.GameTypeResult {
	drawing := models.GenerateDrawing()
	revanchaNumbers := []int{drawing.FirstNumber, drawing.SecondNumber, drawing.ThirdNumber, drawing.FourthNumber, drawing.FifthNumber, drawing.SixthNumber}
	sort.Ints(revanchaNumbers)
	return models.GameTypeResult{GameType: models.GT_Revancha, DrawingResults: revanchaNumbers}
}

func executeSiempreSale(players []models.Player) models.GameTypeResult {
	drawing := models.GenerateDrawing()
	siempreSaleNumbers := []int{drawing.FirstNumber, drawing.SecondNumber, drawing.ThirdNumber, drawing.FourthNumber, drawing.FifthNumber, drawing.SixthNumber}
	sort.Ints(siempreSaleNumbers)
	return models.GameTypeResult{GameType: models.GT_SiempreSale, DrawingResults: siempreSaleNumbers}
}

func executePozoExtra(players []models.Player, tNumbers []int, sNumbers []int, rNumbers []int) models.GameTypeResult {
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
	return models.GameTypeResult{GameType: models.GT_PozoExtra, DrawingResults: pozoExtraNumbers}
}
