package utils

import (
	"q6-golang/models"

	"github.com/shopspring/decimal"
)

func GetWinners(gameType models.GameType, prizeType models.PrizeType, numberOfMatches int, players []models.Player, winners []models.Ticket, prizeAmountTotal decimal.Decimal) models.Winners {
	var prizeAmountPerWinner decimal.Decimal
	numberOfWinners := len(winners)
	numberOfWinners_d := decimal.NewFromInt(int64(numberOfWinners))
	if numberOfWinners == 0 {
		prizeAmountPerWinner = decimal.NewFromInt(0)
	} else {
		prizeAmountPerWinner = prizeAmountTotal.Div(numberOfWinners_d)
	}

	//for each winning ticket, we need to update owner's "prize money"
	for _, ticket := range winners {
		UpdatePrizeWinnerMoney(&players, &ticket, prizeAmountPerWinner)
	}
	return models.Winners{GameType: gameType, PrizeType: prizeType, NumberOfMatches: numberOfMatches, PrizeWinnerList: winners, PrizeAmountPerWinner: prizeAmountPerWinner.Round(2)}
}

func UpdatePrizeWinnerMoney(players *[]models.Player, ticket *models.Ticket, prizeAmountPerWinner decimal.Decimal) {
	player := models.GetTicketOwner(*players, *ticket)
	player.PrizeMoney = decimal.Sum(player.PrizeMoney, prizeAmountPerWinner).Round(2)
}

func CalculateWinners(players []models.Player, drawings models.Q6Results, pG models.Prizes) []models.Winners {
	ch1 := make(chan models.Winners, 1)
	go CheckPrizesTradicionalFirstPrize(ch1, players, drawings.GTRT, pG.TradicionalFirstPrize)

	ch2 := make(chan models.Winners, 1)
	go CheckPrizesTradicionalSecondPrize(ch2, players, drawings.GTRT, pG.TradicionalSecondPrize)

	ch3 := make(chan models.Winners, 1)
	go CheckPrizesTradicionalThirdPrize(ch3, players, drawings.GTRT, pG.TradicionalThirdPrize)

	ch4 := make(chan models.Winners, 1)
	go CheckPrizesSegundaFirstPrize(ch4, players, drawings.GTRS, pG.SegundaFirstPrize)

	ch5 := make(chan models.Winners, 1)
	go CheckPrizesSegundaSecondPrize(ch5, players, drawings.GTRS, pG.SegundaSecondPrize)

	ch6 := make(chan models.Winners, 1)
	go CheckPrizesSegundaThirdPrize(ch6, players, drawings.GTRS, pG.SegundaThirdPrize)

	ch7 := make(chan models.Winners, 1)
	go CheckPrizesRevanchaPrize(ch7, players, drawings.GTRR, pG.RevanchaPrize)

	ch8 := make(chan models.Winners, 1)
	go CheckPrizesSiempreSalePrize(ch8, players, drawings.GTRSS, pG.SiempreSalePrize)

	ch9 := make(chan models.Winners, 1)
	go CheckPrizesPozoExtraPrize(ch9, players, drawings.GTRPE, pG.PozoExtraPrize, drawings)

	tFP_winners := <-ch1
	tSP_winners := <-ch2
	tTP_winners := <-ch3
	sFP_winners := <-ch4
	sSP_winners := <-ch5
	sTP_winners := <-ch6
	r_winners := <-ch7
	sS_winners := <-ch8
	pE_winners := <-ch9

	return []models.Winners{tFP_winners, tSP_winners, tTP_winners, sFP_winners, sSP_winners, sTP_winners, r_winners, sS_winners, pE_winners}
}
