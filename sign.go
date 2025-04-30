package go_myfatoorah

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"sort"
	"strconv"
)

// https://docs.myfatoorah.com/docs/webhook#webhook-signature
// sha256 -> base64
func GenSign(params map[string]interface{}, privateSecret string) string {
	mac := hmac.New(sha256.New, []byte(privateSecret))

	keys := make([]string, 0, len(params))
	for k := range params {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	var buf bytes.Buffer
	for _, k := range keys {
		vs := params[k]
		if vs == "" {
			continue
		}
		if buf.Len() > 0 {
			buf.WriteByte(',')
		}

		buf.WriteString(k)
		buf.WriteByte('=')
		//类型检查
		switch vv := vs.(type) {
		case string:
			buf.WriteString(vv)
		case int:
			buf.WriteString(strconv.FormatInt(int64(vv), 10))
		case int64:
			buf.WriteString(strconv.FormatInt(int64(vv), 10))
		default:
			panic(fmt.Sprintf("params type not supported, k=%s, %+v", k, vv))
		}
	}
	
	fmt.Printf("rawSignStr = %s\n", buf.String())

	_, err := mac.Write([]byte(buf.String()))
	if err != nil {
		return ""
	}
	shaSum := mac.Sum(nil)
	return base64.StdEncoding.EncodeToString(shaSum[:])
}

// 验证签名
func VerifySign(params map[string]interface{}, privateSecret string, sign string) bool {
	//自己算一遍
	selfSign := GenSign(params, privateSecret)
	return selfSign == sign
}
