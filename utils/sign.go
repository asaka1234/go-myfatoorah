package utils

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/samber/lo"
	"github.com/spf13/cast"
	"sort"
	"strings"
)

// 针对webhook回调的签名验证
// https://docs.myfatoorah.com/docs/webhook-signature
func Sign(params map[string]interface{}, key string) (string, error) {
	// 1. Validate key
	if key == "" {
		return "", errors.New("APP_KEY 参数为空，请填写")
	}

	// 2. Get and sort keys
	keys := lo.Keys(params)
	sort.Strings(keys) // ASCII ascending order

	// 3. Build sign string
	var sb strings.Builder
	for _, k := range keys {
		value := cast.ToString(params[k])
		if value != "" {
			//只有非空才可以参与签名
			sb.WriteString(fmt.Sprintf("%s=%s,", k, value))
		}
	}

	// Remove the trailing comma
	str := sb.String()
	if len(str) > 0 {
		str = str[:len(str)-1]
	}

	// 3. Generate HMAC SHA-256 hash
	mac := hmac.New(sha256.New, []byte(key))
	mac.Write([]byte(str))
	hashBytes := mac.Sum(nil)

	// 4. Encode the hash to Base64
	return base64.StdEncoding.EncodeToString(hashBytes), nil
}

// 验证签名
func Verify(data map[string]interface{}, key, signature string) bool {
	generatedSignature, err := Sign(data, key)
	if err != nil {
		fmt.Errorf("Error generating signature: %v", err)
		return false
	}
	return generatedSignature == signature
}
