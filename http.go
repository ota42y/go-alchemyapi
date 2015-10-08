package alchemyapi

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

var apiBase = "http://gateway-a.watsonplatform.net/"

type apiConnection interface {
	get(endPoint string, params url.Values, config *config) ([]byte, error)
	post(endPoint string, params url.Values, config *config) ([]byte, error)
}

// http interface
type httpImp struct {
}

func (h *httpImp) get(endPoint string, params url.Values, config *config) ([]byte, error) {
	return h.connection("GET", endPoint, params, config)
}

func (h *httpImp) post(endPoint string, params url.Values, config *config) ([]byte, error) {
	return h.connection("POST", endPoint, params, config)
}

func (h *httpImp) connection(method string, endPoint string, params url.Values, config *config) ([]byte, error) {
	if config == nil || config.apikey == "" {
		return make([]byte, 0), fmt.Errorf("No auth token")
	}
	params.Add("apikey", config.apikey)
	params.Add("outputMode", "json")

	req, _ := http.NewRequest(method, apiBase+endPoint, nil)
	req.URL.RawQuery = params.Encode()

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return make([]byte, 0), err
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}
