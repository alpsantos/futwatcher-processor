package api

import (
	"github.com/alpsantos/futwatcher-processor/processor"
	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.Engine) {
	processor := processor.NewProcessorService()
	cardController := NewController(processor)

	r.GET("/card", cardController.GetCard)
}
