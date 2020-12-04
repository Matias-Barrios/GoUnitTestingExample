package services

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
)

type HHResult struct {
	Result     string
	StatusCode int
	Error      error
}

type IHttpHelper interface {
	Request(method string, url string, headers map[string][]string, payload interface{}) chan HHResult
	RequestWithOptions(method string, url string, headers map[string][]string, payload interface{}) chan HHResult
}

var singleton IHttpHelper

type HttpHelper struct{}

func (h *HttpHelper) Request(method string, url string, headers map[string][]string, payload interface{}) chan HHResult {
	var result chan HHResult = make(chan HHResult)
	go func(output chan HHResult) {
		payl, err := json.Marshal(payload)
		if err != nil {
			result <- HHResult{
				Result:     "",
				StatusCode: 0,
				Error:      err,
			}
		}
		req, err := http.NewRequest(method, url, bytes.NewReader(payl))
		client := http.Client{
			Timeout: time.Second,
		}
		resp, err := client.Do(req)
		if err != nil {
			result <- HHResult{
				Result:     "",
				StatusCode: 0,
				Error:      err,
			}
		}
		respbytes, err := ioutil.ReadAll(resp.Body)
		defer resp.Body.Close()
		result <- HHResult{
			Result:     string(respbytes),
			StatusCode: resp.StatusCode,
			Error:      nil,
		}

	}(result)
	return result
}

func (h *HttpHelper) RequestWithOptions(method string, url string, headers map[string][]string, payload interface{}) chan HHResult {
	var result chan HHResult = make(chan HHResult)
	go func(output chan<- HHResult) {
		payl, err := json.Marshal(payload)
		if err != nil {
			result <- HHResult{
				Result:     "",
				StatusCode: 0,
				Error:      err,
			}
		}
		req, err := http.NewRequest(method, url, bytes.NewReader(payl))
		client := http.Client{
			Timeout: time.Second,
		}
		resp, err := client.Do(req)
		if err != nil {
			result <- HHResult{
				Result:     "",
				StatusCode: 0,
				Error:      err,
			}
		}
		respbytes, err := ioutil.ReadAll(resp.Body)
		defer resp.Body.Close()
		result <- HHResult{
			Result:     string(respbytes),
			StatusCode: resp.StatusCode,
			Error:      nil,
		}

	}(result)
	return result
}

func NewHttpHelper() IHttpHelper {
	if singleton == nil {
		singleton = &HttpHelper{}
	}
	return singleton
}
