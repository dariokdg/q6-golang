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
	myPlayer := models.GetPlayer(name, age, city)
	var tickets []models.Ticket
	for i := 1; i <= generateRandomNumberOfTickets(); i++ {
		tickets = append(tickets, generateFakeTicket(myPlayer))
	}
	models.AssignTicketsToPlayer(&myPlayer, tickets)
	return myPlayer
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

func generateRandomNumberOfTickets() int {
	return randomdata.Number(5) + 1
}

func generateFakeTicket(player models.Player) models.Ticket {
	nums := models.GenerateDrawing()
	t := models.GetTicket([]int{nums.FirstNumber, nums.SecondNumber, nums.ThirdNumber, nums.FourthNumber, nums.FifthNumber, nums.SixthNumber}, generateFakeGameParticipation(), player.ID)
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
	var tickets []models.Ticket
	myPlayer := models.GetPlayer("Dario De Giacomo", 31, "Arroyo Seco")
	tickets = append(tickets, models.GetTicket([]int{7, 9, 11, 20, 32, 43}, models.GP_TradicionalAndRevanchaAndSiempreSale, myPlayer.ID))
	models.AssignTicketsToPlayer(&myPlayer, tickets)
	return myPlayer
}
