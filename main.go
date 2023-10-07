package main

import (
	"net/http"
	core "q6-golang/core"
	"strconv"

	"github.com/gin-gonic/gin"
)

func startGame(c *gin.Context) {
	testPlayer := core.GetPlayer("Juanito Perez", 20, "Rosario", "Calle Falsa 123", "3415559999", core.GetTicket([]int{1, 5, 6, 12, 25, 41}, core.GP_TradicionalAndRevancha))
	players := []core.Player{testPlayer}
	if _, err := strconv.ParseInt(c.Query("randomPlayers"), 10, 64); err == nil {
		numOfPlayers, _ := strconv.ParseInt(c.Query("randomPlayers"), 10, 64)
		players = append(players, core.GenerateFakePlayers(int(numOfPlayers))...)
	} else {
		players = append(players, core.GenerateFakePlayers(2500)...)
	}
	//q6 := core.ExecuteGame(players) --> this already returns a bunch of info we could hand back to FE via API response
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
