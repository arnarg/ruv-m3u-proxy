package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/arnarg/ruv-m3u-proxy/logging"
)

type RuvResponse struct {
	GeoBlock bool   `json:"geoblock"`
	URL      string `json:"url"`
}

type RuvHandler struct {
	URL string
}

func NewRuvHandler(url string) RuvHandler {
	return RuvHandler{URL: url}
}

func (h RuvHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	logging.LogRequest(r)
	body, err := getRequest(h.URL, r)
	if err != nil {
		w.WriteHeader(502)
	}

	ruvResponse := RuvResponse{}

	err = json.Unmarshal(body, &ruvResponse)
	if err != nil {
		w.WriteHeader(502)
	}

	http.Redirect(w, r, ruvResponse.URL, 302)
	fmt.Printf("Handed URL %s to %s\n", ruvResponse.URL, r.RemoteAddr)
}

func getRequest(url string, r *http.Request) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	// Pass on some headers from the requester
	req.Header = http.Header{
		"User-Agent": r.Header["User-Agent"],
		"Accept":     r.Header["Accept"],
	}

	httpClient := &http.Client{Timeout: time.Second * 10}

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
