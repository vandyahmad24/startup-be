package transaction

import "startup/users"

type GetTransactionCampaignInput struct {
	Id int `uri:"id" binding:"required"`
}

type CreateTransactionInput struct {
	Amount     int `json:"amount" binding:"required"`
	CampaignId int `json:"campaign_id" binding:"required"`
	User       users.User
}

type TransactionData struct {
	VANumbers         []VANumber `json:"va_numbers"`
	TransactionTime   string     `json:"transaction_time"`
	TransactionStatus string     `json:"transaction_status"`
	TransactionID     string     `json:"transaction_id"`
	StatusMessage     string     `json:"status_message"`
	StatusCode        string     `json:"status_code"`
	SignatureKey      string     `json:"signature_key"`
	SettlementTime    string     `json:"settlement_time"`
	PaymentType       string     `json:"payment_type"`
	PaymentAmounts    []int      `json:"payment_amounts"`
	OrderID           string     `json:"order_id"`
	MerchantID        string     `json:"merchant_id"`
	GrossAmount       string     `json:"gross_amount"`
	FraudStatus       string     `json:"fraud_status"`
	ExpiryTime        string     `json:"expiry_time"`
	Currency          string     `json:"currency"`
}

type VANumber struct {
	VANumber string `json:"va_number"`
	Bank     string `json:"bank"`
}
