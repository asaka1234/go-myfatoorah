package go_myfatoorah

import "time"

type MyFatoorahInitParams struct {
	MerchantInfo
	DepositUrl         string `json:"depositUrl" mapstructure:"depositUrl" config:"depositUrl"  yaml:"depositUrl"`
	DepositCallbackUrl string `json:"depositCallbackUrl" mapstructure:"depositCallbackUrl" config:"depositCallbackUrl"  yaml:"depositCallbackUrl"`
}

type MerchantInfo struct {
	MerchantId string `json:"MerchantId" mapstructure:"merchantId" config:"merchantId"  yaml:"merchantId"`
	ApiToken   string `json:"apiToken" mapstructure:"apiToken" config:"apiToken"  yaml:"apiToken"`
	BackKey    string `json:"backKey" mapstructure:"backKey" config:"backKey"  yaml:"backKey"`
}

// ----------generate invoice link-------------------------

/**
 * {
 *     "CustomerName": "cy",
 *     "NotificationOption": "LNK",
 *     "DisplayCurrencyIso": "KWD",
 *     "MobileCountryCode": "+971",
 *     "InvoiceValue": 50,
 *     "ExpiryDate": "2024-06-13T05:07:37.715Z"
 *     "CallBackUrl": "https://www.google.com",
 *     "Language" : "en"
 * }
 */
type MyFatoorahDepositReq struct {
	//must
	InvoiceValue float64 `json:"InvoiceValue"  mapstructure:"InvoiceValue"` // 支付金额
	CustomerName string  `json:"CustomerName"  mapstructure:"CustomerName"` //merchant的客户名字
	//option
	DisplayCurrencyIso string `json:"DisplayCurrencyIso"  mapstructure:"DisplayCurrencyIso"` //币种
	CustomerReference  string `json:"CustomerReference"  mapstructure:"CustomerReference"`   //商户订单号
	//MobileCountryCode  string    `json:"MobileCountryCode,omitempty"` // marked as optional
	//ExpiryDate  time.Time `json:"ExpiryDate"`  //The date you want the invoice link to expire
	//CallBackUrl string `json:"CallBackUrl"` //TODO 应该是前端跳转地址

	//以下让sdk来设置
	//NotificationOption string `json:"NotificationOption"` //枚举, EML,SMS,LNK,ALL  (应该选用LNK :returns the invoice link through the response body only)
	//WebhookUrl string `json:"WebhookUrl"` //结果通知.  这个是一个单独补充, 正常是dashboard后台设置的.
	//Language string `json:"Language"` //EN,AR (Arabic)

}

/**
 * 返回数据
 * {
 *     "IsSuccess": true,
 *     "Message": "Invoice Created Successfully!",
 *     "ValidationErrors": null,
 *     "Data": {
 *         "InvoiceId": 6220128,
 *         "InvoiceURL": "https://portal.myfatoorah.com/ARE/ie/050719108622012861-44c3fe51",
 *         "CustomerReference": null,
 *         "UserDefinedField": null
 *     }
 * }
 */
type MyFatoorahDepositRsp struct {
	IsSuccess        bool                   `json:"IsSuccess"`
	Message          string                 `json:"Message"`
	ValidationErrors *string                `json:"ValidationErrors"` // Using pointer to allow null
	Data             *MyFatoorahDepositData `json:"Data"`             // Using pointer to allow null
}

type MyFatoorahDepositData struct {
	InvoiceId         int     `json:"InvoiceId"` //psp三方的订单id (发票号)
	InvoiceURL        string  `json:"InvoiceURL"`
	CustomerReference *string `json:"CustomerReference"` // 商户订单号
	UserDefinedField  *string `json:"UserDefinedField"`  // The custom field that you have passed in the request
}

// ----------webhook-------------------------
// https://docs.myfatoorah.com/docs/transaction-data-model

/**
 * {
 *    "EventType":1,
 *    "Event":"TransactionsStatusChanged",
 *    "DateTime":"04032021211615",
 *    "CountryIsoCode":"KWT",
 *    "Data":{
 *       "InvoiceId":586170,
 *       "InvoiceReference":"2021000184",
 *       "CreatedDate":"04032021211555",
 *       "CustomerReference":"12223444",
 *       "CustomerName":"Test Webhook",
 *       "CustomerMobile":"96512345678",
 *       "CustomerEmail":"test@mail.com",
 *       "TransactionStatus":"SUCCESS",
 *       "PaymentMethod":"KNET",
 *       "UserDefinedField":null,
 *       "ReferenceId":"106310001097",
 *       "TrackId":"04-03-2021_477336",
 *       "PaymentId":"100202106359084366",
 *       "AuthorizationId":"B68413",
 *       "InvoiceValueInBaseCurrency":"456.75",
 *       "BaseCurrency":"KWD",
 *       "InvoiceValueInDisplayCurreny":"456.75",
 *       "DisplayCurrency":"KWD",
 *       "InvoiceValueInPayCurrency":"456.75",
 *       "PayCurrency":"KWD"
 *    }
 * }
 */
