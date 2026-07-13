package cryptoix

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

const (
	DefaultBaseURL = "https://api.cryptoix.io/v1"
	DefaultPublicURL = "https://cryptoix.io"
	DefaultQRURL = "https://qr.cryptoix.io"
)

type Client struct {
	apiKey string
	baseURL string
	publicURL string
	qrURL string
	authMode AuthMode
	httpClient *http.Client
}

func NewClient(apiKey string, opts ...func(*Client)) *Client {
	c := &Client{apiKey: apiKey, baseURL: DefaultBaseURL, publicURL: DefaultPublicURL, qrURL: DefaultQRURL, authMode: AuthBearer, httpClient: &http.Client{Timeout: 30*time.Second}}
	for _, opt := range opts { opt(c) }
	return c
}

func WithBaseURL(baseURL string) func(*Client) { return func(c *Client) { c.baseURL = strings.TrimRight(baseURL, "/") } }
func WithHTTPClient(h *http.Client) func(*Client) { return func(c *Client) { c.httpClient = h } }
func WithAuthMode(mode AuthMode) func(*Client) { return func(c *Client) { c.authMode = mode } }

func (c *Client) do(ctx context.Context, method, path string, body any, query url.Values, auth bool) (any, error) {
	if query != nil && query.Get("api_key") != "" { return nil, fmt.Errorf("cryptoix: do not send API keys in query parameters; use headers") }
	var reader io.Reader
	if body != nil {
		b, err := json.Marshal(body); if err != nil { return nil, err }
		reader = bytes.NewReader(b)
	}
	reqURL := c.baseURL + "/" + strings.TrimLeft(path, "/")
	if query != nil && len(query) > 0 { reqURL += "?" + query.Encode() }
	req, err := http.NewRequestWithContext(ctx, method, reqURL, reader); if err != nil { return nil, err }
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", "cryptoix-go/0.1.0")
	if body != nil { req.Header.Set("Content-Type", "application/json") }
	if auth {
		if c.authMode == AuthXAPIKey { req.Header.Set("X-API-Key", c.apiKey) } else { req.Header.Set("Authorization", "Bearer "+c.apiKey) }
	}
	resp, err := c.httpClient.Do(req); if err != nil { return nil, err }
	defer resp.Body.Close()
	raw, _ := io.ReadAll(resp.Body)
	var env apiEnvelope
	if len(raw) > 0 { _ = json.Unmarshal(raw, &env) }
	if (env.OK != nil && !*env.OK) || (env.Success != nil && !*env.Success) || resp.StatusCode >= 400 {
		return nil, buildError(resp, raw, env)
	}
	return env.Data, nil
}

func buildError(resp *http.Response, raw []byte, env apiEnvelope) error {
	apiErr := &APIError{StatusCode: resp.StatusCode, Message: http.StatusText(resp.StatusCode), RawBody: raw, RequestID: resp.Header.Get("X-Request-Id")}
	if env.Error != nil {
		if v, ok := env.Error["code"].(string); ok { apiErr.Code = v }
		if v, ok := env.Error["message"].(string); ok { apiErr.Message = v }
		if v, ok := env.Error["request_id"].(string); ok { apiErr.RequestID = v }
		if v, ok := env.Error["details"].(map[string]any); ok { apiErr.Details = v }
	}
	if resp.StatusCode == 429 { retry, _ := strconv.Atoi(resp.Header.Get("Retry-After")); return &RateLimitError{APIError: apiErr, RetryAfter: retry} }
	return apiErr
}

func fmtInt(n int) string { return strconv.Itoa(n) }
