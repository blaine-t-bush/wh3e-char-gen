package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env.development")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	gin.SetMode(os.Getenv("GIN_MODE"))
	router := gin.New()        // Creates a router without any middleware
	router.Use(gin.Logger())   // Writes logs to gin.DefaultWriter, by default os.Stdout
	router.Use(gin.Recovery()) // Recovers from panics and instead delivers a 500

	router.GET("/character/new", func(c *gin.Context) {
		character := GenerateCharacter()
		c.JSON(http.StatusOK, character)
	})

	// listen and serve on http://0.0.0.0:8080, or http://localhost:8080 on windows
	router.Run(fmt.Sprintf(":%s", os.Getenv("PORT")))
}
