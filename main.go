package main

import (
	"net/http"
	core "q6-golang/core"
	"q6-golang/enumerators"
	"strconv"

	"github.com/gin-gonic/gin"
)

func startGame(c *gin.Context) {
	testPlayer := core.GetPlayer("Juanito Perez", 20, "Rosario", "Calle Falsa 123", "3415559999", core.GetTicket([]int{1, 5, 6, 12, 25, 41}, enumerators.GP_TradicionalAndRevancha))
	players := []core.Player{testPlayer}
	if _, err := strconv.ParseInt(c.Query("randomPlayers"), 10, 64); err == nil {
		numOfPlayers, _ := strconv.ParseInt(c.Query("randomPlayers"), 10, 64)
		players = append(players, core.GenerateFakePlayers(int(numOfPlayers))...)
	} else {
		players = append(players, core.GenerateFakePlayers(2500)...)
	}
	core.ExecuteGame(players)
	tradicional := core.GenerateDrawing()
	segunda := core.GenerateDrawing()
	revancha := core.GenerateDrawing()
	siempreSale := core.GenerateDrawing()
	var response = map[string]core.DrawingResult{
		"#1 - Tradicional":  tradicional,
		"#2 - Segunda":      segunda,
		"#3 - Revancha":     revancha,
		"#4 - Siempre Sale": siempreSale,
	}
	c.IndentedJSON(http.StatusOK, response)
}

func main() {
	router := gin.Default()
	router.GET("/play", startGame)
	router.Run("127.0.0.1:10000")
}
