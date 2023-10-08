package core

import (
	"fmt"
	"os"
	"strconv"

	"github.com/jedib0t/go-pretty/v6/table"
)

func printTotalSales(tS TotalSales) {
	fmt.Println("--------------------------------------------")
	fmt.Printf("TOTAL NUMBER OF PLAYERS: %d\n", GetTotalGamePlayers(tS))
	fmt.Printf("    > TRADICIONAL PLAYERS: %d\n", tS.TotalTradicionalPlayers)
	fmt.Printf("    > REVANCHA PLAYERS: %d\n", tS.TotalRevanchaPlayers)
	fmt.Printf("    > SIEMPRE SALE PLAYERS: %d\n", tS.TotalSiempreSalePlayers)
	fmt.Println("--------------------------------------------")
	fmt.Println("\n--------------------------------------------")
	fmt.Printf("TOTAL SALES: $ %s\n", GetTotalGameSales(tS))
	fmt.Printf("    > TRADICIONAL SALES: $ %s\n", tS.TotalTradicionalSales)
	fmt.Printf("    > REVANCHA SALES: $ %s\n", tS.TotalRevanchaSales)
	fmt.Printf("    > SIEMPRE SALE SALES: $ %s\n", tS.TotalSiempreSaleSales)
	fmt.Println("--------------------------------------------")
}

func printPrizes(pG PrizeGenerator) {
	fmt.Println("")
	t := table.NewWriter()
	t.SetTitle("QUINI 6 PRIZE LIST")
	t.Style().Options.DrawBorder = true
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"GAME TYPE", "PRIZE CATEGORY", "NUMBER OF HITS", "TOTAL PRIZE"})
	t.AppendRow([]interface{}{"TRADICIONAL", "FIRST PRIZE", "6", "$ " + pG.TradicionalFirstPrize.Round(2).String()})
	t.AppendSeparator()
	t.AppendRow([]interface{}{"TRADICIONAL", "SECOND PRIZE", "5", "$ " + pG.TradicionalSecondPrize.Round(2).String()})
	t.AppendSeparator()
	t.AppendRow([]interface{}{"TRADICIONAL", "THIRD PRIZE", "4", "$ " + pG.TradicionalThirdPrize.Round(2).String()})
	t.AppendSeparator()
	t.AppendRow([]interface{}{"SEGUNDA", "FIRST PRIZE", "6", "$ " + pG.SegundaFirstPrize.Round(2).String()})
	t.AppendSeparator()
	t.AppendRow([]interface{}{"SEGUNDA", "SECOND PRIZE", "5", "$ " + pG.SegundaSecondPrize.Round(2).String()})
	t.AppendSeparator()
	t.AppendRow([]interface{}{"SEGUNDA", "THIRD PRIZE", "4", "$ " + pG.SegundaThirdPrize.Round(2).String()})
	t.AppendSeparator()
	t.AppendRow([]interface{}{"REVANCHA", "MAIN PRIZE", "6", "$ " + pG.RevanchaPrize.Round(2).String()})
	t.AppendSeparator()
	t.AppendRow([]interface{}{"SIEMPRE SALE", "MAIN PRIZE", "6/5/4/3/2/1", "$ " + pG.SiempreSalePrize.Round(2).String()})
	t.AppendSeparator()
	t.AppendFooter(table.Row{"POZO EXTRA", "MAIN PRIZE", "6", "$ " + pG.PozoExtraPrize.Round(2).String()})
	t.Render()

}

func printDrawingResults(gR GameResults) {

	GTRT := gR.GTRT
	GTRS := gR.GTRS
	GTRR := gR.GTRR
	GTRSS := gR.GTRSS

	fmt.Println("")
	t := table.NewWriter()
	t.SetTitle("QUINI 6 DRAWINGS")
	t.Style().Options.DrawBorder = true
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"GAME TYPE", "FIRST NUMBER", "SECOND NUMBER", "THIRD NUMBER", "FOURTH NUMBER", "FIFTH NUMBER", "SIXTH NUMBER"})
	t.AppendRow([]interface{}{"TRADICIONAL", GTRT.DrawingResults[0], GTRT.DrawingResults[1], GTRT.DrawingResults[2], GTRT.DrawingResults[3], GTRT.DrawingResults[4], GTRT.DrawingResults[5]})
	t.AppendSeparator()
	t.AppendRow([]interface{}{"SEGUNDA", GTRS.DrawingResults[0], GTRS.DrawingResults[1], GTRS.DrawingResults[2], GTRS.DrawingResults[3], GTRS.DrawingResults[4], GTRS.DrawingResults[5]})
	t.AppendSeparator()
	t.AppendRow([]interface{}{"REVANCHA", GTRR.DrawingResults[0], GTRR.DrawingResults[1], GTRR.DrawingResults[2], GTRR.DrawingResults[3], GTRR.DrawingResults[4], GTRR.DrawingResults[5]})
	t.AppendSeparator()
	t.AppendFooter(table.Row{"SIEMPRE SALE", GTRSS.DrawingResults[0], GTRSS.DrawingResults[1], GTRSS.DrawingResults[2], GTRSS.DrawingResults[3], GTRSS.DrawingResults[4], GTRSS.DrawingResults[5]})
	t.Render()

	GTRPE := gR.GTRPE

	n01 := getCorrectValue(GTRPE, 1)
	n02 := getCorrectValue(GTRPE, 2)
	n03 := getCorrectValue(GTRPE, 3)
	n04 := getCorrectValue(GTRPE, 4)
	n05 := getCorrectValue(GTRPE, 5)
	n06 := getCorrectValue(GTRPE, 6)
	n07 := getCorrectValue(GTRPE, 7)
	n08 := getCorrectValue(GTRPE, 8)
	n09 := getCorrectValue(GTRPE, 9)
	n10 := getCorrectValue(GTRPE, 10)
	n11 := getCorrectValue(GTRPE, 11)
	n12 := getCorrectValue(GTRPE, 12)
	n13 := getCorrectValue(GTRPE, 13)
	n14 := getCorrectValue(GTRPE, 14)
	n15 := getCorrectValue(GTRPE, 15)
	n16 := getCorrectValue(GTRPE, 16)
	n17 := getCorrectValue(GTRPE, 17)
	n18 := getCorrectValue(GTRPE, 18)

	t2 := table.NewWriter()
	t2.Style().Options.DrawBorder = true
	t2.SetTitle("POZO EXTRA")
	t2.SetOutputMirror(os.Stdout)
	t2.AppendHeader(table.Row{n01, n02, n03, n04, n05, n06})
	if n07 != " " {
		t2.AppendRow([]interface{}{n07, n08, n09, n10, n11, n12})
		t2.AppendSeparator()
		if n13 != " " {
			t2.AppendRow([]interface{}{n13, n14, n15, n16, n17, n18})
			t2.AppendSeparator()
		}
	}
	fmt.Println("")
	t2.Render()
}

