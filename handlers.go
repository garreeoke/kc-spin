package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
)

// Status codes
const (
	timeFormat = "01-02-2006"
	success    = 200
	failure    = 400
)

// Healthz return ok without auth for testing connectivity
func Healthz(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(200)
	data := []byte("OK")
	w.Write(data)

}

// PodEfficiency ... get efficiency for a
func PodEfficiency(w http.ResponseWriter, r *http.Request) {

	payload := Payload{
		Code:       success,
		W:          w,
		Error:      nil,
	}

	var ef EfficiencyRec
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		payload.processError(err, "reading body")
	}

	err = json.Unmarshal(body, &ef)
	if err != nil {
		payload.processError(err, "unmarshall body")
	}

	err = ef.PodRecommendation()
	if err != nil {
		payload.processError(err, "pod recommendation")
	}

	payload.Data, err = json.Marshal(ef)
	if err != nil {
		payload.processError(err, "payload data marshall")
	}

	payload.completeRequest()
}

func (p *Payload) completeRequest() {
	// Check for error message and end accordingly
	if p.Error != nil {
		p.Code = failure
		log.Println("Error: ", p.Error)
	}

	p.W.Header().Set("Content-Type", "application/json; charset=UTF-8")
	p.W.WriteHeader(p.Code)
	p.W.Write(p.Data)
}

func (p *Payload) processError(err error, msg string) {
	p.Error = errors.New(msg + ":" + err.Error())
	p.completeRequest()
}
