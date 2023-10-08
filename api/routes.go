package api

import (
	"net/http"
	"q6-golang/handlers"
	"q6-golang/models"
	"q6-golang/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

func StartGame(c *gin.Context) {
	testPlayer := models.GetPlayer("Juanito Perez", 20, "Rosario", models.GetTicket([]int{1, 5, 6, 12, 25, 41}, models.GP_TradicionalAndRevancha))
	players := []models.Player{testPlayer}
	if _, err := strconv.ParseInt(c.Query("players"), 10, 64); err == nil {
		numOfPlayers, _ := strconv.ParseInt(c.Query("players"), 10, 64)
		players = append(players, utils.GenerateFakePlayers(int(numOfPlayers))...)
	} else {
		players = append(players, utils.GenerateFakePlayers(2500)...)
	}
	q6 := handlers.ExecuteGame(players)
	c.IndentedJSON(http.StatusOK, q6)
}
