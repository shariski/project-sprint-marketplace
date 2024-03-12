package model

type ResponseFormat struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
