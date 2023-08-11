package transaction

import "gorm.io/gorm"

type Repository interface {
	GetByCampaignID(campaignID uint) ([]Transaction, error)
	GetByUserID(userID uint) ([]Transaction, error)
	GetByID(ID uint) (Transaction, error)
	Save(transaction Transaction) (Transaction, error)
	Update(transaction Transaction) (Transaction, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) GetByCampaignID(campaignID uint) ([]Transaction, error) {
	var transactions []Transaction

	if err := r.db.Preload("User").Where("campaign_id = ?", campaignID).Order("id desc").Find(&transactions).Error; err != nil {
		return transactions, err
	}

	return transactions, nil
}

func (r *repository) GetByUserID(userID uint) ([]Transaction, error) {
	var transactions []Transaction

	if err := r.db.Preload("Campaign.CampaignImages", "campaign_images.is_primary = 1").Where("user_id = ?", userID).Order("id desc").Find(&transactions).Error; err != nil {
		return transactions, err
	}

	return transactions, nil
}

func (r *repository) Save(transaction Transaction) (Transaction, error) {
	if err := r.db.Create(&transaction).Error; err != nil {
		return transaction, err
	}

	return transaction, nil
}

func (r *repository) Update(transaction Transaction) (Transaction, error) {
	if err := r.db.Save(&transaction).Error; err != nil {
		return transaction, err
	}

	return transaction, nil
}

func (r *repository) GetByID(ID uint) (Transaction, error) {
	var transaction Transaction

	if err := r.db.Where("id = ?", ID).Find(&transaction).Error; err != nil {
		return transaction, err
	}

	return transaction, nil
}
