package go_myfatoorah

import (
	"github.com/asaka1234/go-myfatoorah/utils"
	"github.com/go-resty/resty/v2"
)

type Client struct {
	Params *MyFatoorahInitParams

	ryClient  *resty.Client
	debugMode bool
	logger    utils.Logger
}

func NewClient(logger utils.Logger, params *MyFatoorahInitParams) *Client {
	return &Client{
		Params: params,

		ryClient:  resty.New(), //client实例
		debugMode: false,
		logger:    logger,
	}
}

func (cli *Client) SetMerchantInfo(merchant MerchantInfo) {
	cli.Params.MerchantInfo = merchant
}

func (cli *Client) SetDebugModel(debugModel bool) {
	cli.debugMode = debugModel
}
