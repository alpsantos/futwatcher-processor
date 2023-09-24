package routes

import (
	"github.com/alpsantos/futwatcher-processor/adapter/input/controller"
	"github.com/alpsantos/futwatcher-processor/adapter/output/card_http"
	"github.com/alpsantos/futwatcher-processor/application/service"
	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.Engine) {

	cardClient := card_http.NewCardClient()
	cardService := service.NewCardService(cardClient)

	cardController := controller.NewCardController(cardService)

	r.GET("/cards", cardController.GetCard)
}
