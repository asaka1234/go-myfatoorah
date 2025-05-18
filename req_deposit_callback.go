package go_myfatoorah

import (
	"errors"
	"log"
)

// 充值的回调通知处理, 将之集成到http server里
func (cli *Client) DepositCallback(req MyFatoorahDepositBackReq, processor func(MyFatoorahDepositBackReq) error) error {
	//1. 获取返回
	// Validate request
	if req.Data == nil {
		return errors.New("Invalid request")
	}

	data := req.Data
	// 1. Validate the signature
	dataMap := make(map[string]string)
	dataMap["invoiceId"] = string(data.InvoiceId)
	dataMap["invoiceReference"] = data.InvoiceReference
	dataMap["createdDate"] = data.CreatedDate
	if data.CustomerReference != nil {
		dataMap["customerReference"] = *data.CustomerReference
	}
	dataMap["customerName"] = data.CustomerName
	dataMap["customerMobile"] = data.CustomerMobile
	dataMap["customerEmail"] = data.CustomerEmail
	dataMap["transactionStatus"] = data.TransactionStatus
	dataMap["paymentMethod"] = data.PaymentMethod
	if data.UserDefinedField != nil {
		dataMap["userDefinedField"] = *data.UserDefinedField
	}
	dataMap["referenceId"] = data.ReferenceId
	dataMap["trackId"] = data.TrackId
	dataMap["paymentId"] = data.PaymentId
	dataMap["authorizationId"] = data.AuthorizationId
	dataMap["invoiceValueInBaseCurrency"] = data.InvoiceValueInBaseCurrency
	dataMap["baseCurrency"] = data.BaseCurrency
	dataMap["invoiceValueInDisplayCurreny"] = data.InvoiceValueInDisplayCurrency
	dataMap["displayCurrency"] = data.DisplayCurrency
	dataMap["invoiceValueInPayCurrency"] = data.InvoiceValueInPayCurrency
	dataMap["payCurrency"] = data.PayCurrency

	backServiceKey := cli.AccessKey

	// Generate and validate signature
	if !cli.signUtil.ValidateSignature(dataMap, backServiceKey, req.signature) {
		log.Printf("Invalid signature")
		return errors.New("Invalid signature")
	}

	//2. 实际处理
	return processor(req)
}
