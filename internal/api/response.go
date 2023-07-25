package api

type Response struct {
	//nolint:revive // By design
	Data interface{} `json:"data,omitempty"`
	//nolint:revive // By design
	Meta     interface{} `json:"meta,omitempty"`
	Messages []Message   `json:"messages,omitempty"`
	Error    string      `json:"error,omitempty"`
}
