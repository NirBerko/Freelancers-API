package app

import (
	"github.com/jinzhu/gorm"
)

type RequestScope interface {
	Db() *gorm.DB
	SetDB(db *gorm.DB)
	RequestID() string
	SetUserID(userId uint)
	UserID() uint
}

type requestScope struct {
	db        *gorm.DB
	requestID string
	userID    uint
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

func (rs *requestScope) SetDB(db *gorm.DB) {
	rs.db = db
}

func (rs *requestScope) Db() *gorm.DB {
	return rs.db
}
