package client

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/dp0h/wc-api-go/options"
)

// A store that stops answering must not hold a request forever: the client gives up after the
// configured timeout.
func TestNewClientAppliesTheRequestTimeout(t *testing.T) {
	release := make(chan struct{})
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		<-release
	}))
	defer server.Close()
	defer close(release)

	factory := Factory{}
	c := factory.NewClient(options.Basic{
		URL:    server.URL,
		Key:    "ck_test",
		Secret: "cs_test",
		Options: options.Advanced{
			WPAPI:       true,
			WPAPIPrefix: "/wp-json/",
			Version:     "wc/v3",
			Timeout:     1,
		},
	})

	start := time.Now()
	resp, err := c.Get("orders", nil)
	if resp != nil {
		_ = resp.Body.Close()
	}
	if err == nil {
		t.Fatal("expected a timeout error from a store that never answers")
	}
	if elapsed := time.Since(start); elapsed > 5*time.Second {
		t.Fatalf("request took %s; the 1 s timeout was not applied", elapsed)
	}
}
