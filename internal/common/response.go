package common

type Response struct {
	Code int      `json:"code"`
	Data Data     `json:"data,omitempty"`
	Msg  string   `json:"msg,omitempty"`
	Err  []string `json:"error,omitempty"` // only available in debug mode
}

type Data map[string]any
