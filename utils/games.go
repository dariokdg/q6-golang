package utils

import (
	"q6-golang/models"
	"sort"
)

func ExecuteGames(ch chan models.Q6Results, players []models.Player) {
	t := tradicional(players)
	s := segunda(players)
	r := revancha(players)
	ss := siempreSale(players)
	pe := pozoExtra(players, t.DrawingResults, s.DrawingResults, r.DrawingResults)
	ch <- models.Q6Results{GTRT: t, GTRS: s, GTRR: r, GTRSS: ss, GTRPE: pe}
}

func tradicional(players []models.Player) models.Result {
	drawing := models.GenerateDrawing()
	tradicionalNumbers := []int{drawing.FirstNumber, drawing.SecondNumber, drawing.ThirdNumber, drawing.FourthNumber, drawing.FifthNumber, drawing.SixthNumber}
	sort.Ints(tradicionalNumbers)
	return models.Result{GameType: models.GT_Tradicional, DrawingResults: tradicionalNumbers}
}

func segunda(players []models.Player) models.Result {
	drawing := models.GenerateDrawing()
	segundaNumbers := []int{drawing.FirstNumber, drawing.SecondNumber, drawing.ThirdNumber, drawing.FourthNumber, drawing.FifthNumber, drawing.SixthNumber}
	sort.Ints(segundaNumbers)
	return models.Result{GameType: models.GT_Segunda, DrawingResults: segundaNumbers}
}

func revancha(players []models.Player) models.Result {
	drawing := models.GenerateDrawing()
	revanchaNumbers := []int{drawing.FirstNumber, drawing.SecondNumber, drawing.ThirdNumber, drawing.FourthNumber, drawing.FifthNumber, drawing.SixthNumber}
	sort.Ints(revanchaNumbers)
	return models.Result{GameType: models.GT_Revancha, DrawingResults: revanchaNumbers}
}

func siempreSale(players []models.Player) models.Result {
	drawing := models.GenerateDrawing()
	siempreSaleNumbers := []int{drawing.FirstNumber, drawing.SecondNumber, drawing.ThirdNumber, drawing.FourthNumber, drawing.FifthNumber, drawing.SixthNumber}
	sort.Ints(siempreSaleNumbers)
	return models.Result{GameType: models.GT_SiempreSale, DrawingResults: siempreSaleNumbers}
}

func pozoExtra(players []models.Player, tNumbers []int, sNumbers []int, rNumbers []int) models.Result {
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
	return models.Result{GameType: models.GT_PozoExtra, DrawingResults: pozoExtraNumbers}
}
