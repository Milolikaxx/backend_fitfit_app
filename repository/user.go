package repository

import (
	"backend_fitfit_app/model"
	"log"

	"gorm.io/gorm"
)

type userRepo struct {
	db *gorm.DB
}

func NewUserRepository() userRepository {
	db, err := NewDatabaseConnection()
	if err != nil {
		return nil
	}
	return userRepo{db: db}
}

type userRepository interface {
	FindAll() ([]model.User, error)
	Register(model.User) int64
	FindByID(key int) (*model.User, error)
	FindByEmail(key string) (*model.User, error)
}

func (u userRepo) FindAll() ([]model.User, error) {
	users := []model.User{}
	result := u.db.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}
func (u userRepo) FindByID(id int) (*model.User, error) {
	user := model.User{}
	result := u.db.Where("uid = ?", id).Find(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (u userRepo) FindByEmail(email string) (*model.User, error) {
	user := model.User{}
	result := u.db.Where("email = ?", email).Find(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (u userRepo) Register(account model.User) int64 {
	result := db.Create(&account)
	if result.RowsAffected > 0 {
		log.Printf("Register complete\nAffected row : %v", result.RowsAffected)
	} else {
		log.Printf("Register failed %v", result.RowsAffected)
	}
	return result.RowsAffected
}
