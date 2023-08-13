package campaign

import (
	"fmt"

	"github.com/gosimple/slug"
)

type Service interface {
	GetCampaigns(userID uint) ([]Campaign, error)
	GetCampaign(input GetCampaignDetailInput) (Campaign, error)
	GetCampaignBySlug(input GetCampaignDetailBySlug) (Campaign, error)
	CreateCampaign(input CreateCampaignInput) (Campaign, error)
	UpdateCampaign(inputID GetCampaignDetailInput, inputData CreateCampaignInput) (Campaign, error)
	UpdateCampaignImage(inputID GetCampaignDetailInput, imageURL string) (Campaign, error)
	GetHighlightCampaigns() ([]Campaign, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetCampaigns(userID uint) ([]Campaign, error) {
	if userID == 0 {
		campaigns, err := s.repository.FindAll()
		if err != nil {
			return campaigns, err
		}

		return campaigns, nil
	}

	campaigns, err := s.repository.FindByUserID(userID)
	if err != nil {
		return campaigns, err
	}

	return campaigns, nil
}

func (s *service) GetHighlightCampaigns() ([]Campaign, error) {
	campaigns, err := s.repository.GetHighlight()
	if err != nil {
		return campaigns, err
	}

	return campaigns, nil
}

func (s *service) GetCampaign(input GetCampaignDetailInput) (Campaign, error) {
	campaign, err := s.repository.FindByID(input.ID)

	if err != nil {
		return campaign, err
	}

	return campaign, nil
}

func (s *service) GetCampaignBySlug(input GetCampaignDetailBySlug) (Campaign, error) {
	campaign, err := s.repository.FindBySlug(input.Slug)

	if err != nil {
		return campaign, err
	}

	return campaign, nil
}

func (s *service) CreateCampaign(input CreateCampaignInput) (Campaign, error) {
	campaign := Campaign{
		Name:             input.Name,
		ShortDescription: input.ShortDescription,
		Description:      input.Description,
		GoalAmount:       input.GoalAmount,
		UserID:           input.User.ID,
		CampaignImage:    input.CampaignImage,
		Status:           Progress,
	}

	slugCandidate := fmt.Sprintf("%s %d", input.Name, input.User.ID)
	campaign.Slug = slug.Make(slugCandidate)

	newCampaign, err := s.repository.Save(campaign)
	if err != nil {
		return newCampaign, err
	}

	return newCampaign, nil
}

func (s *service) UpdateCampaign(inputID GetCampaignDetailInput, inputData CreateCampaignInput) (Campaign, error) {
	campaign, err := s.repository.FindByID(inputID.ID)
	if err != nil {
		return campaign, err
	}

	campaign.Name = inputData.Name
	campaign.ShortDescription = inputData.ShortDescription
	campaign.Description = inputData.Description
	campaign.CampaignImage = inputData.CampaignImage
	campaign.GoalAmount = inputData.GoalAmount

	updateCampaign, err := s.repository.Update(campaign)
	if err != nil {
		return updateCampaign, err
	}

	return updateCampaign, nil
}

func (s *service) UpdateCampaignImage(inputID GetCampaignDetailInput, imageURl string) (Campaign, error) {
	campaign, err := s.repository.FindByID(inputID.ID)
	if err != nil {
		return campaign, err
	}

	campaign.CampaignImage = imageURl

	updateCampaign, err := s.repository.Update(campaign)
	if err != nil {
		return updateCampaign, err
	}

	return updateCampaign, nil
}
