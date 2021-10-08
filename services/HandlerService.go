package services

import (
	"fmt"
	"janus-proxy-example/constanta"
	errorModel "janus-proxy-example/error-model"
	"janus-proxy-example/services/janus"
)

type Handler struct{}

func (input Handler) CreateHandler(transaction string, inputStruct map[string]interface{}) (interface{}, errorModel.ErrorModel) {
	fmt.Println("Hit API Create Handler")

	err := input.validateCreateHandler(inputStruct)
	if err.ErrorCode != 200 {
		return nil, err
	}

	output, err := janus.JanusHandler.CreateHandler(transaction, inputStruct[constanta.FieldSessionID].(int64), constanta.JanusVideoRoomPlugin)
	if err.ErrorCode != 200 {
		return nil, err
	}

	return output, err
}

func (input Handler) validateCreateHandler(inputData map[string]interface{}) (err errorModel.ErrorModel) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
			err = errorModel.GenerateJSONInvalidModel()
		}
	}()

	if inputData[constanta.FieldSessionID] == nil {
		return errorModel.GenerateEmptyFieldError("Session ID")
	}

	inputData[constanta.FieldSessionID] = int64(inputData[constanta.FieldSessionID].(float64))
	if inputData[constanta.FieldSessionID].(int64) == 0 {
		return errorModel.GenerateEmptyFieldError("Session ID")
	}

	return errorModel.GenerateNonErrorModel()
}
