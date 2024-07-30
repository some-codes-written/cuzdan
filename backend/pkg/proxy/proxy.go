package proxy

type ProxyRequest struct {
	Url     string `json:"url" example:"https://www.google.com"`
	Payload string `json:"payload" example:"{ \"name\": \"John Doe\" }"`
}
