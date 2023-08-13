package campaign

import (
	"encoding/json"
	"github.com/rizkyunm/senabung-api/user"
)

type GetCampaignDetailInput struct {
	ID uint `uri:"id" binding:"required"`
}

type GetCampaignDetailBySlug struct {
	Slug string `uri:"slug" binding:"required"`
}

type CreateCampaignInput struct {
	Name             string      `json:"name" binding:"required"`
	ShortDescription string      `json:"short_description" binding:"required"`
	Description      string      `json:"description" binding:"required"`
	GoalAmount       json.Number `json:"goal_amount" binding:"required"`
	CampaignImage    string      `json:"campaign_image"`
	User             user.User
}
