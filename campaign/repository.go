package campaign

import "gorm.io/gorm"

type Repository interface {
	FindAll() ([]Campaign, error)
	FindByUserId(Id int) ([]Campaign, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]Campaign, error) {
	var campaign []Campaign
	err := r.db.Preload("CampaignImages", "campaign_image.is_primary = 1").Find(&campaign).Error
	if err != nil {
		return campaign, err
	}
	return campaign, nil
}

func (r *repository) FindByUserId(Id int) ([]Campaign, error) {
	var campaign []Campaign
	err := r.db.Where("user_id = ?", Id).Preload("CampaignImages", "campaign_image.is_primary = 1").Find(&campaign).Error
	if err != nil {
		return nil, err
	}
	return campaign, nil
}
