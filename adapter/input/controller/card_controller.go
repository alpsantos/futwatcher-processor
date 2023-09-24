package controller

import (
	"net/http"

	"github.com/alpsantos/futwatcher-processor/adapter/input/model/request"
	"github.com/alpsantos/futwatcher-processor/application/domain"
	"github.com/alpsantos/futwatcher-processor/application/port/input"
	"github.com/alpsantos/futwatcher-processor/configuration/logger"
	"github.com/alpsantos/futwatcher-processor/configuration/validation"
	"github.com/gin-gonic/gin"
)

type cardController struct {
	cardUseCase input.CardUseCase
}

func NewCardController(cardUseCase input.CardUseCase) *cardController {
	return &cardController{
		cardUseCase: cardUseCase,
	}
}

func (cc *cardController) GetCard(c *gin.Context) {
	logger.Info("Init GetCard controller method called")
	//https://www.futwiz.com/en/app/sold23/21767/console
	request := request.CardRequest{}

	if err := c.ShouldBindQuery(&request); err != nil {
		logger.Error("Error on bind query", err)
		errRest := validation.ValidateUserError(err)
		c.JSON(errRest.Code, errRest)
	}

	cardRequest := domain.CardRequest{
		PlayerId: request.Id,
	}

	response, err := cc.cardUseCase.GetCardService(cardRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}

	c.JSON(http.StatusOK, response)
}
