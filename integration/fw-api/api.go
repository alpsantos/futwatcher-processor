package fw_api

import (
	"encoding/json"
	"fmt"

	"github.com/alpsantos/futwatcher-processor/configuration/logger"

	"github.com/go-resty/resty/v2"
)

type FwDataClient interface {
	GetPlayerDetails(request PlayerRequest) (*PlayerResponse, error)
	GetPlayerHistory(request PlayerRequest) (*PlayerHistoryResponse, error)
}

type FwDataClientImpl struct{}

const (
	GetPlayerDataUrl    string = "https://www.futwiz.com/en/app/sold24/%d/console"
	GetPlayerHistoryUrl string = "https://www.futwiz.com/en/app/price_history_player24_multi?p=%d&h"
)

var (
	httpClient = resty.New()
)

func (c *FwDataClientImpl) GetPlayerHistory(request PlayerRequest) (*PlayerHistoryResponse, error) {

	requestUrl := fmt.Sprintf(GetPlayerHistoryUrl, request.Id)

	data := &PlayerHistoryResponse{}

	jsonData, err := httpClient.R().
		SetHeader("Accept", "application/json").
		Get(requestUrl)

	if err != nil {
		return nil, err
	}

	logger.Info(fmt.Sprintf("Player API Response: %s", jsonData.Body()))

	if err := json.Unmarshal([]byte(jsonData.Body()), &data); err != nil {
		logger.Error("Error at unmarshal player data", err)
		return nil, err
	}

	return data, nil
}

func (c *FwDataClientImpl) GetPlayerDetails(request PlayerRequest) (*PlayerResponse, error) {

	requestUrl := fmt.Sprintf(GetPlayerDataUrl, request.Id)

	playerResponse := &PlayerResponse{}

	jsonData, err := httpClient.R().
		SetHeader("Accept", "application/json").
		Get(requestUrl)

	if err != nil {
		return nil, err
	}

	logger.Info(fmt.Sprintf("Player API Response: %s", jsonData.Body()))

	var data map[string]interface{}

	if err := json.Unmarshal([]byte(jsonData.Body()), &data); err != nil {
		logger.Error("Error at unmarshal player data", err)
		return nil, err
	}

	player := data["player"].(map[string]interface{})

	if _, ok := player["line_id"]; !ok {
		logger.Error("Maybe a not found player", err)
		return nil, fmt.Errorf("Maybe a not found player id %d", request.Id)
	}

	prices := data["prices"].(map[string]interface{})["xb"].(map[string]interface{})

	playerResponse.Id = player["line_id"].(string)
	playerResponse.Urlname = player["urlname"].(string)
	playerResponse.Price = prices["binuf"].(string)
	playerResponse.Ud = prices["ud"].(string)

	return playerResponse, nil
}
