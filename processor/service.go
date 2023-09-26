package processor

import (
	"github.com/alpsantos/futwatcher-processor/domain"
	fw_api "github.com/alpsantos/futwatcher-processor/integration/fw-api"
	"github.com/jinzhu/copier"
)

type ProcessorService interface {
	GetPlayerData(playerId int) (domain.Player, error)
}

type ProcessorServiceImpl struct {
	dataClient fw_api.FwDataClient
}

func NewProcessorService() *ProcessorServiceImpl {
	return &ProcessorServiceImpl{
		dataClient: &fw_api.FwDataClientImpl{},
	}
}

func (p *ProcessorServiceImpl) GetPlayerData(playerId int) (response domain.Player, err error) {

	request := fw_api.PlayerRequest{
		Id: playerId,
	}

	res, err := p.dataClient.GetPlayerDetails(request)
	if err != nil {
		return
	}

	player := domain.Player{}
	copier.Copy(&player, &res)

	return player, nil
}

func (p *ProcessorServiceImpl) GetPlayerHistory(playerId int) (response domain.PlayerHistory, err error) {

	request := fw_api.PlayerRequest{
		Id: playerId,
	}

	res, err := p.dataClient.GetPlayerHistory(request)
	if err != nil {
		return
	}

	history := domain.PlayerHistory{}
	copier.Copy(&history, &res)

	return history, nil
}
