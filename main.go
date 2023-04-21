package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()
	router.GET("/hello", world)

	if err := router.Run("localhost:8080"); err != nil {
		panic(err)
	}
}

func world(ctx *gin.Context) {
	ctx.String(http.StatusOK, "world")
}