//https://docs.myfatoorah.com/docs/webhook-v1   版本v1
type MyFatoorahDepositBackReq struct {
	EventType      int                           `json:"EventType"` // 枚举：1 For Transaction Status Changed, 2 For Refund Status Changed, 3 For Balance Transferred, 4 For Supplier Status Changed, 5 For Recurring Status Changed
	Event          string                        `json:"Event"`     // 枚举：TransactionsStatusChanged,RefundStatusChanged,BalanceTransferred,SupplierStatusChanged
	DateTime       string                        `json:"DateTime"`  // ddMMyyyyHHmmss
	CountryIsoCode string                        `json:"CountryIsoCode"`
	Data           *MyFatoorahDepositBackReqData `json:"Data"`
	// Signature from header MyFatoorah-Signature
	//Signature string
}

type MyFatoorahDepositBackReqData struct {
	InvoiceId                     int     `json:"InvoiceId" mapstructure:"InvoiceId"` //psp订单号(发票号)
	InvoiceReference              string  `json:"InvoiceReference" mapstructure:"InvoiceReference"`
	CreatedDate                   string  `json:"CreatedDate" mapstructure:"CreatedDate"`             //ddMMyyyyHHmmss
	CustomerReference             string  `json:"CustomerReference" mapstructure:"CustomerReference"` //merchant订单号
	CustomerName                  string  `json:"CustomerName" mapstructure:"CustomerName"`
	CustomerMobile                string  `json:"CustomerMobile" mapstructure:"CustomerMobile"`
	CustomerEmail                 string  `json:"CustomerEmail" mapstructure:"CustomerEmail"`
	TransactionStatus             string  `json:"TransactionStatus" mapstructure:"TransactionStatus"` //枚举： SUCCESS, FAILED,CANCELED,AUTHORIZE
	PaymentMethod                 string  `json:"PaymentMethod" mapstructure:"PaymentMethod"`
	UserDefinedField              *string `json:"UserDefinedField" mapstructure:"UserDefinedField"`
	ReferenceId                   string  `json:"ReferenceId" mapstructure:"ReferenceId"`
	TrackId                       string  `json:"TrackId" mapstructure:"TrackId"`
	PaymentId                     string  `json:"PaymentId" mapstructure:"PaymentId"`
	AuthorizationId               string  `json:"AuthorizationId" mapstructure:"AuthorizationId"`
	InvoiceValueInBaseCurrency    string  `json:"InvoiceValueInBaseCurrency" mapstructure:"InvoiceValueInBaseCurrency"`
	BaseCurrency                  string  `json:"BaseCurrency" mapstructure:"BaseCurrency"`
	InvoiceValueInDisplayCurrency string  `json:"InvoiceValueInDisplayCurreny" mapstructure:"InvoiceValueInDisplayCurreny"` // Note: Typo in JSON field name preserved
	DisplayCurrency               string  `json:"DisplayCurrency" mapstructure:"DisplayCurrency"`
	InvoiceValueInPayCurrency     string  `json:"InvoiceValueInPayCurrency" mapstructure:"InvoiceValueInPayCurrency"`
	PayCurrency                   string  `json:"PayCurrency" mapstructure:"PayCurrency"` //币种
}

//------------

type MyFatoorahDepositBackRsp struct {
	IsSuccess        bool                          `json:"IsSuccess"`
	Message          string                        `json:"Message"`
	ValidationErrors interface{}                   `json:"ValidationErrors"` // Using interface{} to handle any type
	Data             *MyFatoorahDepositBackRspData `json:"Data"`
}

// TODO 感觉完全用不到
type MyFatoorahDepositBackRspData struct {
	InvoiceId           int           `json:"InvoiceId"`
	InvoiceStatus       string        `json:"InvoiceStatus"`
	InvoiceReference    string        `json:"InvoiceReference"`
	CustomerReference   string        `json:"CustomerReference"` // Pointer for nullable field
	CreatedDate         time.Time     `json:"CreatedDate"`       // Using time.Time for datetime
	ExpiryDate          string        `json:"ExpiryDate"`        // Keeping as string for custom format
	ExpiryTime          string        `json:"ExpiryTime"`        // Keeping as string for custom format
	InvoiceValue        float64       `json:"InvoiceValue"`
	Comments            *string       `json:"Comments"` // Pointer for nullable field
	CustomerName        string        `json:"CustomerName"`
	CustomerMobile      string        `json:"CustomerMobile"`
	CustomerEmail       *string       `json:"CustomerEmail"`    // Pointer for nullable field
	UserDefinedField    *string       `json:"UserDefinedField"` // Pointer for nullable field
	InvoiceDisplayValue string        `json:"InvoiceDisplayValue"`
	DueDeposit          float64       `json:"DueDeposit"`
	DepositStatus       string        `json:"DepositStatus"`
	InvoiceItems        []interface{} `json:"InvoiceItems"`
	InvoiceTransactions []interface{} `json:"InvoiceTransactions"` // Using interface{} for unknown structure
	Suppliers           []interface{} `json:"Suppliers"`           // Using interface{} for unknown structure
}
