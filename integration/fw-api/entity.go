package fw_api

type PlayerRequest struct {
	Id int
}

type PlayerResponse struct {
	Id      string
	Urlname string
	Price   string
	Ud      string
}

type PlayerHistoryResponse struct {
	Console map[int64]float64
	PC      map[int64]float64
}