func getCorrectValue(g GameTypeResult, index int) string {
	if len(g.DrawingResults) > index {
		return strconv.Itoa(g.DrawingResults[index-1])
	}
	return " "
}

func printWinners(drawings GameResults, pG PrizeGenerator, winners []Winner) {
	tFP_winners := winners[0]
	tSP_winners := winners[1]
	tTP_winners := winners[2]

	sFP_winners := winners[3]
	sSP_winners := winners[4]
	sTP_winners := winners[5]

	r_winners := winners[6]

	sS_winners := winners[7]

	pE_winners := winners[8]

	fmt.Println("")
	t := table.NewWriter()
	t.SetTitle("QUINI 6 RESULTS")
	t.Style().Options.DrawBorder = true
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"GAME TYPE", "PRIZE CATEGORY", "TOTAL PRIZE AMOUNT", "NUMBER OF WINNERS", "NUMBER OF HITS", "PRIZE FOR EACH WINNER"})
	t.AppendRow([]interface{}{"TRADICIONAL", "FIRST PRIZE", "$ " + pG.TradicionalFirstPrize.Round(2).String(), strconv.Itoa(len(tFP_winners.PrizeWinnerList)), "6", "$ " + tFP_winners.PrizeAmountPerWinner.Round(2).String()})
	t.AppendSeparator()
	t.AppendRow([]interface{}{"TRADICIONAL", "SECOND PRIZE", "$ " + pG.TradicionalSecondPrize.Round(2).String(), strconv.Itoa(len(tSP_winners.PrizeWinnerList)), "5", "$ " + tSP_winners.PrizeAmountPerWinner.Round(2).String()})
	t.AppendSeparator()
	t.AppendRow([]interface{}{"TRADICIONAL", "THIRD PRIZE", "$ " + pG.TradicionalThirdPrize.Round(2).String(), strconv.Itoa(len(tTP_winners.PrizeWinnerList)), "4", "$ " + tTP_winners.PrizeAmountPerWinner.Round(2).String()})
	t.AppendSeparator()
	t.AppendRow([]interface{}{"SEGUNDA", "FIRST PRIZE", "$ " + pG.SegundaFirstPrize.Round(2).String(), strconv.Itoa(len(sFP_winners.PrizeWinnerList)), "6", "$ " + sFP_winners.PrizeAmountPerWinner.Round(2).String()})
	t.AppendSeparator()
	t.AppendRow([]interface{}{"SEGUNDA", "SECOND PRIZE", "$ " + pG.SegundaSecondPrize.Round(2).String(), strconv.Itoa(len(sSP_winners.PrizeWinnerList)), "5", "$ " + sSP_winners.PrizeAmountPerWinner.Round(2).String()})
	t.AppendSeparator()
	t.AppendRow([]interface{}{"SEGUNDA", "THIRD PRIZE", "$ " + pG.SegundaThirdPrize.Round(2).String(), strconv.Itoa(len(sTP_winners.PrizeWinnerList)), "4", "$ " + sTP_winners.PrizeAmountPerWinner.Round(2).String()})
	t.AppendSeparator()
	t.AppendRow([]interface{}{"REVANCHA", "MAIN PRIZE", "$ " + pG.RevanchaPrize.Round(2).String(), strconv.Itoa(len(r_winners.PrizeWinnerList)), "6", "$ " + r_winners.PrizeAmountPerWinner.Round(2).String()})
	t.AppendSeparator()
	t.AppendRow([]interface{}{"SIEMPRE SALE", "MAIN PRIZE", "$ " + pG.SiempreSalePrize.Round(2).String(), strconv.Itoa(len(sS_winners.PrizeWinnerList)), sS_winners.NumberOfMatches, "$ " + sS_winners.PrizeAmountPerWinner.Round(2).String()})
	t.AppendSeparator()
	t.AppendFooter(table.Row{"POZO EXTRA", "MAIN PRIZE", "$ " + pG.PozoExtraPrize.Round(2).String(), strconv.Itoa(len(pE_winners.PrizeWinnerList)), "6", "$ " + pE_winners.PrizeAmountPerWinner.Round(2).String()})
	t.Render()
}
