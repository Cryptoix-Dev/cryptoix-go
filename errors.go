package cryptoix

import "fmt"

type APIError struct {
	StatusCode int
	Code       string
	Message    string
	Details    map[string]any
	RequestID  string
	RawBody    []byte
}

func (e *APIError) Error() string {
	if e.Code != "" {
		return fmt.Sprintf("cryptoix: %s (%s)", e.Message, e.Code)
	}
	return fmt.Sprintf("cryptoix: %s", e.Message)
}

type RateLimitError struct {
	*APIError
	RetryAfter int
}
