package model

type SocketMessage struct {
	Request string `json:"request"`
	Transaction string `json:"transaction"`
	Body map[string]interface{} `json:"body"`
}