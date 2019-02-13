package services

import (
	"freelancers/app"
	"freelancers/models"
	"freelancers/models/UIModels"
	"freelancers/util"
)

type (
	authDAO interface {
		Login(rs app.RequestScope, user *models.User) models.User
		Register(rs app.RequestScope, user *models.User) error
	}

	AuthService struct {
		dao authDAO
	}
)

func NewAuthService(dao authDAO) *AuthService {
	return &AuthService{dao}
}

func (s *AuthService) Login(rs app.RequestScope, user *models.User) *UIModels.User {
	findUser := s.dao.Login(rs, user)

	passwordMath := util.ComparePasswords(findUser.Password, []byte(user.Password))

	if passwordMath {
		return &UIModels.User{
			ID:        findUser.GetID(),
			Email:     findUser.GetEmail(),
			FirstName: findUser.GetFirstName(),
			LastName:  findUser.GetLastName(),
		}
	} else {
		return nil
	}
}

func (s *AuthService) Register(rs app.RequestScope, user *models.User) error {
	user.Password = util.HashAndSalt([]byte(user.Password))

	return s.dao.Register(rs, user)
}
