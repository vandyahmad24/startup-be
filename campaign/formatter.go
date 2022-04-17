package campaign

import "fmt"

type CampaignFormatter struct {
	ID               int    `json:"id"`
	UserId           int    `json:"user_id"`
	Name             string `json:"name"`
	ShortDescription string `json:"short_description"`
	ImageUrl         string `json:"image_url"`
	GoalAmount       int    `json:"goal_amount"`
	CurrentAmount    int    `json:"current_amount"`
	Slug             string `json:"slug"`
}

func FormatCampaign(campaign Campaign) CampaignFormatter {
	var images string
	if len(campaign.CampaignImages) > 0 {
		image := campaign.CampaignImages[0].FileName
		images = fmt.Sprintf("images/%s", image)
	}

	campaignFormatter := CampaignFormatter{
		ID:               campaign.Id,
		UserId:           campaign.UserId,
		Name:             campaign.Name,
		ShortDescription: campaign.ShortDescription,
		GoalAmount:       campaign.GoalAmount,
		CurrentAmount:    campaign.CurrentAmount,
		ImageUrl:         images,
		Slug:             campaign.Slug,
	}

	return campaignFormatter
}
func FormatCampaigns(campaings []Campaign) []CampaignFormatter {

	if len(campaings) == 0 {
		return []CampaignFormatter{}
	}
	var campaignFormatters []CampaignFormatter
	for _, campaign := range campaings {
		campaignFormatter := FormatCampaign(campaign)
		campaignFormatters = append(campaignFormatters, campaignFormatter)
	}
	return campaignFormatters
}
