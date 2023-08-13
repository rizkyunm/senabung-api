package campaign

type CampaignFormatter struct {
	ID               uint           `json:"id"`
	Name             string         `json:"name"`
	ShortDescription string         `json:"short_description"`
	CampaignImage    string         `json:"campaign_image"`
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
		CampaignImage:    campaign.CampaignImage,
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
	ID               uint                  `json:"id"`
	Name             string                `json:"name"`
	ShortDescription string                `json:"short_description"`
	Description      string                `json:"description"`
	GoalAmount       float64               `json:"goal_amount"`
	CurrentAmount    float64               `json:"current_amount"`
	BackerCount      int                   `json:"backer_count"`
	UserID           uint                  `json:"user_id"`
	Slug             string                `json:"slug"`
	Status           CampaignStatus        `json:"status"`
	User             CampaignUserFormatter `json:"user"`
	CampaignImage    string                `json:"campaign_image"`
}

type CampaignUserFormatter struct {
	Name     string `json:"name"`
	ImageURL string `json:"image_url"`
}

func FormatCampaignDetail(campaign Campaign) CampaignDetailFormatter {
	campaignDetailFormatter := CampaignDetailFormatter{
		ID:               campaign.ID,
		Name:             campaign.Name,
		ShortDescription: campaign.ShortDescription,
		Description:      campaign.Description,
		CampaignImage:    campaign.CampaignImage,
		GoalAmount:       campaign.GoalAmount,
		CurrentAmount:    campaign.CurrentAmount,
		BackerCount:      campaign.BackerCount,
		UserID:           campaign.UserID,
		Slug:             campaign.Slug,
		Status:           campaign.Status,
	}

	user := campaign.User

	campaignUserFormatter := CampaignUserFormatter{
		Name:     user.Name,
		ImageURL: user.AvatarFileName,
	}

	campaignDetailFormatter.User = campaignUserFormatter

	return campaignDetailFormatter
}
