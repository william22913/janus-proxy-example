package services

import (
	"fmt"
	errorModel "janus-proxy-example/error-model"
	"janus-proxy-example/services/janus"
)

type Session struct{
}

func (input Session) CreateSession(transaction string, _ map[string]interface{}) (interface{}, errorModel.ErrorModel){
	fmt.Println("Hit API Create Session")

	output, err := janus.JanusSession.CreateSession(transaction)
	if err.ErrorCode != 200 {
		return nil, err
	}

	KeepAliveSessionService{output.SessionID}.KeepAliveSession()

	return output, err
}
