package transaction

import "gorm.io/gorm"

type repository struct {
	db *gorm.DB
}

type Repository interface {
	GetByCampaignId(campaginId int) ([]Transaction, error)
	GetByUserId(userId int) ([]Transaction, error)
	Save(transaction Transaction) (Transaction, error)
	Update(transaction Transaction) (Transaction, error)
	GetByCode(code string) (Transaction, error)
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

func (r *repository) GetByUserId(userId int) ([]Transaction, error) {
	var transaction []Transaction

	err := r.db.Debug().
		Preload("Campaign.CampaignImages", "campaign_image.is_primary = 1").
		Model(&Transaction{}).
		Where("user_id = ?", userId).
		Order("id desc").
		Find(&transaction).Error
	if err != nil {
		return transaction, err
	}
	return transaction, nil
}
func (r *repository) Save(transaction Transaction) (Transaction, error) {
	err := r.db.Create(&transaction).Error
	if err != nil {
		return Transaction{}, err
	}
	return transaction, nil
}

func (r *repository) Update(transaction Transaction) (Transaction, error) {
	err := r.db.Save(&transaction).Error
	if err != nil {
		return Transaction{}, err
	}
	return transaction, nil
}

func (r *repository) GetByCode(code string) (Transaction, error) {
	var transaction Transaction

	err := r.db.Debug().Where("code = ?", code).Find(&transaction).Error
	if err != nil {
		return transaction, err
	}
	return transaction, nil
}
