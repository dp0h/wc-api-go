package net

import (
	"github.com/dp0h/wc-api-go/request"
	"net/http"
	"strings"
)

// Sender provides HTTP Requests
type Sender struct {
	requestEnricher RequestEnricher
	urlBuilder      URLBuilder
	httpClient      Client
	requestCreator  RequestCreator
}

// Send method sends requests to WooCommerce API
func (s *Sender) Send(req request.Request) (resp *http.Response, err error) {
	request := s.prepareRequest(req)
	return s.httpClient.Do(request)
}

func (s *Sender) prepareRequest(req request.Request) *http.Request {
	URL := s.urlBuilder.GetURL(req)
	hasBody := ("POST" == req.Method || "PUT" == req.Method) && len(req.Body) > 0
	var request *http.Request
	if hasBody {
		reader := strings.NewReader(req.Body)
		request, _ = s.requestCreator.NewRequest(req.Method, URL, reader)
		request.Header.Add("Content-type", "application/json")
	} else {
		request, _ = s.requestCreator.NewRequest(req.Method, URL, nil)
	}
	s.requestEnricher.EnrichRequest(request, URL)
	if req.Values != nil {
		request.Form = req.Values
	}
	return request
}

// SetRequestEnricher ...
func (s *Sender) SetRequestEnricher(a RequestEnricher) {
	s.requestEnricher = a
}

// SetURLBuilder ...
func (s *Sender) SetURLBuilder(urlBuilder URLBuilder) {
	s.urlBuilder = urlBuilder
}

// SetHTTPClient ...
func (s *Sender) SetHTTPClient(c Client) {
	s.httpClient = c
}

// SetRequestCreator ...
func (s *Sender) SetRequestCreator(rc RequestCreator) {
	s.requestCreator = rc
}
