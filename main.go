package main

import (
	"q6-golang/api"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/play", api.Play)
	router.Run("127.0.0.1:10000")
}
