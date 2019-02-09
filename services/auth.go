package services

import (
	"freelancers/app"
	"freelancers/models"
	"freelancers/util"
)

type authDAO interface {
	Login(rs app.RequestScope, user *models.User) *models.User
	Register(rs app.RequestScope, user *models.User) error
}

type AuthService struct {
	dao authDAO
}

func NewAuthService(dao authDAO) *AuthService {
	return &AuthService{dao}
}

func (s *AuthService) Login(rs app.RequestScope, user *models.User) {
	s.dao.Login(rs, user)
}

func (s *AuthService) Register(rs app.RequestScope, user *models.User) error {
	user.Password = util.HashAndSalt([]byte(user.Password))

	return s.dao.Register(rs, user)
}
