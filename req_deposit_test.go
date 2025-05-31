package go_myfatoorah

import (
	"fmt"
	"testing"
	"time"
)

func TestDeposit(t *testing.T) {

	//构造client
	cli := NewClient(nil, MyFatoorahInitParams{MerchantInfo{MERCHANT, API_TOKEN, BACK_KEY}, DEPOSIT_URL, DEPOSIT_CALLBACK_URL})

	//发请求
	resp, err := cli.Deposit(GenDepositRequestDemo())
	if err != nil {
		fmt.Printf("err:%s\n", err.Error())
		return
	}
	fmt.Printf("resp:%+v\n", resp)
}

func GenDepositRequestDemo() MyFatoorahDepositReq {
	return MyFatoorahDepositReq{
		CustomerName:       "cy",
		NotificationOption: "LNK",
		DisplayCurrencyIso: "KWD",
		MobileCountryCode:  "+971",
		InvoiceValue:       50,
		ExpiryDate:         time.Now().AddDate(0, 0, 1),
		CallBackUrl:        "https://www.google.com",
		Language:           "EN",
	}
}
