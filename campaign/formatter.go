package campaign

import "strings"

type CampaignFormatter struct {
	ID               uint           `json:"id"`
	Name             string         `json:"name"`
	ShortDescription string         `json:"short_description"`
	ImageURL         string         `json:"image_url"`
	GoalAmount       float64        `json:"goal_amount"`
	CurrentAmount    float64        `json:"current_amount"`
	Slug             string         `json:"slug"`
	Status           CampaignStatus `json:"status"`
}

func FormatCampaign(campaign Campaign) CampaignFormatter {
	campaignFormatter := CampaignFormatter{
		ID:               campaign.ID,
		Name:             campaign.Name,
		ShortDescription: campaign.ShortDescription,
		GoalAmount:       campaign.GoalAmount,
		CurrentAmount:    campaign.CurrentAmount,
		Slug:             campaign.Slug,
		Status:           campaign.Status,
		ImageURL:         "",
	}

	if len(campaign.CampaignImages) > 0 {
		campaignFormatter.ImageURL = campaign.CampaignImages[0].FileName
	}

	return campaignFormatter
}

func FormatCampaigns(campaigns []Campaign) []CampaignFormatter {
	campaignsFormatter := []CampaignFormatter{}

	for _, campaign := range campaigns {
		campaignFormatter := FormatCampaign((campaign))
		campaignsFormatter = append(campaignsFormatter, campaignFormatter)
	}

	return campaignsFormatter
}

type CampaignDetailFormatter struct {
	ID               uint                      `json:"id"`
	Name             string                    `json:"name"`
	ShortDescription string                    `json:"short_description"`
	Description      string                    `json:"description"`
	ImageURL         string                    `json:"image_url"`
	GoalAmount       float64                   `json:"goal_amount"`
	CurrentAmount    float64                   `json:"current_amount"`
	BackerCount      int                       `json:"backer_count"`
	UserID           uint                      `json:"user_id"`
	Slug             string                    `json:"slug"`
	Status           CampaignStatus            `json:"status"`
	Perks            []string                  `json:"perks"`
	User             CampaignUserFormatter     `json:"user"`
	Images           []CampaignImagesFormatter `json:"images"`
}

type CampaignUserFormatter struct {
	Name     string `json:"name"`
	ImageURL string `json:"image_url"`
}

type CampaignImagesFormatter struct {
	ImageURL  string `json:"image_url"`
	IsPrimary bool   `json:"is_primary"`
}

func FormatCampaignDetail(campaign Campaign) CampaignDetailFormatter {
	campaignDetailFormatter := CampaignDetailFormatter{
		ID:               campaign.ID,
		Name:             campaign.Name,
		ShortDescription: campaign.ShortDescription,
		Description:      campaign.Description,
		ImageURL:         "",
		GoalAmount:       campaign.GoalAmount,
		CurrentAmount:    campaign.CurrentAmount,
		BackerCount:      campaign.BackerCount,
		UserID:           campaign.UserID,
		Slug:             campaign.Slug,
		Status:           campaign.Status,
	}

	if len(campaign.CampaignImages) > 0 {
		campaignDetailFormatter.ImageURL = campaign.CampaignImages[0].FileName
	}

	var perks []string

	for _, perk := range strings.Split(campaign.Perks, ",") {
		perks = append(perks, strings.TrimSpace(perk))
	}

	campaignDetailFormatter.Perks = perks

	user := campaign.User

	campaignUserFormatter := CampaignUserFormatter{
		Name:     user.Name,
		ImageURL: user.AvatarFileName,
	}

	campaignDetailFormatter.User = campaignUserFormatter

	images := []CampaignImagesFormatter{}

	for _, image := range campaign.CampaignImages {
		campaignImagesFormatter := CampaignImagesFormatter{
			ImageURL:  image.FileName,
			IsPrimary: false,
		}

		if image.IsPrimary == 1 {
			campaignImagesFormatter.IsPrimary = true
		}

		images = append(images, campaignImagesFormatter)
	}

	campaignDetailFormatter.Images = images

	return campaignDetailFormatter
}
