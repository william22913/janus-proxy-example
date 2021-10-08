package services

import (
	"fmt"
	errorModel "janus-proxy-example/error-model"
	"janus-proxy-example/model"
)

var (
	serviceList = serviceInit()
)

func serviceInit() (output map[string]func(string, map[string]interface{}) (interface{},errorModel.ErrorModel)){
	output = make(map[string]func(string, map[string]interface{}) (interface{},errorModel.ErrorModel))
	output["session-create"] = Session{}.CreateSession
	output["handler-create"] = Handler{}.CreateHandler

	return output
}

func StartService(messageIn model.SocketMessage) (interface{},errorModel.ErrorModel){
	fmt.Println("Start Service")
	if serviceList[messageIn.Request] != nil {
		return serviceList[messageIn.Request](messageIn.Transaction, messageIn.Body)
	} else {
		return nil, errorModel.GenerateUnsupportedServiceModel()
	}
}


