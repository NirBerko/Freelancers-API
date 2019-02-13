package services

import (
	"freelancers/app"
	"freelancers/models"
	"freelancers/models/UIModels"
)

type (
	userDao interface {
		GetUserByID(rs app.RequestScope, id uint) models.User
	}

	UserService struct {
		dao userDao
	}
)

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
