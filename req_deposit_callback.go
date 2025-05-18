package go_myfatoorah

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// 充值的回调通知处理, 将之集成到http server里
func DepositNotify(w http.ResponseWriter, r *http.Request, processor func(MyFatoorahDepositBackReq) error) error {

	//1. 获取返回
	var req MyFatoorahDepositBackReq
	err := json.NewDecoder(r.Body).Decode(&req) // 解析响应体为JSON格式
	if err != nil {
		fmt.Println("解析JSON失败：", err)
		return err
	}

	//2. 实际处理
	return processor(req)
}
