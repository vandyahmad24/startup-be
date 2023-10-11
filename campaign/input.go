package campaign

import "startup/users"

type GetCampaignDetailInput struct {
	Id int `uri:"id" binding:"required"`
}

type CreateCampaignInput struct {
	Name             string     `json:"name" binding:"required"`
	ShortDescription string     `json:"short_description" binding:"required"`
	Description      string     `json:"description" binding:"required"`
	GoalAmount       int        `json:"goal_amount" binding:"required"`
	Perks            string     `json:"perks" binding:"required"`
	User             users.User `json:"-"`
}

type CreateCampaingImage struct {
	CampaignId int  `form:"campaign_id" binding:"required"`
	IsPrimary  bool `form:"is_primary"`
}
