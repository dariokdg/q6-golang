package models

import (
	"slices"

	"github.com/Pallinder/go-randomdata"
)

type Drawing struct {
	FirstNumber  int `json:"firstNumber"`
	SecondNumber int `json:"secondNumber"`
	ThirdNumber  int `json:"thirdNumber"`
	FourthNumber int `json:"fourthNumber"`
	FifthNumber  int `json:"fifthNumber"`
	SixthNumber  int `json:"sixthNumber"`
}

func GenerateDrawing() Drawing {
	var dR Drawing
	dR.FirstNumber = getFirstNumber(&dR)
	dR.SecondNumber = getSecondNumber(&dR)
	dR.ThirdNumber = getThirdNumber(&dR)
	dR.FourthNumber = getFourthNumber(&dR)
	dR.FifthNumber = getFifthNumber(&dR)
	dR.SixthNumber = getSixthNumber(&dR)
	return dR
}

func getFirstNumber(dR *Drawing) int {
	return getQ6Number()
}

func getSecondNumber(dR *Drawing) int {
	keepRunning := true
	var secondNumber int
	drawingResultNumbers := []int{dR.FirstNumber}
	for keepRunning {
		secondNumber = getQ6Number()
		if !slices.Contains(drawingResultNumbers, secondNumber) {
			keepRunning = false
		}
	}
	return secondNumber
}

func getThirdNumber(dR *Drawing) int {
	keepRunning := true
	var thirdNumber int
	drawingResultNumbers := []int{dR.FirstNumber, dR.SecondNumber}
	for keepRunning {
		thirdNumber = getQ6Number()
		if !slices.Contains(drawingResultNumbers, thirdNumber) {
			keepRunning = false
		}
	}
	return thirdNumber
}

func getFourthNumber(dR *Drawing) int {
	keepRunning := true
	var fourthNumber int
	drawingResultNumbers := []int{dR.FirstNumber, dR.SecondNumber, dR.ThirdNumber}
	for keepRunning {
		fourthNumber = getQ6Number()
		if !slices.Contains(drawingResultNumbers, fourthNumber) {
			keepRunning = false
		}
	}
	return fourthNumber
}

func getFifthNumber(dR *Drawing) int {
	keepRunning := true
	var fifthNumber int
	drawingResultNumbers := []int{dR.FirstNumber, dR.SecondNumber, dR.ThirdNumber, dR.FourthNumber}
	for keepRunning {
		fifthNumber = getQ6Number()
		if !slices.Contains(drawingResultNumbers, fifthNumber) {
			keepRunning = false
		}
	}
	return fifthNumber
}

func getSixthNumber(dR *Drawing) int {
	keepRunning := true
	var sixthNumber int
	drawingResultNumbers := []int{dR.FirstNumber, dR.SecondNumber, dR.ThirdNumber, dR.FourthNumber, dR.FifthNumber}
	for keepRunning {
		sixthNumber = getQ6Number()
		if !slices.Contains(drawingResultNumbers, sixthNumber) {
			keepRunning = false
		}
	}
	return sixthNumber
}

func getQ6Number() int {
	return randomdata.Number(45) + 1
}
