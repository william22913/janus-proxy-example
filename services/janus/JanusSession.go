package janus

import (
	"encoding/json"
	"janus-proxy-example/constanta"
	"janus-proxy-example/dto/out"
	errorModel "janus-proxy-example/error-model"
	janus_dto_in "janus-proxy-example/services/janus/janus-dto-in"
	janus_dto_out "janus-proxy-example/services/janus/janus-dto-out"
	"janus-proxy-example/services/util"
)

var JanusSession session

type session struct {

}
func (input session) CreateSession(transaction string)(output out.CreateSessionDTOOut, errs errorModel.ErrorModel){
	var httpStatus int
	var bodyResult string
	var err error

	var requestBody = janus_dto_in.SessionDTOIn{
		Janus:       constanta.JanusCreateSession,
		Transaction: transaction,
	}

	httpStatus, _, bodyResult, err = util.HitAPI("POST", "http://"+constanta.JanusURL, nil, requestBody)
	if err != nil {
		errs = errorModel.GenerateUnknownErrorModel(err)
		return
	}

	if httpStatus == 200{
		var bodyResultObj janus_dto_out.SessionDTOOut
		_ = json.Unmarshal([]byte(bodyResult), &bodyResultObj)
		output.SessionID = bodyResultObj.Data.Id
		errs = errorModel.GenerateNonErrorModel()
	} else{
		errs = errorModel.GenerateUnknownErrorModel(err)
	}

	return
}