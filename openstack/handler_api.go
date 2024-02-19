package openstack

import (
	"errors"
	"github.com/rs/zerolog/log"
	"net/http"
	"strconv"
	"strings"
)

func GetMetrics(provId string, catalogType string, catalogName string, path string, response interface{}) error {
	fnc := "GetMetricsAll"
	// get endpoint

	endpoint, err := GetEndpoint(provId, catalogType, catalogName)
	if err != nil {
		log.Error().Msgf("%s : To get endpoint is failed.. - error : %s", fnc, err)
		return err
	}

	if strings.EqualFold(endpoint, "") {
		return errors.New("no endpoint!")
	}

	header := map[string]string{}

	if o, ok := Openstacks[provId]; ok && o.Token.Token != "" {
		header["X-Auth-Token"] = o.Token.Token
	}

	_, statusCode, err := GetData(endpoint+path, header, &response)

	if err != nil {
		return err
	}
	if statusCode != http.StatusOK {
		err = errors.New(strconv.Itoa(statusCode) + " : " + http.StatusText(statusCode))
		return err
	}

	return nil
}
