package response

type PlayerResponse struct {
	Id      string `json:"player_id"`
	Urlname string `json:"urlname"`
	Price   string `json:"price"`
	Ud      string `json:"ud"`
}
