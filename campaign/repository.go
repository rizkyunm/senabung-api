package campaign

import "gorm.io/gorm"

type Repository interface {
	FindAll() ([]Campaign, error)
	FindByUserID(userID uint) ([]Campaign, error)
	FindByID(ID uint) (Campaign, error)
	FindBySlug(slug string) (Campaign, error)
	Save(campaign Campaign) (Campaign, error)
	Update(campaign Campaign) (Campaign, error)
	GetHighlight() ([]Campaign, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]Campaign, error) {
	var campaigns []Campaign

	if err := r.db.Find(&campaigns).Error; err != nil {
		return campaigns, err
	}

	return campaigns, nil
}

func (r *repository) GetHighlight() ([]Campaign, error) {
	var campaigns []Campaign

	if err := r.db.Find(&campaigns).Limit(6).Order("created_at desc").Error; err != nil {
		return campaigns, err
	}

	return campaigns, nil
}

func (r *repository) FindByUserID(userID uint) ([]Campaign, error) {
	var campaigns []Campaign

	if err := r.db.Where("user_id = ?", userID).Find(&campaigns).Error; err != nil {
		return campaigns, err
	}

	return campaigns, nil
}

func (r *repository) FindByID(ID uint) (Campaign, error) {
	var campaign Campaign

	if err := r.db.Where("id = ?", ID).Preload("User").Find(&campaign).Error; err != nil {
		return campaign, err
	}

	return campaign, nil
}

func (r *repository) FindBySlug(slug string) (Campaign, error) {
	var campaign Campaign

	if err := r.db.Where("slug = ?", slug).Preload("User").Find(&campaign).Error; err != nil {
		return campaign, err
	}

	return campaign, nil
}

func (r *repository) Save(campaign Campaign) (Campaign, error) {
	if err := r.db.Create(&campaign).Error; err != nil {
		return campaign, err
	}

	return campaign, nil
}

func (r *repository) Update(campaign Campaign) (Campaign, error) {
	if err := r.db.Save(&campaign).Error; err != nil {
		return campaign, err
	}

	return campaign, nil
}
