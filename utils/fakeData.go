package utils

import (
	"q6-golang/models"

	"github.com/Pallinder/go-randomdata"
)

func GenerateFakePlayers(numberOfPlayersToGenerate int) []models.Player {
	var players []models.Player
	for i := 0; i < numberOfPlayersToGenerate; i++ {
		players = append(players, generateFakePlayer())
	}
	return players
}

func generateFakePlayer() models.Player {
	name := generateRandomName()
	age := generateRandomAdultAge()
	city := generateRandomCity()
	ticket := generateFakeTicket()
	return models.GetPlayer(name, age, city, ticket)
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

func generateFakeTicket() models.Ticket {
	nums := models.GenerateDrawing()
	t := models.GetTicket([]int{nums.FirstNumber, nums.SecondNumber, nums.ThirdNumber, nums.FourthNumber, nums.FifthNumber, nums.SixthNumber}, generateFakeGameParticipation())
	return t
}

func generateFakeGameParticipation() models.GameParticipation {
	n := randomdata.Number(3)
	if n == 0 {
		return models.GP_TradicionalOnly
	} else if n == 1 {
		return models.GP_TradicionalAndRevancha
	} else {
		return models.GP_TradicionalAndRevanchaAndSiempreSale
	}
}

func GetCustomTestPlayer() models.Player {
	return models.GetPlayer("Dario De Giacomo", 31, "Arroyo Seco", models.GetTicket([]int{7, 9, 11, 20, 32, 43}, models.GP_TradicionalAndRevanchaAndSiempreSale))
}
