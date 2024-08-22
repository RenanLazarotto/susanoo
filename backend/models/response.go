package models

type Meta struct {
	Count int `json:"count"`
}

type Response struct {
	Meta  Meta        `json:"meta,omitempty"`
	Data  interface{} `json:"data,omitempty"`
	Error interface{} `json:"errors,omitempty"`
}
