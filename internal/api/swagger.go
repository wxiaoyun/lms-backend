package api

// Used solely for swagger documentation
type SwgResponse[T any] struct {
	Data     T         `json:"data,omitempty"`
	Meta     Meta      `json:"meta,omitempty"`
	Messages []Message `json:"messages,omitempty"`
}

// Used solely for swagger documentation
type SwgMsgResponse struct {
	Messages []Message `json:"messages,omitempty"`
}

// Used solely for swagger documentation
type SwgErrResponse struct {
	Messages []Message `json:"messages,omitempty"`
	Error    string    `json:"error,omitempty"`
}
