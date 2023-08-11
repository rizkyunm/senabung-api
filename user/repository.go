package user

import (
	"gorm.io/gorm"
)

type Repository interface {
	Save(user User) (User, error)
	FindByEmail(email string) (User, error)
	FindByID(ID uint) (User, error)
	Update(user User) (User, error)
	FindAll() ([]User, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Save(user User) (User, error) {
	if err := r.db.Create(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

func (r *repository) FindByEmail(email string) (User, error) {
	var user User

	if err := r.db.Where("email = ?", email).Find(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}

func (r *repository) FindByID(ID uint) (User, error) {
	var user User

	if err := r.db.Where("id = ?", ID).Find(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}

func (r *repository) Update(user User) (User, error) {
	if err := r.db.Save(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}

func (r *repository) FindAll() ([]User, error) {
	var users []User

	if err := r.db.Find(&users).Error; err != nil {
		return users, err
	}

	return users, nil
}
