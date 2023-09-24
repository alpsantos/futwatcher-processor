package main

import (
	"github.com/alpsantos/futwatcher-processor/api"
	"github.com/alpsantos/futwatcher-processor/configuration/logger"

	"github.com/gin-gonic/gin"
)

func main() {
	logger.Info("Starting the application...")
	router := gin.Default()

	api.InitRoutes(router)

	if err := router.Run(":8080"); err != nil {
		logger.Error("Error while running the application", err)
		return
	}
}
