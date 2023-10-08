package core

import "github.com/Pallinder/go-randomdata"

func GenerateFakePlayers(numberOfPlayersToGenerate int) []Player {
	var players []Player
	for i := 0; i < numberOfPlayersToGenerate; i++ {
		players = append(players, generateFakePlayer())
	}
	return players
}

func generateFakePlayer() Player {
	name := generateRandomName()
	age := generateRandomAdultAge()
	city := generateRandomCity()
	ticket := generateFakeTicket()
	return GetPlayer(name, age, city, ticket)
}

func generateRandomName() string {
	return randomdata.FullName(randomdata.RandomGender)
}

func generateRandomAdultAge() int {
	return randomdata.Number(81) + 18
}

func generateRandomCity() string {
	return randomdata.City()
}

func generateFakeTicket() Ticket {
	nums := GenerateDrawing()
	t := GetTicket([]int{nums.FirstNumber, nums.SecondNumber, nums.ThirdNumber, nums.FourthNumber, nums.FifthNumber, nums.SixthNumber}, generateFakeGameParticipation())
	return t
}

func generateFakeGameParticipation() GameParticipation {
	n := randomdata.Number(3)
	if n == 0 {
		return GP_TradicionalOnly
	} else if n == 1 {
		return GP_TradicionalAndRevancha
	} else {
		return GP_TradicionalAndRevanchaAndSiempreSale
	}
}
