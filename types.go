package cryptoix

type AuthMode string

const (
	AuthBearer AuthMode = "bearer"
	AuthXAPIKey AuthMode = "x-api-key"
)

type Config struct {
	APIKey    string
	BaseURL   string
	PublicURL string
	QRURL     string
	AuthMode  AuthMode
}

type Meta struct {
	RequestID  string `json:"request_id,omitempty"`
	Page       int    `json:"page,omitempty"`
	PerPage    int    `json:"per_page,omitempty"`
	Total      int    `json:"total,omitempty"`
	TotalPages int    `json:"total_pages,omitempty"`
}

type apiEnvelope struct {
	OK      *bool          `json:"ok,omitempty"`
	Success *bool          `json:"success,omitempty"`
	Data    any            `json:"data,omitempty"`
	Meta    Meta           `json:"meta,omitempty"`
	Error   map[string]any `json:"error,omitempty"`
	Errors  []string       `json:"errors,omitempty"`
}
