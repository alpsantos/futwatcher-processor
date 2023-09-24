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
