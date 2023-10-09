package utils

import "q6-golang/models"

func CalculateTotalTickets(ch chan models.Tickets, players []models.Player) {
	totalTradicionalTickets := 0
	totalRevanchaTickets := 0
	totalSiempreSaleTickets := 0
	for _, player := range players {
		for _, ticket := range player.Tickets {
			if ticket.Participation == models.GP_TradicionalAndRevanchaAndSiempreSale {
				totalSiempreSaleTickets++
			} else if ticket.Participation == models.GP_TradicionalAndRevancha {
				totalRevanchaTickets++
			} else {
				totalTradicionalTickets++
			}
		}
	}
	ch <- models.Tickets{TTickets: totalTradicionalTickets, TRTickets: totalRevanchaTickets, TRSSTickets: totalSiempreSaleTickets}
}
