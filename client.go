package go_myfatoorah

import (
	"github.com/asaka1234/go-myfatoorah/utils"
	"github.com/go-resty/resty/v2"
)

type Client struct {
	Merchant           string
	ApiToken           string
	BackKey            string
	DepositUrl         string
	DepositCallbackUrl string

	ryClient *resty.Client
	logger   utils.Logger
}

func NewClient(logger utils.Logger, merchantID string, apiToken string, backKey string, depositUrl, depositCallbackUrl string) *Client {
	return &Client{
		Merchant:           merchantID,
		ApiToken:           apiToken,
		BackKey:            backKey,
		DepositUrl:         depositUrl,
		DepositCallbackUrl: depositCallbackUrl,

		ryClient: resty.New(), //client实例
		logger:   logger,
	}
}

func (cli *Client) SetMerchantInfo(merchantId, apiToken, backKey string) {
	cli.Merchant = merchantId
	cli.ApiToken = apiToken
	cli.BackKey = backKey
}
