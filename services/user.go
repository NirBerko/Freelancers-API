package services

import (
	"freelancers/app"
	"freelancers/models"
	"freelancers/models/UIModels"
	"github.com/jinzhu/gorm"
)

type (
	userDao interface {
		GetUserByID(rs app.RequestScope, id uint64) models.User
	}

	UserService struct {
		dao userDao
	}
)

type aa struct {
	db        *gorm.DB
	requestID string
	userID    uint64
}

func NewUserService(dao userDao) *UserService {
	return &UserService{dao}
}

func (s *UserService) GetUserDetails(rs app.RequestScope) UIModels.User {
	findUser := s.dao.GetUserByID(rs, rs.UserID())

	return UIModels.User{
		ID:        findUser.GetID(),
		Email:     findUser.GetEmail(),
		FirstName: findUser.GetFirstName(),
		LastName:  findUser.GetLastName(),
	}
}
