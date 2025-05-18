package go_myfatoorah

import (
	"github.com/asaka1234/go-myfatoorah/utils"
	"github.com/go-resty/resty/v2"
)

type Client struct {
	Merchant           string
	AccessKey          string
	DepositUrl         string
	DepositCallbackUrl string

	ryClient *resty.Client
	logger   utils.Logger
}

func NewClient(logger utils.Logger, merchantID string, accessKey string, depositUrl, depositCallbackUrl string) *Client {
	return &Client{
		Merchant:           merchantID,
		AccessKey:          accessKey,
		DepositUrl:         depositUrl,
		DepositCallbackUrl: depositCallbackUrl,

		ryClient: resty.New(), //client实例
		logger:   logger,
	}
}

func (cli *Client) SetMerchant(merchantId string) {
	cli.Merchant = merchantId
}
