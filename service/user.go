package service

import (
	"backend_fitfit_app/model"
	"backend_fitfit_app/repository"
	"log"

	"golang.org/x/crypto/bcrypt"
)

type userServ struct{}

func NewUserService() UserService {
	return userServ{}
}

var userRepo = repository.NewUserRepository()

type UserService interface {
	GetAllUsers() ([]model.User, error)
	Register(model.User) int64
	GetUserByID(key int) (*model.User, error)
	GetUserByEmail(key string) (*model.User, error)
}

func (u userServ) GetAllUsers() ([]model.User, error) {
	users, err := userRepo.FindAll()
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (u userServ) GetUserByID(id int) (*model.User, error) {
	user, err := userRepo.FindByID(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u userServ) GetUserByEmail(email string) (*model.User, error) {
	user, err := userRepo.FindByEmail(email)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u userServ) Register(user model.User) int64 {
	usr, _ := userRepo.FindByEmail(user.Email)
	if usr.Uid == 0 {
		pwdHash := hashPassword(user.Password)
		user.Password = pwdHash
		rowsAff := userRepo.Register(user)
		if rowsAff > 0 {
			return 1
		} else if rowsAff == 0 {
			return 0
		} else {
			return -1
		}
	} else {
		return 0
	}
}

func hashPassword(pwd string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	if err != nil {
		log.Println(err)
	}
	return string(hash)
}
