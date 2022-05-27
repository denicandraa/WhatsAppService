package models

type ParamSendText struct {
	Data []DataSendText `json:"data" form:"data" binding:"required"`
}

type DataSendText struct {
	Number  string `json:"number" form:"number" binding:"required"`
	Message string `json:"message" form:"message" binding:"required"`
}
