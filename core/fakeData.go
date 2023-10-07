package core

import "math/rand"

func GenerateFakePlayers(numberOfPlayersToGenerate int) []Player {
	var players []Player
	for i := 0; i < numberOfPlayersToGenerate; i++ {
		players = append(players, generateFakePlayer())
	}
	return players
}

func generateFakePlayer() Player {
	n := "Testy name"
	age := rand.Intn(81) + 18
	c := "Rosario"
	ad := "Calle Falsa 123"
	p := "3410000000"
	t := generateFakeTicket()
	return GetPlayer(n, age, c, ad, p, t)
}

func generateFakeTicket() Ticket {
	nums := GenerateDrawing()
	t := GetTicket([]int{nums.FirstNumber, nums.SecondNumber, nums.ThirdNumber, nums.FourthNumber, nums.FifthNumber, nums.SixthNumber}, generateFakeGameParticipation())
	return t
}

func generateFakeGameParticipation() GameParticipation {
	n := rand.Intn(3)
	if n == 0 {
		return GP_TradicionalOnly
	} else if n == 1 {
		return GP_TradicionalAndRevancha
	} else {
		return GP_TradicionalAndRevanchaAndSiempreSale
	}
}
