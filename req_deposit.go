package go_myfatoorah

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// 下单(充值/提现是同一个接口)
func (cli *Client) Deposit(req MyFatoorahDepositReq) (*MyFatoorahDepositRsp, error) {

	/**
	 * curl -X POST \
	 * --header 'Content-Type: application/json' \
	 * --header 'Accept: application/json' \
	 * --header 'Authorization: bearer CoUYHpUydszMZ8-vSVVM-kHVbO_aTTosv9zJqBPSibJ5T7nR8GJKPJVPo8zfllE8SDBX9UUZcoX_7OR0Vmv2tW5KrhRAAZkdP4zmbPzUxYb32sWcBKZus_HX-cUuuPK5nBKCYNsYATlkUzS7FDeAI3O2Yz8fzCLyUwCNoDw0v6lMADweivaR4R9rIX3Kp1DpkDY0cv9DwTh74IaCtfWH1F2n_KRy2ndf1xaI_b3dsTpL7-z-rAmfSpoF8QlvlSnhx0R4o_FDLQMnOI5tvjcK6Rgh-kOO0fnA1CmzPYjHXVLiAEa1nPt9giCqzralL2j47fiLszvTeMWy6nb-dXXDGrC9IKEYkUyQuvGQZMs2uuWyC2UvIORDBHZHAO4PLjuaUMB5uqyxXXHJndfPikFGeCtANigotEA4Fz3ptV-riR_mXS3ftpDwAnBOXBj2M5HNsfkC4T4NZkNoxJm9rNVI07DsNtIjW2UGd3zg-DCsTFOVz_qPNVwx7YH3W6DrBovY2CxVQszuqjrLTS6vZ4nInpkUlK--41ruhHZjtIQIx9PsIz38iXbhIBIVkP2xV6HpttH-xzCCake-4vX8nzjTWAdqnpAN_Z_oH9liE_N36Qs__gBs5T6hWPYLJRsX8BiMAe12VUo_ZL2JVJJqlyxxnoQCk2ALQtK-J95ITcE3PXyPYWi9zL8QydnbM1irC4QXMfDZVQ' \
	 * --data '{
	 *    "CustomerName": "larry",
	 *    "NotificationOption": "LNK",
	 *    "DisplayCurrencyIso": "AED",
	 *    "InvoiceValue": 1,
	 *    "ExpiryDate": "2024-06-13T05:07:37.715Z",
	 *    "CallBackUrl": "https://userportal.cptinternational.com/en/user/login",
	 *    "Language" : "en"
	 *  }' \
	 * 'https://apitest.myfatoorah.com/v2/SendPayment'
	 */

	rawURL := cli.DepositUrl + "/v2/SendPayment"

	// Prepare request body
	requestBody := map[string]interface{}{
		"CustomerName":       req.CustomerName,
		"NotificationOption": req.NotificationOption,
		"DisplayCurrencyIso": req.DisplayCurrencyIso,
		"InvoiceValue":       req.InvoiceValue,
		"Language":           "en",
		// Uncomment these if needed
		// "ExpiryDate":       req.ExpiryDate,
		// "CallBackUrl":      backUrl,
	}

	// Convert request body to JSON
	jsonBody, err := json.Marshal(requestBody)
	if err != nil {
		return nil, fmt.Errorf("error marshaling request body: %v", err)
	}

	// Determine access key based on currency
	accessKey := cli.AccessKey

	// Create HTTP request
	httpReq, err := http.NewRequest("POST", rawURL, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}

	// Set headers
	httpReq.Header.Set("Content-Type", "application/json")
	httpReq.Header.Set("Accept", "application/json")
	httpReq.Header.Set("Authorization", "bearer "+accessKey)

	// Log request (you might want to use a proper logging library)
	cli.logger.Infof("MyFatoorah Deposit request: %+v, headers: %+v\n", requestBody, httpReq.Header)

	// Send request
	client := &http.Client{Timeout: 30 * time.Second}
	resp, err := client.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %v", err)
	}
	defer resp.Body.Close()

	// Read response
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response: %v", err)
	}

	// Parse response
	var rsp MyFatoorahDepositRsp
	err = json.Unmarshal(body, &rsp)
	if err != nil {
		return nil, fmt.Errorf("error parsing response: %v", err)
	}

	// Handle non-successful responses
	if !rsp.IsSuccess {
		// Try to extract validation errors
		var errorResponse struct {
			ValidationErrors *string `json:"ValidationErrors"`
		}
		if err := json.Unmarshal(body, &errorResponse); err == nil {
			rsp.ValidationErrors = errorResponse.ValidationErrors
		}
	}

	return &rsp, nil
}
