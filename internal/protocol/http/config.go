package http

type config struct {
	UserAgent   string `json:"userAgent"`
	Connections int    `json:"connections"`
}
