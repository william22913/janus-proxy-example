package janus

import (
	"encoding/json"
	"fmt"
	"janus-proxy-example/constanta"
	"janus-proxy-example/dto/out"
	errorModel "janus-proxy-example/error-model"
	janus_dto_in "janus-proxy-example/services/janus/janus-dto-in"
	janus_dto_out "janus-proxy-example/services/janus/janus-dto-out"
	"janus-proxy-example/services/util"
	"strconv"
)

var JanusHandler handler

type handler struct {}

func (input handler) CreateHandler(transaction string, sessionID int64, plugin string)(output out.CreateHandlerDTOOut, errs errorModel.ErrorModel){
	var httpStatus int
	var bodyResult string
	var err error

	var requestBody = janus_dto_in.HandlerDTOIn{
		SessionDTOIn: janus_dto_in.SessionDTOIn{
			Janus:       constanta.JanusCreateHandler,
			Transaction: transaction,
		},
		Plugin: plugin,
	}

	sessionIDStr := strconv.Itoa(int(sessionID))
	httpStatus, _, bodyResult, err = util.HitAPI("POST", "http://"+constanta.JanusURL+"/"+sessionIDStr, nil, requestBody)
	if err != nil {
		errs = errorModel.GenerateUnknownErrorModel(err)
		return
	}

	if httpStatus == 200 {
		fmt.Println(bodyResult)
		var bodyResultObj janus_dto_out.HandlerDTOOut
		_ = json.Unmarshal([]byte(bodyResult), &bodyResultObj)
		output.SessionID = bodyResultObj.SessionId
		output.HandlerID = bodyResultObj.Data.Id
		errs = errorModel.GenerateNonErrorModel()
	} else{
		errs = errorModel.GenerateUnknownErrorModel(err)
	}

	return
}