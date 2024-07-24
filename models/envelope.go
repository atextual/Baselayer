package models

type ResponseEnvelope struct {
	Message string `json:"message"`
	Data    any    `json:"data"`
}
