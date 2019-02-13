package services

import (
	"freelancers/app"
	"freelancers/models"
	"freelancers/models/UIModels"
	"freelancers/util"
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

func (s *UserService) GetUserDetails(rs app.RequestScope) util.ResultParser {
	findUser := s.dao.GetUserByID(rs, rs.UserID())

	user := UIModels.User{
		ID:        findUser.GetID(),
		Email:     findUser.GetEmail(),
		FirstName: findUser.GetFirstName(),
		LastName:  findUser.GetLastName(),
	}

	return util.ResultParser{
		Data:   user,
		IsDone: true,
		Error:  nil,
	}
}
