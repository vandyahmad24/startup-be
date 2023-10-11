package campaign

import (
	"fmt"
	"github.com/gosimple/slug"
)

type Service interface {
	FindCampaigns(userId int) ([]Campaign, error)
	GetCampaignById(input GetCampaignDetailInput) (Campaign, error)
	CreateCampaign(input CreateCampaignInput) (Campaign, error)
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
