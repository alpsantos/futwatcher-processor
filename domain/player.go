package domain

type Player struct {
	Id      string
	Urlname string
	Price   string
	Ud      string
}

type PlayerHistory struct {
	Player  Player    `json:"player"`
	History []History `json:"history"`
}

type History struct {
	Timestamp string `json:"timestamp"`
	Price     int    `json:"price"`
}
