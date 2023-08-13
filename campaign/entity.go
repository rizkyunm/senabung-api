package campaign

import (
	"database/sql/driver"
	"github.com/rizkyunm/senabung-api/user"
	"gorm.io/gorm"
)

type CampaignStatus string

const (
	Pending  CampaignStatus = "Pending"
	Progress CampaignStatus = "Progress"
	Closed   CampaignStatus = "Closed"
)

func (r *CampaignStatus) Scan(src any) error {
	*r = CampaignStatus(src.([]byte))
	return nil
}

func (r CampaignStatus) Value() (driver.Value, error) {
	return string(r), nil
}

func (r *CampaignStatus) UnmarshalJSON(bytes []byte) error {
	*r = CampaignStatus(bytes)
	return nil
}

func (r CampaignStatus) MarshalJSON() ([]byte, error) {
	return []byte(`"` + string(r) + `"`), nil
}

type Campaign struct {
	gorm.Model
	UserID           uint
	Name             string         `gorm:"type:varchar(100)"`
	ShortDescription string         `gorm:"type:text"`
	Description      string         `gorm:"type:text"`
	BackerCount      int            `gorm:"type:int"`
	GoalAmount       float64        `gorm:"type:decimal(10,2)"`
	CurrentAmount    float64        `gorm:"type:decimal(10,2)"`
	Slug             string         `gorm:"type:varchar(255);index"`
	CampaignImage    string         `gorm:"type:varchar(255)"`
	Status           CampaignStatus `gorm:"type:varchar(20)"`
	User             user.User      `gorm:"foreignKey:UserID"`
}
