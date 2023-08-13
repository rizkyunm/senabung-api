package user

import (
	"database/sql/driver"
	"encoding/json"
	"gorm.io/gorm"
)

type Role string

const (
	Admin  Role = "admin"
	Client Role = "client"
)

func (r *Role) Scan(src any) error {
	*r = Role(src.([]byte))
	return nil
}

func (r Role) Value() (driver.Value, error) {
	return string(r), nil
}

func (r *Role) UnmarshalJSON(bytes []byte) error {
	var str string
	if err := json.Unmarshal(bytes, &str); err != nil {
		return err
	}
	*r = Role(str)
	return nil
}

func (r Role) MarshalJSON() ([]byte, error) {
	return json.Marshal(string(r))
}

type User struct {
	gorm.Model
	Name           string `gorm:"type:varchar(100)"`
	PhoneNumber    string `gorm:"type:varchar(20)"`
	Email          string `gorm:"type:varchar(100)"`
	PasswordHash   string `gorm:"type:varchar(255)"`
	AvatarFileName string `gorm:"type:varchar(255)"`
	Role           Role   `gorm:"type:varchar(20)"`
}
