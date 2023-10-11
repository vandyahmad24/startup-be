package transaction

import "gorm.io/gorm"

type repository struct {
	db *gorm.DB
}

type Repository interface {
	GetByCampaignId(campaginId int) ([]Transaction, error)
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{
		db: db,
	}
}

func (r *repository) GetByCampaignId(campaginId int) ([]Transaction, error) {
	var transaction []Transaction

	err := r.db.Preload("User").Model(&Transaction{}).Where("campaign_id = ?", campaginId).Order("id desc").Find(&transaction).Error
	if err != nil {
		return transaction, err
	}
	return transaction, nil
}
