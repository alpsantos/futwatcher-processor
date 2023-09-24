package input

import "github.com/alpsantos/futwatcher-processor/application/domain"

type CardUseCase interface {
	GetCardService(request domain.CardRequest) (*domain.Player, error)
}
