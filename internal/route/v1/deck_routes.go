package route

import (
	"github.com/gin-gonic/gin"
	"github.com/ricardofedrigo/everyone_against_everyone/internal/controller"
)

func SetupDeckRoutes(router *gin.Engine) {
	deckController := controller.NewDeckController()

	router.GET("/deck", deckController.GetDeck)
}
