package go_myfatoorah

import (
	"errors"
	"github.com/asaka1234/go-myfatoorah/utils"
	"github.com/mitchellh/mapstructure"
	"log"
)

// 其中sign是从header读取出来的
func (cli *Client) DepositCallback(req MyFatoorahDepositBackReq, sign string, processor func(MyFatoorahDepositBackReq) error) error {
	//1. 获取返回
	if req.Data == nil {
		return errors.New("Invalid request")
	}

	//自己算一下签名
	var params map[string]interface{}
	mapstructure.Decode(req.Data, &params)

	// Generate and validate signature
	// 注意: 这里要在psp后台先打开这个设置
	if !utils.Verify(params, cli.Params.BackKey, sign) {
		log.Printf("Invalid signature")
		return errors.New("Invalid signature")
	}

	//2. 实际处理
	return processor(req)
}
