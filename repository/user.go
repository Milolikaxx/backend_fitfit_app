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
	FindByName(key string) (*model.User, error)
	UpdateUser(model.User, int) int64
}

func (u userRepo) FindAll() ([]model.User, error) {
	users := []model.User{}
	result := u.db.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}

func (userRepo) FindByID(id int) (*model.User, error) {
	user := model.User{}
	result := db.Where("uid = ?", id).Find(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (userRepo) FindByEmail(email string) (*model.User, error) {
	user := model.User{}
	result := db.Where("email = ?", email).Find(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (userRepo) FindByName(name string) (*model.User, error) {
	user := model.User{}
	result := db.Where("name= ?", name).Find(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (userRepo) Register(user model.User) int64 {
	result := db.Create(&user)
	if result.RowsAffected > 0 {
		log.Printf("Register complete\nAffected row : %v", result.RowsAffected)
	} else {
		log.Printf("Register failed %v", result.RowsAffected)
	}
	return result.RowsAffected
}

func (userRepo) UpdateUser(user model.User, id int) int64 {
	result := db.Model(&model.User{}).Where("uid = ?", id).Updates(&user)
	if result.RowsAffected > 0 {
		log.Printf("Update User complete\nAffected row : %v", result.RowsAffected)
	} else {
		log.Printf("Update User failed\nAffected row : %v", result.RowsAffected)
	}
	return result.RowsAffected
}
