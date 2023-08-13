package transaction

import (
	"github.com/rizkyunm/senabung-api/campaign"
	"github.com/rizkyunm/senabung-api/user"
	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	CampaignID uint
	UserID     uint
	Amount     float64           `gorm:"type:decimal(10,2)"`
	Status     string            `gorm:"type:varchar(20)"`
	Code       string            `gorm:"type:varchar(50)"`
	PaymentURL string            `gorm:"type:varchar(255)"`
	User       user.User         `gorm:"foreignKey:UserID"`
	Campaign   campaign.Campaign `gorm:"foreignKey:CampaignID"`
	OrderNo    string            `gorm:"type:varchar(50);index"`
}
