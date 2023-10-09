package utils

import "q6-golang/models"

func CalculateWinners(players []models.Player, drawings models.Q6Results, pG models.Prizes) []models.Winners {
	tFP_winners := CheckPrizesTradicionalFirstPrize(players, drawings.GTRT, pG.TradicionalFirstPrize)
	tSP_winners := CheckPrizesTradicionalSecondPrize(players, drawings.GTRT, pG.TradicionalSecondPrize)
	tTP_winners := CheckPrizesTradicionalThirdPrize(players, drawings.GTRT, pG.TradicionalThirdPrize)

	sFP_winners := CheckPrizesSegundaFirstPrize(players, drawings.GTRS, pG.SegundaFirstPrize)
	sSP_winners := CheckPrizesSegundaSecondPrize(players, drawings.GTRS, pG.SegundaSecondPrize)
	sTP_winners := CheckPrizesSegundaThirdPrize(players, drawings.GTRS, pG.SegundaThirdPrize)

	var validRevanchaPlayers []models.Player
	for _, p := range players {
		if p.Quini6Ticket.Participation == models.GP_TradicionalAndRevancha || p.Quini6Ticket.Participation == models.GP_TradicionalAndRevanchaAndSiempreSale {
			validRevanchaPlayers = append(validRevanchaPlayers, p)
		}
	}
	r_winners := CheckPrizesRevanchaPrize(validRevanchaPlayers, drawings.GTRR, pG.RevanchaPrize)

	var validSiempreSalePlayers []models.Player
	for _, p := range players {
		if p.Quini6Ticket.Participation == models.GP_TradicionalAndRevancha || p.Quini6Ticket.Participation == models.GP_TradicionalAndRevanchaAndSiempreSale {
			validSiempreSalePlayers = append(validSiempreSalePlayers, p)
		}
	}
	sS_winners := CheckPrizesSiempreSalePrize(validSiempreSalePlayers, drawings.GTRSS, pG.SiempreSalePrize)

	pE_winners := CheckPrizesPozoExtraPrize(players, drawings.GTRPE, pG.PozoExtraPrize, drawings)

	return []models.Winners{tFP_winners, tSP_winners, tTP_winners, sFP_winners, sSP_winners, sTP_winners, r_winners, sS_winners, pE_winners}
}
