package go_myfatoorah

import (
	"crypto/tls"
	"fmt"
	"github.com/asaka1234/go-myfatoorah/utils"
	jsoniter "github.com/json-iterator/go"
	"github.com/mitchellh/mapstructure"
)

// https://docs.myfatoorah.com/docs/send-payment
// https://apitest.myfatoorah.com/swagger/ui/index#!/Payment/Payment_SendPayment
func (cli *Client) Deposit(req MyFatoorahDepositReq) (*MyFatoorahDepositRsp, error) {

	cli.logger.Infof("go_myfatoorah==>deposit, req:%+v", req)

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

	rawURL := cli.Params.DepositUrl

	var params map[string]interface{}
	mapstructure.Decode(req, &params)

	params["NotificationOption"] = "LNK"
	params["WebhookUrl"] = cli.Params.DepositCallbackUrl //指定webhook, 这样会直接走这个
	params["Language"] = "EN"                            //临时写死

	//----------------------
	var result MyFatoorahDepositRsp

	resp, err := cli.ryClient.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true}).
		SetCloseConnection(true).
		R().
		SetBody(params).
		SetHeaders(getAuthHeaders(cli.Params.ApiToken)).
		SetDebug(cli.debugMode).
		SetResult(&result).
		SetError(&result).
		Post(rawURL)

	restLog, _ := jsoniter.ConfigCompatibleWithStandardLibrary.Marshal(utils.GetRestyLog(resp))
	cli.logger.Infof("PSPResty#myfatoorah#deposit->%+v", string(restLog))

	if err != nil {
		return nil, err
	}

	if resp.StatusCode() != 200 {
		//反序列化错误会在此捕捉
		return nil, fmt.Errorf("status code: %d", resp.StatusCode())
	}

	if resp.Error() != nil {
		//反序列化错误会在此捕捉
		return nil, fmt.Errorf("%v, body:%s", resp.Error(), resp.Body())
	}

	return &result, nil
}
