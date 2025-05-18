package go_myfatoorah

import "time"

// ----------generate invoice link-------------------------

/**
 * {
 *     "CustomerName": "larry",
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
	InvoiceValue       float64 `json:"InvoiceValue"` // using float64 for decimal numbers
	CustomerName       string  `json:"CustomerName"`
	NotificationOption string  `json:"NotificationOption"`
	//option
	DisplayCurrencyIso string    `json:"DisplayCurrencyIso"`
	MobileCountryCode  string    `json:"MobileCountryCode,omitempty"` // marked as optional
	ExpiryDate         time.Time `json:"ExpiryDate"`
	CallBackUrl        string    `json:"CallBackUrl"`
	Language           string    `json:"Language"`
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
	InvoiceId         int     `json:"InvoiceId"`
	InvoiceURL        string  `json:"InvoiceURL"`
	CustomerReference *string `json:"CustomerReference"` // Using pointer to allow null
	UserDefinedField  *string `json:"UserDefinedField"`  // Using pointer to allow null
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
type MyFatoorahDepositBackReq struct {
	EventType      int                           `json:"EventType"`
	Event          string                        `json:"Event"`
	DateTime       string                        `json:"DateTime"`
	CountryIsoCode string                        `json:"CountryIsoCode"`
	Data           *MyFatoorahDepositBackReqData `json:"Data"`
	// Signature field is not in the JSON structure but exists in the Java class
	signature string // Unexported since setter doesn't expose it
}

type MyFatoorahDepositBackReqData struct {
	InvoiceId                     int     `json:"InvoiceId"`
	InvoiceReference              string  `json:"InvoiceReference"`
	CreatedDate                   string  `json:"CreatedDate"`
	CustomerReference             *string `json:"CustomerReference"`
	CustomerName                  string  `json:"CustomerName"`
	CustomerMobile                string  `json:"CustomerMobile"`
	CustomerEmail                 string  `json:"CustomerEmail"`
	TransactionStatus             string  `json:"TransactionStatus"`
	PaymentMethod                 string  `json:"PaymentMethod"`
	UserDefinedField              *string `json:"UserDefinedField"`
	ReferenceId                   string  `json:"ReferenceId"`
	TrackId                       string  `json:"TrackId"`
	PaymentId                     string  `json:"PaymentId"`
	AuthorizationId               string  `json:"AuthorizationId"`
	InvoiceValueInBaseCurrency    string  `json:"InvoiceValueInBaseCurrency"`
	BaseCurrency                  string  `json:"BaseCurrency"`
	InvoiceValueInDisplayCurrency string  `json:"InvoiceValueInDisplayCurreny"` // Note: Typo in JSON field name preserved
	DisplayCurrency               string  `json:"DisplayCurrency"`
	InvoiceValueInPayCurrency     string  `json:"InvoiceValueInPayCurrency"`
	PayCurrency                   string  `json:"PayCurrency"`
}

//------------

type MyFatoorahDepositBackRsp struct {
	IsSuccess        bool                          `json:"IsSuccess"`
	Message          string                        `json:"Message"`
	ValidationErrors interface{}                   `json:"ValidationErrors"` // Using interface{} to handle any type
	Data             *MyFatoorahDepositBackRspData `json:"Data"`
}

type MyFatoorahDepositBackRspData struct {
	InvoiceId           int           `json:"InvoiceId"`
	InvoiceStatus       string        `json:"InvoiceStatus"`
	InvoiceReference    string        `json:"InvoiceReference"`
	CustomerReference   *string       `json:"CustomerReference"` // Pointer for nullable field
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
