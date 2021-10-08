package util

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func HitAPI(method string, url string, header map[string]string, body interface{})(httpStatus int,headerResult map[string][]string, bodyResult string, err error){
	client := &http.Client{}

	bodyStr, _ := json.Marshal(body)
	req, err := http.NewRequest(method, url, bytes.NewBuffer(bodyStr))
	if err != nil {
		return
	}

	if header == nil {
		header = make(map[string]string)
	}

	for key := range header {
		req.Header.Set(key, header[key])
	}

	res, err := client.Do(req)
	if err != nil {
		return
	}

	defer res.Body.Close()
	//Read the response body

	bodyStr, err = ioutil.ReadAll(res.Body)
	if err != nil {
		return
	}

	httpStatus = res.StatusCode
	headerResult = res.Header
	bodyResult = string(bodyStr)
	return
}