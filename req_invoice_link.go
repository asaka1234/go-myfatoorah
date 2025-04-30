package go_myfatoorah

import (
	"crypto/tls"
)

// 获取一个invoice link
// https://docs.myfatoorah.com/docs/send-payment
func (cli *Client) SendPayment(req SendPaymentRequest) (*SendPaymentResponse, error) {

	reqPath := "/v2/SendPayment"
	rawURL := cli.BaseURL + reqPath

	//返回值会放到这里
	var result SendPaymentResponse

	_, err := cli.ryClient.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true}).
		SetCloseConnection(true).
		R().
		SetBody(req).
		SetHeaders(getAuthHeaders(cli.Token)).
		SetResult(&result).
		Post(rawURL)

	//fmt.Printf("accessToken: %+v\n", resp)

	if err != nil {
		return nil, err
	}

	return &result, err
}
