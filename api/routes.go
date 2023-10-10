package api

import (
	"net/http"
	"q6-golang/handlers"
	"q6-golang/models"
	"q6-golang/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Play(c *gin.Context) {
	var players []models.Player

	//check if we want to include a premade custom player in the game
	param, exists := c.GetQuery("premadePlayer")
	if exists && param == "true" {
		players = append(players, utils.GetPremadeTestPlayer())
	}

	//check if we want to include a custom player in the game
	_, existsCustom := c.GetQuery("n1")
	if existsCustom {
		if _, err := strconv.ParseInt(c.Query("n1"), 10, 64); err == nil {
			n1, _ := strconv.ParseInt(c.Query("n1"), 10, 64)
			n2, _ := strconv.ParseInt(c.Query("n2"), 10, 64)
			n3, _ := strconv.ParseInt(c.Query("n3"), 10, 64)
			n4, _ := strconv.ParseInt(c.Query("n4"), 10, 64)
			n5, _ := strconv.ParseInt(c.Query("n5"), 10, 64)
			n6, _ := strconv.ParseInt(c.Query("n6"), 10, 64)
			players = append(players, utils.GetCustomTestPlayer(int(n1), int(n2), int(n3), int(n4), int(n5), int(n6)))
		}
	}

	//add X number of players to the game
	// the base number is 100 by default
	// (in case no query parameter was added)
	if _, err := strconv.ParseInt(c.Query("players"), 10, 64); err == nil {
		numOfPlayers, _ := strconv.ParseInt(c.Query("players"), 10, 64)
		players = append(players, utils.GenerateFakePlayers(int(numOfPlayers))...)
	} else {
		players = append(players, utils.GenerateFakePlayers(100)...)
	}

	pTest, isTest := c.GetQuery("test")
	if isTest && pTest == "true" {
		handlers.ExecuteGame(players)
	} else {
		Quini6Game := handlers.ExecuteGame(players)
		c.IndentedJSON(http.StatusOK, Quini6Game)
	}
}
