package go_myfatoorah

import (
	"github.com/asaka1234/go-myfatoorah/utils"
	"github.com/go-resty/resty/v2"
)

type Client struct {
	params MyFatoorahInitParams

	ryClient *resty.Client
	logger   utils.Logger
}

func NewClient(logger utils.Logger, params MyFatoorahInitParams) *Client {
	return &Client{
		params: params,

		ryClient: resty.New(), //client实例
		logger:   logger,
	}
}

func (cli *Client) SetMerchantInfo(merchantId, apiToken, backKey string) {
	cli.params.Merchant = merchantId
	cli.params.ApiToken = apiToken
	cli.params.BackKey = backKey
}
