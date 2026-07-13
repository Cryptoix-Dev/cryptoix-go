package cryptoix

import (
	"io"
	"context"
	"net/http"
	"strings"
	"testing"
)

type roundTripFunc func(*http.Request) (*http.Response, error)

func (f roundTripFunc) RoundTrip(r *http.Request) (*http.Response, error) { return f(r) }

func TestAuthHeaderAndSuccessEnvelope(t *testing.T) {
	transport := roundTripFunc(func(r *http.Request) (*http.Response, error) {
		if r.Header.Get("Authorization") != "Bearer ak_live_xxxxxxxxxxxx" { t.Fatalf("missing auth header") }
		return &http.Response{StatusCode: 200, Header: http.Header{"Content-Type": []string{"application/json"}}, Body: io.NopCloser(strings.NewReader(`{"ok":true,"data":{"hello":"world"},"meta":{"request_id":"req_1"}}`))}, nil
	})
	client := NewClient("ak_live_xxxxxxxxxxxx", WithBaseURL("https://api.test/v1"), WithHTTPClient(&http.Client{Transport: transport}))
	data, err := client.ListBalances(context.Background())
	if err != nil { t.Fatal(err) }
	m := data.(map[string]any)
	if m["hello"] != "world" { t.Fatalf("unexpected data: %#v", data) }
}

func TestAPIError(t *testing.T) {
	transport := roundTripFunc(func(r *http.Request) (*http.Response, error) {
		return &http.Response{StatusCode: 403, Header: http.Header{"Content-Type": []string{"application/json"}}, Body: io.NopCloser(strings.NewReader(`{"ok":false,"error":{"code":"missing_scope","message":"Nope","request_id":"req_2"}}`))}, nil
	})
	client := NewClient("ak_live_xxxxxxxxxxxx", WithBaseURL("https://api.test/v1"), WithHTTPClient(&http.Client{Transport: transport}))
	_, err := client.ListBalances(context.Background())
	if err == nil { t.Fatal("expected error") }
	apiErr, ok := err.(*APIError)
	if !ok || apiErr.Code != "missing_scope" || apiErr.RequestID != "req_2" { t.Fatalf("bad error: %#v", err) }
}
