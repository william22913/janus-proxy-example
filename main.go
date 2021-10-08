package main

import (
	"encoding/json"
	"fmt"
	"janus-proxy-example/dto/out"
	errorModel "janus-proxy-example/error-model"
	"janus-proxy-example/model"
	"janus-proxy-example/services"
	"net/http"
	"strconv"

	ws "github.com/gorilla/websocket"
)

func main() {
	http.HandleFunc("/ws", webSocketEndpoint)
	port := strconv.Itoa(8000)
	fmt.Println("Server Start in port " + port)
	fmt.Println(http.ListenAndServe("0.0.0.0:"+port, nil))
}

var upgrade = ws.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func corsOriginHandler(response http.ResponseWriter) {
	response.Header().Set("Access-Control-Allow-Origin", "*")
}

func serveOutput(response http.ResponseWriter, errs errorModel.ErrorModel) {
	corsOriginHandler(response)

	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(errs.ErrorCode)

	_ = json.NewEncoder(response).Encode(out.StandardOutput{
		Success: false,
		Payload: errs.Message,
	})

	if errs.CausedBy != nil {
		fmt.Println(errs.CausedBy)
	}
}

func webSocketEndpoint(response http.ResponseWriter, request *http.Request) {
	var errs errorModel.ErrorModel

	defer func() {
		serveOutput(response, errs)
	}()

	upgrade.CheckOrigin = func(r *http.Request) bool {
		return true
	}

	socket, err := upgrade.Upgrade(response, request, nil)
	if err != nil {
		errs = errorModel.GenerateUnknownErrorModel(err)
		return
	}

	fmt.Println("Client Connected")
	reader(socket)
}

func reader(conn *ws.Conn) {
	for {
		msgType, data, err := conn.ReadMessage()
		if err != nil {
			fmt.Println(err)
			return
		}

		var message model.SocketMessage

		err = json.Unmarshal(data, &message)
		if err != nil {
			fmt.Println(err)
			continue
		}

		var output = out.StandardOutput{
			Success: true,
			Request: message.Request,
		}

		payload, errs := services.StartService(message)
		if errs.ErrorCode != 200 {
			output = out.StandardOutput{
				Transaction: message.Transaction,
				Success: false,
				Payload: errs.Message,
			}
		} else {
			output.Payload = payload
		}

		data, _ = json.Marshal(output)

		if err = conn.WriteMessage(msgType, data); err != nil {
			fmt.Println(err)
			return
		}
	}

}
