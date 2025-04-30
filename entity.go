package go_myfatoorah

// ----------generate invoice link-------------------------

type SendPaymentRequest struct {
	//must
	InvoiceValue       int    `json:"InvoiceValue"`
	CustomerName       string `json:"CustomerName"`
	NotificationOption string `json:"NotificationOption"`
	//option
	MobileCountryCode  string `json:"MobileCountryCode"`
	CustomerMobile     string `json:"CustomerMobile"`
	CustomerEmail      string `json:"CustomerEmail"`
	DisplayCurrencyIso string `json:"DisplayCurrencyIso"`
	CallBackUrl        string `json:"CallBackUrl"`
	ErrorUrl           string `json:"ErrorUrl"`
	Language           string `json:"Language"`
	CustomerReference  string `json:"CustomerReference"`
	CustomerAddress    struct {
		Block               string `json:"Block"`
		Street              string `json:"Street"`
		HouseBuildingNo     string `json:"HouseBuildingNo"`
		Address             string `json:"Address"`
		AddressInstructions string `json:"AddressInstructions"`
	} `json:"CustomerAddress"`
	InvoiceItems []struct {
		ItemName  string `json:"ItemName"`
		Quantity  int    `json:"Quantity"`
		UnitPrice int    `json:"UnitPrice"`
	} `json:"InvoiceItems"`
}

type SendPaymentResponse struct {
	IsSuccess        bool        `json:"IsSuccess"`
	Message          string      `json:"Message"`
	ValidationErrors interface{} `json:"ValidationErrors"`
	Data             struct {
		InvoiceId         int         `json:"InvoiceId"`
		InvoiceURL        string      `json:"InvoiceURL"` //inboice url
		CustomerReference string      `json:"CustomerReference"`
		UserDefinedField  interface{} `json:"UserDefinedField"`
	} `json:"Data"`
}

// ----------webhook-------------------------
// https://docs.myfatoorah.com/docs/transaction-data-model

type WebHookTransactionRequest struct {
	EventType      int    `json:"EventType"`
	Event          string `json:"Event"`
	DateTime       string `json:"DateTime"`
	CountryIsoCode string `json:"CountryIsoCode"`
	Data           struct {
		InvoiceId                    int         `json:"InvoiceId"`
		InvoiceReference             string      `json:"InvoiceReference"`
		CreatedDate                  string      `json:"CreatedDate"`
		CustomerReference            string      `json:"CustomerReference"`
		CustomerName                 string      `json:"CustomerName"`
		CustomerMobile               string      `json:"CustomerMobile"`
		CustomerEmail                string      `json:"CustomerEmail"`
		TransactionStatus            string      `json:"TransactionStatus"`
		PaymentMethod                string      `json:"PaymentMethod"`
		UserDefinedField             interface{} `json:"UserDefinedField"`
		ReferenceId                  string      `json:"ReferenceId"`
		TrackId                      string      `json:"TrackId"`
		PaymentId                    string      `json:"PaymentId"`
		AuthorizationId              string      `json:"AuthorizationId"`
		InvoiceValueInBaseCurrency   string      `json:"InvoiceValueInBaseCurrency"`
		BaseCurrency                 string      `json:"BaseCurrency"`
		InvoiceValueInDisplayCurreny string      `json:"InvoiceValueInDisplayCurreny"`
		DisplayCurrency              string      `json:"DisplayCurrency"`
		InvoiceValueInPayCurrency    string      `json:"InvoiceValueInPayCurrency"`
		PayCurrency                  string      `json:"PayCurrency"`
	} `json:"Data"`
}
