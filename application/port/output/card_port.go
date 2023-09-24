package output

import "github.com/alpsantos/futwatcher-processor/application/domain"

type GetCardPort interface {
	GetCardPort(domain.CardRequest) (*domain.Player, error)
}
