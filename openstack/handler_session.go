package openstack

import (
	"bytes"
	"encoding/json"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"golang_code/openstack/os_api_model"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"
)

var Openstacks = map[string]openstack{}
var mu sync.Mutex

type openstack struct {
	url  string
	auth struct {
		url      string
		domain   string
		user     string
		password string
		reqBody  []byte
	}
	Token struct {
		Token     string
		expiresAt time.Time
	}

	catalogs []os_api_model.ResAuthTokenCatalog
}

func GetAuthToken(provId string) error {
	fnc := "GetAuthToken"
	err := error(nil)

	if _, ok := Openstacks[provId]; !ok {
		mu.Lock()
		Openstacks[provId] = openstack{}
		mu.Unlock()
	}

	o := Openstacks[provId]
	// 토큰 발급주기 체크
	//	* 확인하고자 하는 계정이 저장되어있는 계정과 같은 지 판단 후 발급주기 체크.
	//		- 글로벌로 저장된 계정과 같은 경우 토큰 발급 주기 체크 후 주기에 따라 토큰 재발급
	//		- 글로벌로 저장된 계정과 다른 경우 토큰 재발급 및 글로벌 변수에 계정 새로 저장
	if strings.Contains(o.auth.url, endpoint) &&
		o.auth.domain == domain &&
		o.auth.user == user &&
		o.auth.password == password {

		if o.Token.Token != "" && time.Now().Before(o.Token.expiresAt.Add(-5*time.Minute)) {
			return nil
		}

	} else {
		o.url = endpoint
		o.auth.url = endpoint + ":" + keystonePort + "/auth/tokens"
		o.auth.domain = domain
		o.auth.user = user
		o.auth.password = password
		err = o.setReqBodyForAuthToken(&o.auth.reqBody)
	}

	// call service_api
	var authTokenBody os_api_model.ResAuthBody
	resHeader, statusCode, err := PostData(o.auth.url, nil, &authTokenBody, bytes.NewBuffer(o.auth.reqBody))
	if err != nil {
		return err
	}

	// check response status. 200 or 201
	if !(statusCode == http.StatusOK || statusCode == http.StatusCreated) {
		statusMessage := strconv.Itoa(statusCode) + " : " + http.StatusText(statusCode)
		err = errors.New(statusMessage)
		log.Error().Msgf("%s : The statuscode is not normal. - error : %s", fnc, err)
		return err
	}

	// extract auth token in header
	tokenList := []string{resHeader.Get("X-Subject-Token")}

	t := ""
	if tokenList == nil || len(tokenList) == 0 {
		log.Error().Msgf("%s : no token in the header. - error : %s", fnc, err)
	} else {
		for _, token := range tokenList {
			t = token
		}
	}

	if authTokenBody.Token.Catalog == nil {
		return errors.New("getEndpoint() Failed. catalog.Endpoints is nil")
	}

	o.catalogs = authTokenBody.Token.Catalog

	exTime := authTokenBody.Token.Expires_at.Sub(authTokenBody.Token.Issued_at)

	o.Token.Token = t
	o.Token.expiresAt = time.Now().Add(exTime).Add(-2 * time.Minute)

	Openstacks[provId] = o
	return nil
}

func GetEndpoint(provId string, catalogType string, catalogName string) (string, error) {
	o := Openstacks[provId]

	// extract endpoints to need
	endpoints := []os_api_model.ResAuthTokenCatalogEndPoints{}
	for _, catalog := range o.catalogs {
		if catalog.TypeName == catalogType && catalog.Name == catalogName {
			endpoints = catalog.Endpoints
			break
		}
	}

	if endpoints == nil {
		return "", errors.New("GetEndpoint() Failed. catalog.Endpoints is nil")
	}

	// extract a valid URL
	for _, endpoint := range endpoints {
		if strings.ToLower(endpoint.InterfaceName) == "public" {
			//return endpoint.Url, nil
			//port만 보내고 kafka에서 받은 url과 조합
			return o.url + ":" + strings.Split(strings.Split(endpoint.Url, ":")[2], "/")[0], nil
		}
	}

	return "", errors.New("GetEndpoint() Failed. endpoint.Url is nil")
}

func (o *openstack) setReqBodyForAuthToken(reqBody *[]byte) error {
	fnc := "createReqBodyForAuthToken"
	err := error(nil)
	/*
		{
		    "auth": {
		        "identity": {
		            "methods": [
		                "password"
		            ],
		            "password": {
		                "user": {
		                    "name": "admin",
		                    "domain": {
		                        "name": "default"
		                    },
		                    "password": "okestro2018"
		                }
		            }
		        },
		        "scope": {
		            "system": {
		                "all": true
		            }
		        }
			}
		}
	*/
	req := map[string]interface{}{
		"auth": map[string]interface{}{
			"identity": map[string]interface{}{
				"methods": []string{"password"},
				"password": map[string]interface{}{
					"user": map[string]interface{}{
						"name": o.auth.user,
						"domain": map[string]interface{}{
							"name": o.auth.domain},
						"password": o.auth.password},
				},
			},
			"scope": map[string]interface{}{
				"system": map[string]interface{}{
					"all": true,
				},
			},
		},
	}

	*reqBody, err = json.Marshal(req)
	if err != nil {
		log.Error().Msgf("%s : To get auth token is Failed. - error : %s", fnc, err)
		return err
	}

	return nil
}
