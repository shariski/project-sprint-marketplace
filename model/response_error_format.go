package model

type ResponseErrorFormat struct {
	Message string `json:"message"`
	Errors  string `json:"errors"`
}
