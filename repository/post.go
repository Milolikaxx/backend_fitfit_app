package repository

import (
	"backend_fitfit_app/model"
	"log"

	"gorm.io/gorm"
)

type postRepo struct {
	db *gorm.DB
}

func NewPostRepository() postRepository {
	db, err := NewDatabaseConnection()
	if err != nil {
		return nil
	}
	return postRepo{db: db}
}

type postRepository interface {
	FindAll() ([]model.Post, error)
	FindByID(key int) (*model.Post, error)
	AddPost(model.Post) int64
	UpdatePost(wp model.Post, id int) int64
}

func (postRepo) FindAll() ([]model.Post, error) {
	posts := []model.Post{}
	result := db.Joins("User").Omit("post.created_at", "post.updated_at", "birthday", "email", "password", "google_id", "created_at", "updated_at").Find(&posts)
	// result := db.Joins("Playlist").Find(&posts)
	if result.Error != nil {
		return nil, result.Error
	}
	return posts, nil
}

func (postRepo) FindByID(id int) (*model.Post, error) {
	post := model.Post{}
	result := db.Where("wpid = ?", id).Find(&post)
	if result.Error != nil {
		return nil, result.Error
	}
	return &post, nil
}

func (postRepo) AddPost(post model.Post) int64 {
	result := db.Create(&post)
	if result.RowsAffected > 0 {
		log.Printf("Add Post complete\nAffected row : %v", result.RowsAffected)
	} else {
		log.Printf("Add Post failed %v", result.RowsAffected)
	}
	return result.RowsAffected
}

func (postRepo) UpdatePost(post model.Post, id int) int64 {
	result := db.Model(&model.WorkoutProfile{}).Where("wpid = ?", id).Updates(&post)
	if result.RowsAffected > 0 {
		log.Printf("Update Post complete\nAffected row : %v", result.RowsAffected)
	} else {
		log.Printf("Update Post failed\nAffected row : %v", result.RowsAffected)
	}
	return result.RowsAffected
}
