package transaction

type GetTransactionCampaignInput struct {
	Id int `uri:"id" binding:"required"`
}
