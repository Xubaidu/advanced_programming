package schema

type SendMailRequest struct {
	From        string
	To          string
	Subject     string
	ContentType string
	Body        string
}

type SendMailResponse struct {
	Msg string `json:"msg"`
}
