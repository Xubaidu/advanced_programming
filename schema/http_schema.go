package schema

type HttpMethodRequest struct {
	Method string            `json:"method"`
	URL    string            `json:"url"`
	Header map[string]string `json:"header"`
	Params interface{}       `json:"params"`
}

type HttpRequest struct {
	Method      string      `json:"method"`
	URL         string      `json:"url"`
	ContentType string      `json:"content_type"`
	Token       string      `json:"token"`
	ReqParams   interface{} `json:"req_params"`
}
