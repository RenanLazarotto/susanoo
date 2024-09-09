package models

type ResponseData interface {
	interface{} | []interface{}
}

type Response struct {
	Count   int          `json:"count,omitempty"`
	Data    ResponseData `json:"data,omitempty"`
	Errors  []string     `json:"errors,omitempty"`
	Message string       `json:"message,omitempty"`
}
