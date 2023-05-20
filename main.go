package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	r.GET("/", func(ctx *gin.Context) {
		ctx.IndentedJSON(http.StatusOK, gin.H{
			"message": "successful",
		})

	})
	r.Run(":" + port)

}
