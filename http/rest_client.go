package http

import (
	"bytes"
	"encoding/json"
	"github.com/rs/zerolog/log"
	"net/http"
	"strings"
)

func GetData(apiUri string, headers map[string]string, mappingData interface{}) (http.Header, int, error) {
	return callExecute(setRequest("GET", apiUri, headers, nil), mappingData)
}

func PostData(apiUri string, headers map[string]string, mappingData interface{}, reqJson *bytes.Buffer) (http.Header, int, error) {
	return callExecute(setRequest("POST", apiUri, headers, reqJson), mappingData)
}

func setRequest(method string, apiUri string, headers map[string]string, reqJson *bytes.Buffer) *http.Request {
	fnc := "setRequest"

	var req *http.Request
	var err error

	if strings.ToUpper(method) == "GET" || strings.ToUpper(method) == "DELETE" {
		req, err = http.NewRequest("GET", apiUri, nil)
	} else {
		req, err = http.NewRequest("POST", apiUri, reqJson)
	}
	if err != nil {
		log.Error().Msgf("%s : To create new request is failed. - error : %s", fnc, err)
	}

	// define header
	req.Header.Add("Content-Type", "application/json")

	for key, value := range headers {
		req.Header.Add(key, value)
	}

	return req
}

func callExecute(req *http.Request, mappingData interface{}) (http.Header, int, error) {
	fnc := "callExecute"

	respHeader, respBody, statusCode, err := ExecuteService(req)
	if err != nil || !(statusCode == http.StatusOK || statusCode == http.StatusCreated) {
		return nil, statusCode, err
	}

	if err := json.Unmarshal(respBody, &mappingData); err != nil {
		log.Error().Msgf("%s : Unmarshaling Fail. - error : %s", fnc, err)
		return nil, statusCode, err
	}

	return respHeader, statusCode, nil
}
