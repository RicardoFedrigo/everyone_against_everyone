package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ricardofedrigo/everyone_against_everyone/internal/service"
)

type DeckController struct {
}

func NewDeckController() *DeckController {
	return &DeckController{}
}

func (c *DeckController) GetDeck(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, service.ReturnDeck())
}
