package campaign

import "gorm.io/gorm"

type Repository interface {
	FindAll() ([]Campaign, error)
	FindByUserId(Id int) ([]Campaign, error)
	FindByID(Id int) (Campaign, error)
	Save(campaign Campaign) (Campaign, error)
	Update(campaign Campaign) (Campaign, error)
	SaveImage(image CampaignImage) (CampaignImage, error)
	MarkAllImageAsNonPrimary(campaignId int) (bool, error)
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

func (r *repository) FindByID(Id int) (Campaign, error) {
	var campaign Campaign
	err := r.db.Preload("CampaignImages").Preload("User").Where("id = ?", Id).Find(&campaign).Error
	if err != nil {
		return campaign, err
	}
	return campaign, nil
}

func (r *repository) Save(campaign Campaign) (Campaign, error) {
	err := r.db.Create(&campaign).Error
	if err != nil {
		return Campaign{}, err
	}
	return campaign, nil
}

func (r *repository) Update(campaign Campaign) (Campaign, error) {
	err := r.db.Save(&campaign).Error
	if err != nil {
		return Campaign{}, err
	}
	return campaign, nil
}

func (r *repository) SaveImage(image CampaignImage) (CampaignImage, error) {
	err := r.db.Create(&image).Error
	if err != nil {
		return CampaignImage{}, err
	}
	return image, nil
}

func (r *repository) MarkAllImageAsNonPrimary(campaignId int) (bool, error) {
	err := r.db.Model(&CampaignImage{}).Where("campaign_id = ?", campaignId).Update("is_primary", false).Error
	if err != nil {
		return false, err
	}
	return true, nil
}
