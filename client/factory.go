package client

import (
	"crypto/tls"
	"github.com/dp0h/wc-api-go/auth"
	"github.com/dp0h/wc-api-go/net"
	"github.com/dp0h/wc-api-go/options"
	"github.com/dp0h/wc-api-go/url"
	"net/http"
)

// Factory Structure
type Factory struct {
}

// NewClient method creates new Client
func (f *Factory) NewClient(o options.Basic) Client {
	authenticator := f.NewAuthenticator(o)

	urlBuilder := url.Builder{}
	urlBuilder.SetOptions(o)
	urlBuilder.SetQueryEnricher(authenticator)

	sender := f.NewSender(urlBuilder, o)
	c := Client{
		sender: &sender,
	}
	return c
}

// NewSender method creates new Sender
func (f *Factory) NewSender(u url.Builder, o options.Basic) net.Sender {
	tr := &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}
	httpClient := http.Client{Transport: tr}
	requestCreator := f.NewRequestCreator()
	requestEnricher := f.NewAuthenticator(o)

	sender := net.Sender{}
	sender.SetURLBuilder(&u)
	sender.SetHTTPClient(&httpClient)
	sender.SetRequestCreator(&requestCreator)
	sender.SetRequestEnricher(requestEnricher)
	return sender
}

// NewRequestCreator ...
func (f *Factory) NewRequestCreator() net.HTTP {
	return net.HTTP{}
}

// NewAuthenticator ...
func (f *Factory) NewAuthenticator(o options.Basic) *auth.Authenticator {
	oauth := auth.OAuth{}
	oauth.SetMicrotimer(&auth.MicroTimer{})

	ba := auth.BasicAuthentication{}

	authenticator := auth.Authenticator{}
	authenticator.SetOAuth(oauth)
	authenticator.SetBasicAuth(ba)
	authenticator.SetOptions(o)

	return &authenticator
}
