package main

//APIData - struct of APIData
type APIData struct {
	StoreKey            int     `json:"storeKey"`
	HomeStoreKey        int     `json:"homeStoreKey"`
	CustKey             int     `json:"custKey"`
	CardHolderKey       int     `json:"cardholderKey"`
	InvoiceSN           string  `json:"invoiceSN"`
	SzCODInvoiceType    string  `json:"szCODInvoiceType"`
	SzCODDocumentNumber string  `json:"szCODDocumentNumber"`
	Date                string  `json:"date"`
	Time                string  `json:"time"`
	AmountToPay         float64 `json:"amountToPay"`
	AmountCollected     float64 `json:"amountCollected"`
	Rules               []Rule  `json:"rules"`
}

//Rule - struct of data in
type Rule struct {
	Rule  string  `json:"rule"`
	Qty   float64 `json:"qty"`
	Total float64 `json:"total"`
}
