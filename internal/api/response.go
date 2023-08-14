// nolint revive
package api

type Response struct {
	Data     interface{} `json:"data,omitempty"`
	Meta     interface{} `json:"meta,omitempty"`
	Messages []Message   `json:"messages,omitempty"`
	Error    string      `json:"error,omitempty"`
}
