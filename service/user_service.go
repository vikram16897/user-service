package service

import (
	"github.com/user-service/model"
	"github.com/user-service/repository"
	"github.com/user-service/utils"
)

type UserService interface {
	CreateUser(user *model.User) error
	GetUsers() ([]model.User, error)
	GetUsersByEmail(email string) (model.User, error)
	GetUsersByID(id int) (model.User, error)
	UpdateUser(user *model.User) error
}

type userService struct {
	repo repository.UserRepository
}

func NewUserRepository(r repository.UserRepository) userService {
	return userService{r}
}

func (u userService) CreateUser(user *model.User) error {
	user.Password, _ = utils.Hash(user.Password)
	return u.repo.Create(user)
}

func (u userService) GetUsers() ([]model.User, error) {
	return u.repo.GetAll()
}

func (u userService) GetUsersByEmail(email string) (model.User, error) {
	return u.repo.GetUsersByEmail(email)
}

func (u userService) GetUsersByID(id int) (model.User, error) {
	return u.repo.GetUsersByID(id)
}

func (u userService) UpdateUser(user *model.User) error {
	return u.repo.UpdateUser(user)
}
