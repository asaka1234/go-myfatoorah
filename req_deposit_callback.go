package go_myfatoorah

// 充值的回调通知处理, 将之集成到http server里
func (cli *Client) DepositCallback(req MyFatoorahDepositBackReq, processor func(MyFatoorahDepositBackReq) error) error {
	//1. 获取返回
	//TODO 鉴权

	//2. 实际处理
	return processor(req)
}
