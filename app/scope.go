package app

import (
	"freelancers/models"
	"github.com/jinzhu/gorm"
)

type RequestScope interface {
	Db() *gorm.DB
	SetDB(db *gorm.DB)
	RequestID() string
	SetUserID(userId uint)
	UserID() uint
	SetUser(user models.User)
	User() models.User
}

type requestScope struct {
	db        *gorm.DB
	requestID string
	userID    uint
	user      models.User
}

func (rs *requestScope) RequestID() string {
	return rs.requestID
}

func (rs *requestScope) SetUserID(userId uint) {
	rs.userID = userId
}

func (rs *requestScope) UserID() uint {
	return rs.userID
}

func (rs *requestScope) SetUser(user models.User) {
	rs.user = user
}

func (rs *requestScope) User() models.User {
	return rs.user
}

func (rs *requestScope) SetDB(db *gorm.DB) {
	rs.db = db
}

func (rs *requestScope) Db() *gorm.DB {
	return rs.db
}
