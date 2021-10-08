package janus_dto_in

type HandlerDTOIn struct{
	SessionDTOIn
	Plugin string `json:"plugin"`
}
