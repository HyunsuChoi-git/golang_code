package http

import (
	"crypto/tls"
	"github.com/rs/zerolog/log"
	"io"
	"net/http"
	"time"
)

// NewClientInit Rest API(HTTP1.*) 연결 공통 로직 - 커넥션 정보 설정
func NewClientInit() *http.Client {
	var defaultClient *http.Client

	defaultTransportPointer, ok := http.DefaultTransport.(*http.Transport)
	if !ok {
		log.Error().Msgf("defaultRoundTripper not an *http.Transport")
	}
	defaultTransport := *defaultTransportPointer
	defaultTransport.IdleConnTimeout = 90 * time.Second
	defaultTransport.MaxIdleConns = 100
	defaultTransport.MaxIdleConnsPerHost = 100

	tr := &http.Transport{
		IdleConnTimeout:     90 * time.Second,
		MaxIdleConns:        100,
		MaxIdleConnsPerHost: 100,
		TLSClientConfig:     &tls.Config{InsecureSkipVerify: true},
	}

	defaultClient = &http.Client{
		Timeout:   time.Second * 40,
		Transport: tr,
	}

	return defaultClient
}

// ExecuteService Rest API(HTTP1.*) 연결 공통 로직 - 메서드 별 구분
func ExecuteService(req *http.Request) (http.Header, []byte, int, error) {
	client := NewClientInit()

	// request service_api
	resp, err := client.Do(req)
	if err != nil {
		log.Error().Msgf("Http Client Do Error: %s", err)
		return nil, nil, 0, err
	}

	// check response status
	if resp.StatusCode == http.StatusForbidden {
		log.Error().Msgf("Http Client Response StatusCode Error: %s", resp.StatusCode)
		return nil, nil, resp.StatusCode, nil
	}

	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Error().Msgf("ReadIo Error: %s", err)
	}

	return resp.Header, respBody, resp.StatusCode, nil
}
