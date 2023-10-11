package campaign

import (
	"startup/users"
	"time"
)

type Campaign struct {
	Id               int             `json:"id"`
	UserId           int             `json:"user_id"`
	Name             string          `json:"name"`
	ShortDescription string          `json:"short_description"`
	Description      string          `json:"description"`
	Perk             string          `json:"perk"`
	BackerCount      int             `json:"backer_count"`
	GoalAmount       int             `json:"goal_amount"`
	CurrentAmount    int             `json:"current_amount"`
	Slug             string          `json:"slug"`
	CreatedAt        time.Time       `json:"created_at"`
	UpdatedAt        time.Time       `json:"updated_at"`
	CampaignImages   []CampaignImage `json:"campaign_images"`
	User             users.User      `json:"user"`
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
