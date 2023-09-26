package api

import (
	"net/http"

	"github.com/alpsantos/futwatcher-processor/configuration/logger"
	"github.com/alpsantos/futwatcher-processor/configuration/validation"
	"github.com/alpsantos/futwatcher-processor/processor"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
)

type PlayerController struct {
	processorService processor.ProcessorService
}

func NewController(processor processor.ProcessorService) *PlayerController {
	return &PlayerController{
		processorService: processor,
	}
}

func (pc *PlayerController) GetPlayer(c *gin.Context) {
	logger.Info("Init GetCard controller method called")
	//https://www.futwiz.com/en/app/sold23/21767/console
	request := PlayerRequest{}

	if err := c.ShouldBindQuery(&request); err != nil {
		logger.Error("Error on bind query", err)
		errRest := validation.ValidateUserError(err)
		c.JSON(errRest.Code, errRest)
	}

	res, err := pc.processorService.GetPlayerData(request.Id)
	if err != nil {
		errRest := validation.ValidateUserError(err)
		c.JSON(errRest.Code, errRest)
	}

	response := PlayerResponse{}
	copier.Copy(&response, &res)

	c.JSON(http.StatusOK, response)
}
