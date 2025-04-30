package go_myfatoorah

import "github.com/go-resty/resty/v2"

type Client struct {
	Token    string
	BaseURL  string
	ryClient *resty.Client
}

func NewClient(token string, baseURL string) *Client {
	return &Client{
		Token:    token,
		BaseURL:  baseURL,
		ryClient: resty.New(), //client实例
	}
}
