package app

import (
	"github.com/jinzhu/gorm"
)

type RequestScope interface {
	Db() *gorm.DB
	SetDB(db *gorm.DB)
	RequestID() string
	SetUserID(userId uint64)
	UserID() uint64
}

type requestScope struct {
	db        *gorm.DB
	requestID string
	userID    uint64
}

func (rs *requestScope) RequestID() string {
	return rs.requestID
}

func (rs *requestScope) SetUserID(userId uint64) {
	rs.userID = userId
}

func (rs *requestScope) UserID() uint64 {
	return rs.userID
}

func (rs *requestScope) SetDB(db *gorm.DB) {
	rs.db = db
}

func (rs *requestScope) Db() *gorm.DB {
	return rs.db
}
