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
	var players []models.Player

	//check if we want to include a custom player in the game
	param, exists := c.GetQuery("customPlayer")
	if exists && param == "true" {
		players = append(players, utils.GetCustomTestPlayer())
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

	Quini6Game := handlers.ExecuteGame(players)
	c.IndentedJSON(http.StatusOK, Quini6Game)
}
