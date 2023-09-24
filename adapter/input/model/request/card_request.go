package request

type CardRequest struct {
	Id int `form:"id" json:"id" binding:"required"`
	// From string `form:"from" binding:"required, datetime" time_format:"2006-01-02"`
}
