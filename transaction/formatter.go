package transaction

import "time"

type CampaignTransactionFormatter struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	Amount    int       `json:"amount"`
	CreatedAt time.Time `json:"created_at"`
}

type UserTransactionFormatter struct {
	Id        int               `json:"id"`
	Amount    int               `json:"amount"`
	Status    string            `json:"status"`
	CreatedAt time.Time         `json:"created_at"`
	Campaign  CampaignFormatter `json:"campaign"`
}
type CampaignFormatter struct {
	Name     string `json:"string"`
	ImageUrl string `json:"imageUrl"`
}

type TransactionFormatter struct {
	Id         int       `json:"id"`
	CampaignId int       `json:"campaign_id"`
	Amount     int       `json:"amount"`
	UserId     int       `json:"user_id"`
	Status     string    `json:"status"`
	PaymentUrl string    `json:"payment_url"`
	Code       string    `json:"code"`
	CreatedAt  time.Time `json:"created_at"`
}

func FormatCampaignTransaction(transaction Transaction) CampaignTransactionFormatter {
	formatter := CampaignTransactionFormatter{
		Id:        transaction.Id,
		Name:      transaction.User.Name,
		Amount:    transaction.Amount,
		CreatedAt: transaction.CreatedAt,
	}
	return formatter
}
func FormatCampaignTransactions(transaction []Transaction) []CampaignTransactionFormatter {
	var response []CampaignTransactionFormatter

	if len(transaction) == 0 {
		return response
	}

	for _, v := range transaction {
		res := FormatCampaignTransaction(v)
		response = append(response, res)
	}
	return response

}

func FormatUserTransactionFormatter(transaction Transaction) UserTransactionFormatter {
	var imageUrl string
	if len(transaction.Campaign.CampaignImages) > 0 {
		imageUrl = transaction.Campaign.CampaignImages[0].FileName

	}

	formatter := UserTransactionFormatter{
		Id:        transaction.Id,
		Amount:    transaction.Amount,
		Status:    transaction.Status,
		CreatedAt: transaction.CreatedAt,
		Campaign: CampaignFormatter{
			Name:     transaction.Campaign.Name,
			ImageUrl: imageUrl,
		},
	}
	return formatter
}

func FormatUserTransactionFormatters(transaction []Transaction) []UserTransactionFormatter {
	var response []UserTransactionFormatter

	if len(transaction) == 0 {
		return response
	}

	for _, v := range transaction {
		res := FormatUserTransactionFormatter(v)
		response = append(response, res)
	}
	return response

}

func FormatTransactionFormatter(transaction Transaction) TransactionFormatter {
	return TransactionFormatter{
		Id:         transaction.Id,
		CampaignId: transaction.CampaignId,
		Amount:     transaction.Amount,
		UserId:     transaction.UserId,
		Status:     transaction.Status,
		PaymentUrl: transaction.PaymentUrl,
		Code:       transaction.Code,
		CreatedAt:  transaction.CreatedAt,
	}
}
