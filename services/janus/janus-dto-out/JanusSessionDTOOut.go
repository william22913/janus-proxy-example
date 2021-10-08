package janus_dto_out

type SessionDTOOut struct {
	Janus       string `json:"janus"`
	Transaction string `json:"transaction"`
	Data        struct {
		Id int64 `json:"id"`
	} `json:"data"`
}