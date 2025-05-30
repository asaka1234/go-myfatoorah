package go_myfatoorah

import (
	"github.com/asaka1234/go-myfatoorah/utils"
	"github.com/go-resty/resty/v2"
)

type Client struct {
	Params MyFatoorahInitParams

	ryClient *resty.Client
	logger   utils.Logger
}

func NewClient(logger utils.Logger, params MyFatoorahInitParams) *Client {
	return &Client{
		Params: params,

		ryClient: resty.New(), //client实例
		logger:   logger,
	}
}

func (cli *Client) SetMerchantInfo(merchantId, apiToken, backKey string) {
	cli.Params.Merchant = merchantId
	cli.Params.ApiToken = apiToken
	cli.Params.BackKey = backKey
}
