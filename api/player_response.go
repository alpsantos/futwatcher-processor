package api

type PlayerResponse struct {
	Id      string `json:"id"`
	Urlname string `json:"urlname"`
	Price   string `json:"price"`
	Ud      string `json:"ud"`
}
