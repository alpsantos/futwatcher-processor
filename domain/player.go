package domain

type Player struct {
	Id      string
	Urlname string
	Price   string
	Ud      string
}

type PlayerHistory struct {
	Player  Player            `json:"player"`
	History map[int64]float64 `json:"history"`
}
