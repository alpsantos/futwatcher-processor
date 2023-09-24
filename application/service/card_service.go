package service

import (
	"fmt"

	"github.com/alpsantos/futwatcher-processor/application/domain"
	"github.com/alpsantos/futwatcher-processor/application/port/output"
	"github.com/alpsantos/futwatcher-processor/configuration/logger"
)

type cardService struct {
	cardPort output.GetCardPort
}

func NewCardService(cardPort output.GetCardPort) *cardService {
	return &cardService{
		cardPort: cardPort,
	}
}

func (s *cardService) GetCardService(request domain.CardRequest) (*domain.Player, error) {
	logger.Info(fmt.Sprintf("Getting card %d", request.PlayerId))

	 cardDomainResponse, err := s.cardPort.GetCardPort(request)
	
	 return cardDomainResponse, err
}
