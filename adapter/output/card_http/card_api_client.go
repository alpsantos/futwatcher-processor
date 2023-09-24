package card_http

import (
	"encoding/json"
	"fmt"

	"github.com/alpsantos/futwatcher-processor/adapter/output/model/response"
	"github.com/alpsantos/futwatcher-processor/application/domain"
	"github.com/alpsantos/futwatcher-processor/configuration/logger"
	"github.com/go-resty/resty/v2"
	"github.com/jinzhu/copier"
)

type cardClient struct{}

func NewCardClient() *cardClient {
	client = resty.New()
	return &cardClient{}
}

var (
	client *resty.Client
)

func (c *cardClient) GetCardPort(request domain.CardRequest) (*domain.Player, error) {

	// client := resty.New()
	//https://www.futwiz.com/en/app/sold24/18343/console

	url := fmt.Sprintf("https://www.futwiz.com/en/app/sold24/%d/console", request.PlayerId)

	clientResponse := &response.PlayerResponse{}
	// clientResponse := domain.Player{}

	jsonData, err := client.R().
		SetHeader("Accept", "application/json").
		Get(url)

	if err != nil {
		return nil, err
	}

	logger.Info(fmt.Sprintf("Card API Response: %s", jsonData.Body()))

	var data map[string]interface{}

	// Decodifica o JSON para o mapa
	if err := json.Unmarshal([]byte(jsonData.Body()), &data); err != nil {
		fmt.Println("Erro ao decodificar JSON:", err)
		return nil, err
	}

	// Agora você pode acessar os dados do mapa conforme necessário
	player := data["player"].(map[string]interface{})

	if _, ok := player["line_id"]; !ok {
		fmt.Println("something went wrong with player id response", err)
		return nil, fmt.Errorf("something went wrong with player id %d response", request.PlayerId)
	}

	prices := data["prices"].(map[string]interface{})["xb"].(map[string]interface{})

	fmt.Println("Player Line ID:", player["line_id"])
	fmt.Println("Player URL Name:", player["urlname"])

	clientResponse.Id = player["line_id"].(string)
	clientResponse.Urlname = player["urlname"].(string)
	clientResponse.Price = prices["binuf"].(string)
	clientResponse.Ud = prices["ud"].(string)

	domainPlayer := &domain.Player{}
	copier.Copy(domainPlayer, clientResponse)
	return domainPlayer, nil
}
