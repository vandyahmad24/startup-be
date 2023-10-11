package transaction

import (
	"startup/campaign"
	"startup/users"
	"time"
)

type Transaction struct {
	Id         int
	CampaignId int
	UserId     int
	Amount     int
	Status     string
	Code       string
	PaymentUrl string
	User       users.User
	Campaign   campaign.Campaign
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

// TableName overrides the table name used by User to `transaction`
func (Transaction) TableName() string {
	return "transactions"
}
