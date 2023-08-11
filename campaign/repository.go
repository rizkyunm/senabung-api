package campaign

import "gorm.io/gorm"

type Repository interface {
	FindAll() ([]Campaign, error)
	FindByUserID(userID uint) ([]Campaign, error)
	FindByID(ID uint) (Campaign, error)
	Save(campaign Campaign) (Campaign, error)
	Update(campaign Campaign) (Campaign, error)
	CreateImage(campaignImage CampaignImage) (CampaignImage, error)
	MarkAllImagesAsNonPrimary(campaignID uint) (bool, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]Campaign, error) {
	var campaigns []Campaign

	if err := r.db.Preload("CampaignImages", "campaign_images.is_primary = 1").Find(&campaigns).Error; err != nil {
		return campaigns, err
	}

	return campaigns, nil
}

func (r *repository) FindByUserID(userID uint) ([]Campaign, error) {
	var campaigns []Campaign

	if err := r.db.Where("user_id = ?", userID).Preload("CampaignImages", "campaign_images.is_primary = 1").Find(&campaigns).Error; err != nil {
		return campaigns, err
	}

	return campaigns, nil
}

func (r *repository) FindByID(ID uint) (Campaign, error) {
	var campaign Campaign

	if err := r.db.Where("id = ?", ID).Preload("User").Preload("CampaignImages").Find(&campaign).Error; err != nil {
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

func (r *repository) CreateImage(campaignImage CampaignImage) (CampaignImage, error) {
	if err := r.db.Create(&campaignImage).Error; err != nil {
		return campaignImage, err
	}

	return campaignImage, nil
}

func (r *repository) MarkAllImagesAsNonPrimary(campaignID uint) (bool, error) {
	if err := r.db.Model(&CampaignImage{}).Where("campaign_id = ?", campaignID).Update("is_primary", false).Error; err != nil {
		return false, err
	}

	return true, nil
}
