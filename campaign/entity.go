package campaign

import (
	"startup/users"
	"time"
)

type Campaign struct {
	Id               int
	UserId           int
	Name             string
	ShortDescription string
	Description      string
	Perk             string
	BackerCount      int
	GoalAmount       int
	CurrentAmount    int
	Slug             string
	CreatedAt        time.Time
	UpdatedAt        time.Time
	CampaignImages   []CampaignImage
	User             users.User
}

// TableName overrides the table name used by User to `campaign`
func (Campaign) TableName() string {
	return "campaign"
}

type CampaignImage struct {
	Id         int
	CampaignId int
	FileName   string
	IsPrimary  int
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

func (CampaignImage) TableName() string {
	return "campaign_image"
}
