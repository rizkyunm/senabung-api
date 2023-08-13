package transaction

import "gorm.io/gorm"

type Repository interface {
	GetByCampaignID(campaignID uint) ([]Transaction, error)
	GetByUserID(userID uint) ([]Transaction, error)
	GetByID(ID uint) (Transaction, error)
	GetByOrderID(orderNo string) (Transaction, error)
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

	if err := r.db.Preload("User").Where("campaign_id = ?", campaignID).Order("created_at desc").Limit(10).Find(&transactions).Error; err != nil {
		return transactions, err
	}

	return transactions, nil
}

func (r *repository) GetByUserID(userID uint) ([]Transaction, error) {
	var transactions []Transaction

	if err := r.db.Preload("Campaign").Where("user_id = ?", userID).Order("created_at desc").Limit(10).Find(&transactions).Error; err != nil {
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

	if err := r.db.Where("id = ?", ID).Preload("User").Preload("Campaign").Find(&transaction).Error; err != nil {
		return transaction, err
	}

	return transaction, nil
}

func (r *repository) GetByOrderID(orderNo string) (Transaction, error) {
	var transaction Transaction

	if err := r.db.Where("order_no = ?", orderNo).Preload("User").Preload("Campaign").Find(&transaction).Error; err != nil {
		return transaction, err
	}

	return transaction, nil
}
