package services

type KeepAliveSessionService struct {
	SessionID int64
}

func (input KeepAliveSessionService) KeepAliveSession() {
	// todo go routine keep alive session
}
