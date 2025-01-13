package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/ricardofedrigo/everyone_against_everyone/internal/route/v1"
	// "github.com/ricardofedrigo/everyone_against_everyone/internal/db"
)

func main() {
	// db.InitDb()
	// db.MigrateDB()
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.Use(gin.Recovery())

	route.SetupDeckRoutes(router)

	log.Println(os.Getenv("PORT"))

	if err := router.Run(":3000"); err != nil {
		log.Fatalln("Error starting server")
		panic("Server not started")
	}

}
