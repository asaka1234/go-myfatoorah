package go_myfatoorah

import (
	"fmt"
	"testing"
)

func TestDepositCallback(t *testing.T) {
	vLog := VLog{}

	req := MyFatoorahDepositBackReq{
		EventType:      1,
		Event:          "TransactionsStatusChanged",
		DateTime:       "17062025190705",
		CountryIsoCode: "KWT",
		Data: &MyFatoorahDepositBackReqData{
			InvoiceId:                     5838959,
			InvoiceReference:              "2025000197",
			CreatedDate:                   "17062025190433",
			CustomerReference:             "23452222",
			CustomerName:                  "cy",
			CustomerMobile:                "+965",
			CustomerEmail:                 nil,
			TransactionStatus:             "FAILED",
			PaymentMethod:                 "VISA/MASTER",
			UserDefinedField:              nil,
			ReferenceId:                   "07075838959279331675",
			TrackId:                       "17-06-2025_2793316",
			PaymentId:                     "07075838959279331675",
			AuthorizationId:               "07075838959279331675",
			InvoiceValueInBaseCurrency:    "50",
			BaseCurrency:                  "KWD",
			InvoiceValueInDisplayCurrency: "50",
			DisplayCurrency:               "KWD",
			InvoiceValueInPayCurrency:     "50",
			PayCurrency:                   "KWD",
		},
	}

	sign := "voK/CKpQOKQkCiqpw8xuyWCX6Ws9iIVS7DYiYQ/5Olk="

	//构造client
	cli := NewClient(vLog, &MyFatoorahInitParams{MerchantInfo{MERCHANT, API_TOKEN, BACK_KEY}, DEPOSIT_URL, DEPOSIT_CALLBACK_URL})

	//发请求
	err := cli.DepositCallback(req, sign, processor)
	if err != nil {
		fmt.Printf("err:%s\n", err.Error())
		return
	}
}

func processor(MyFatoorahDepositBackReq) error {
	return nil
}
