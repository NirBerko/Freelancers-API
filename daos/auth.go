package daos

import (
	"freelancers/app"
	"freelancers/models"
)

type AuthDAO struct{}

func NewAuthDAO() *AuthDAO {
	return &AuthDAO{}
}

func (dao *AuthDAO) Login(rs app.RequestScope, user *models.User) (findUser models.User) {
	rs.Db().Where("email = ?", user.GetEmail()).First(&findUser)

	return
}

func (dao *AuthDAO) Register(rs app.RequestScope, user *models.User) error {
	create := rs.Db().Create(&user)

	return create.Error
}
