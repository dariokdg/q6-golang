package main

import (
	"net/http"
	core "q6-golang/core"
	"strconv"

	"github.com/gin-gonic/gin"
)

func startGame(c *gin.Context) {
	testPlayer := core.GetPlayer("Juanito Perez", 20, "Rosario", core.GetTicket([]int{1, 5, 6, 12, 25, 41}, core.GP_TradicionalAndRevancha))
	players := []core.Player{testPlayer}
	if _, err := strconv.ParseInt(c.Query("players"), 10, 64); err == nil {
		numOfPlayers, _ := strconv.ParseInt(c.Query("players"), 10, 64)
		players = append(players, core.GenerateFakePlayers(int(numOfPlayers))...)
	} else {
		players = append(players, core.GenerateFakePlayers(2500)...)
	}
	q6 := core.ExecuteGame(players)
	c.IndentedJSON(http.StatusOK, q6)
}

func main() {
	router := gin.Default()
	router.GET("/play", startGame)
	router.Run("127.0.0.1:10000")
}
