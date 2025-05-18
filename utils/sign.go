package utils

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"sort"
	"strings"
)

type SignatureUtils struct {
	logger Logger
}

func NewSignatureUtils(logger Logger) *SignatureUtils {
	return &SignatureUtils{logger: logger}
}

// callback的数据的签名
func (s *SignatureUtils) GenerateSignature(data map[string]string, key string) (string, error) {
	// 1. Sort the data keys case-insensitively
	keys := make([]string, 0, len(data))
	for k := range data {
		keys = append(keys, k)
	}

	sort.Slice(keys, func(i, j int) bool {
		return strings.ToLower(keys[i]) < strings.ToLower(keys[j])
	})

	// 2. Create a string from the sorted data
	var dataString strings.Builder
	for _, k := range keys {
		v := data[k]
		if v != "" { // In Go, empty string is the zero value
			dataString.WriteString(k)
			dataString.WriteString("=")
			dataString.WriteString(v)
			dataString.WriteString(",")
		}
	}

	// Remove the trailing comma
	str := dataString.String()
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

func (s *SignatureUtils) ValidateSignature(data map[string]string, key, signature string) bool {
	generatedSignature, err := s.GenerateSignature(data, key)
	if err != nil {
		s.logger.Errorf("Error generating signature: %v", err)
		return false
	}

	return generatedSignature == signature
}
