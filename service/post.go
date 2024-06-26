package service

import (
	"backend_fitfit_app/model"
	"backend_fitfit_app/repository"
)

type postServ struct{}

func NewPostService() PostService {
	return postServ{}
}

var postRepo = repository.NewPostRepository()

type PostService interface {
	GetAllPosts() ([]model.Post, error)
	GetPostByID(key int) (*model.Post, error)
	Save(model.Post) int64
	Update(wp model.Post, id int) int64
	Delete(id int) (int64, error)
}

func (postServ) GetAllPosts() ([]model.Post, error) {
	posts, err := postRepo.FindAll()
	if err != nil {
		return nil, err
	}
	return posts, nil
}

func (postServ) GetPostByID(id int) (*model.Post, error) {
	post, err := postRepo.FindByID(id)
	if err != nil {
		return nil, err
	}
	return post, nil
}

func (postServ) Save(post model.Post) int64 {
	rowsAff := postRepo.AddPost(post)
	if rowsAff > 0 {
		return 1
	} else if rowsAff == 0 {
		return 0
	} else {
		return -1
	}
}

func (p postServ) Update(post model.Post, id int) int64 {
	rowsAff := postRepo.UpdatePost(post, id)
	if rowsAff > 0 {
		return 1
	} else if rowsAff == 0 {
		return 0
	} else {
		return -1
	}
}

func (postServ) Delete(id int) (int64, error) {
	row, error := postRepo.DeletePost(id)
	if row > 0 {
		return 1, error
	} else if row == 0 {
		return 0, error
	} else {
		return -1, error
	}
}