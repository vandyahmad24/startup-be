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

type TransactionNotificationInput struct {
	TransactionStatus string `json:"transaction_status"`
	Code              string `json:"code"`
	PaymentType       string `json:"payment_type"`
	FraudStatus       string `json:"fraud_status"`
}
