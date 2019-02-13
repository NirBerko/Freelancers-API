package daos

import (
	"freelancers/app"
	"freelancers/models"
)

type UserDAO struct{}

func NewUserDAO() *UserDAO {
	return &UserDAO{}
}

func (dao *UserDAO) GetUserByID(rs app.RequestScope, id uint64) (findUser models.User) {
	rs.Db().Where("id = ?", id).First(&findUser)

	return
}
