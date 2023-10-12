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
	TransactionEventID string `json:"transaction_event_id"`
	SourceDC           string `json:"source_dc"`
	OrderID            string `json:"order_id" binding:"required"`
	MerchantID         string `json:"merchant_id"`
	MerchantEmail      string `json:"merchant_email"`
	Endpoint           string `json:"endpoint"`
	ContentType        string `json:"content_type"`
	Body               struct {
		VANumbers []struct {
			VANumber string `json:"va_number"`
			Bank     string `json:"bank"`
		} `json:"va_numbers"`
		TransactionTime   string        `json:"transaction_time"`
		TransactionStatus string        `json:"transaction_status" binding:"required"`
		TransactionID     string        `json:"transaction_id"`
		StatusMessage     string        `json:"status_message"`
		StatusCode        string        `json:"status_code"`
		SignatureKey      string        `json:"signature_key"`
		PaymentType       string        `json:"payment_type" binding:"required"`
		PaymentAmounts    []interface{} `json:"payment_amounts"`
		OrderID           string        `json:"order_id"`
		MerchantID        string        `json:"merchant_id"`
		GrossAmount       string        `json:"gross_amount"`
		FraudStatus       string        `json:"fraud_status"`
		ExpiryTime        string        `json:"expiry_time"`
		Currency          string        `json:"currency"`
	} `json:"body"`
}
