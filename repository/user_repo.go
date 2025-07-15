package repository

import (
	"github.com/user-service/model"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user *model.User) error
	GetAll() ([]model.User, error)
	GetUsersByEmail(email string) (model.User, error)
	GetUsersByID(id int) (model.User, error)
	UpdateUser(user *model.User) error
}

type UserRepo struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepo {
	db.AutoMigrate(&model.User{})
	return &UserRepo{db}
}

func (r *UserRepo) Create(user *model.User) error {
	return r.db.Create(&user).Error
}

func (r *UserRepo) GetAll() ([]model.User, error) {
	var users []model.User
	err := r.db.Find(&users).Error

	if err != nil {
		return nil, err
	}

	return users, err
}

func (r *UserRepo) GetUsersByEmail(email string) (model.User, error) {
	var user model.User
	err := r.db.Where("email = ?", email).First(&user).Error

	if err != nil {
		return model.User{}, err
	}

	return user, err
}

func (r *UserRepo) GetUsersByID(id int) (model.User, error) {
	var user model.User
	err := r.db.Where("id = ?", id).First(&user).Error

	if err != nil {
		return model.User{}, err
	}

	return user, err
}

func (r *UserRepo) UpdateUser(user *model.User) error {

	err := r.db.Model(&user).Where("email = ?", user.Email).Updates(&user).Error

	if err != nil {
		return err
	}

	return err
}
