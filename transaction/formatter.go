package transaction

import "time"

type CampaignTransactionFormatter struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	Amount    int       `json:"amount"`
	CreatedAt time.Time `json:"created_at"`
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
