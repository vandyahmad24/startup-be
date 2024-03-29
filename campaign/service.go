package campaign

import (
	"errors"
	"fmt"
	"github.com/gosimple/slug"
)

type Service interface {
	FindCampaigns(userId int) ([]Campaign, error)
	GetCampaignById(input GetCampaignDetailInput) (Campaign, error)
	CreateCampaign(input CreateCampaignInput) (Campaign, error)
	UpdateCampaign(input GetCampaignDetailInput, inputData CreateCampaignInput) (Campaign, error)
	CreateCampaingImage(input CreateCampaingImage, fileLocation string) (CampaignImage, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository: repository}
}

func (s *service) FindCampaigns(userId int) ([]Campaign, error) {
	if userId != 0 {
		campaign, err := s.repository.FindByUserId(userId)
		if err != nil {
			return campaign, err
		}
		return campaign, err
	}

	campaign, err := s.repository.FindAll()
	if err != nil {
		return campaign, err
	}
	fmt.Println(campaign)
	return campaign, err

}

func (s *service) GetCampaignById(input GetCampaignDetailInput) (Campaign, error) {
	campaign, err := s.repository.FindByID(input.Id)
	if err != nil {
		return campaign, err
	}
	return campaign, nil
}

func (s *service) CreateCampaign(input CreateCampaignInput) (Campaign, error) {

	slugCandidate := slug.Make(fmt.Sprintf("%s-%d", input.Name, input.User.ID))
	campaign := Campaign{
		Name:             input.Name,
		ShortDescription: input.ShortDescription,
		Description:      input.Description,
		Perk:             input.Perks,
		GoalAmount:       input.GoalAmount,
		UserId:           input.User.ID,
		Slug:             slugCandidate,
	}
	newCampaign, err := s.repository.Save(campaign)
	if err != nil {
		return Campaign{}, err
	}

	return newCampaign, nil
}

func (s *service) UpdateCampaign(input GetCampaignDetailInput, inputData CreateCampaignInput) (Campaign, error) {
	campaign, err := s.repository.FindByID(input.Id)
	if err != nil {
		return Campaign{}, err
	}

	if campaign.UserId != input.Id {
		return Campaign{}, errors.New("Not an owner of the campaign")
	}

	campaign.Name = inputData.Name
	campaign.Description = inputData.Description
	campaign.ShortDescription = inputData.ShortDescription
	campaign.Perk = inputData.Perks
	campaign.GoalAmount = inputData.GoalAmount

	campaign, err = s.repository.Update(campaign)
	if err != nil {
		return Campaign{}, err
	}
	return campaign, nil
}

func (s *service) CreateCampaingImage(input CreateCampaingImage, fileLocation string) (CampaignImage, error) {

	isPrimary := 0
	if input.IsPrimary {
		_, err := s.repository.MarkAllImageAsNonPrimary(input.CampaignId)
		if err != nil {
			return CampaignImage{}, nil
		}
		isPrimary = 1
	}

	campaignImage := CampaignImage{
		CampaignId: input.CampaignId,
		FileName:   fileLocation,
		IsPrimary:  isPrimary,
	}

	newImage, err := s.repository.SaveImage(campaignImage)
	if err != nil {
		return CampaignImage{}, err
	}
	return newImage, nil
}
