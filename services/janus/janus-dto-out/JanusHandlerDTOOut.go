package janus_dto_out

type HandlerDTOOut struct {
	Janus       string `json:"janus"`
	SessionId   int64  `json:"session_id"`
	Transaction string `json:"transaction"`
	Data        struct {
		Id int64 `json:"id"`
	} `json:"data"`
}
