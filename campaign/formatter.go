package campaign

import (
	"fmt"
	"strings"
)

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

type CampaignDetailFormatter struct {
	ID               int                      `json:"id"`
	Name             string                   `json:"name"`
	ShortDescription string                   `json:"short_description"`
	Description      string                   `json:"description"`
	ImageUrl         string                   `json:"image_url"`
	GoalAmount       int                      `json:"goal_amount"`
	CurrentAmount    int                      `json:"current_amount"`
	UserId           int                      `json:"user_id"`
	Slug             string                   `json:"slug"`
	Perks            []string                 `json:"perks"`
	User             CampaignUserFormatter    `json:"user"`
	Images           []CampaignImageFormatter `json:"images"`
}

type CampaignUserFormatter struct {
	Name     string `json:"name"`
	ImageUrl string `json:"image_url"`
}

type CampaignImageFormatter struct {
	ImageUrl  string `json:"image_url"`
	IsPrimary bool   `json:"is_primary"`
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

func FormatCampaignDetail(campaign Campaign) CampaignDetailFormatter {
	var images string
	if len(campaign.CampaignImages) > 0 {
		image := campaign.CampaignImages[0].FileName
		images = fmt.Sprintf("images/%s", image)
	}
	var perks []string
	for _, perk := range strings.Split(campaign.Perk, ",") {
		perks = append(perks, strings.TrimSpace(perk))
	}

	user := CampaignUserFormatter{
		Name:     campaign.User.Name,
		ImageUrl: campaign.User.Avatar,
	}
	imagesField := []CampaignImageFormatter{}

	for _, v := range campaign.CampaignImages {
		campaignImageFormatter := CampaignImageFormatter{}
		campaignImageFormatter.ImageUrl = v.FileName
		campaignImageFormatter.IsPrimary = false
		if v.IsPrimary == 1 {
			campaignImageFormatter.IsPrimary = true
		}
		imagesField = append(imagesField, campaignImageFormatter)
	}

	campaignFormatter := CampaignDetailFormatter{
		ID:               campaign.Id,
		Name:             campaign.Name,
		ShortDescription: campaign.ShortDescription,
		Description:      campaign.Description,
		ImageUrl:         images,
		GoalAmount:       campaign.GoalAmount,
		CurrentAmount:    campaign.CurrentAmount,
		UserId:           campaign.UserId,
		Slug:             campaign.Slug,
		Perks:            perks,
		User:             user,
		Images:           imagesField,
	}
	return campaignFormatter
}
