package api

type ApiResponse struct {
	Data     interface{} `json:"data;omitempty"`
	Meta     interface{} `json:"meta;omitempty"`
	Messages []string    `json:"messages;omitempty"`
	Error    string      `json:"error;omitempty"`
}
