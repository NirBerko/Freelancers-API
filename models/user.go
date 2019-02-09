package models

type User struct {
	DBBase
	Name     string `gorm:"column:name" json:"name"`
	Phone    string `gorm:"column:phone" json:"phone"`
	Password string `gorm:"column:password" json:"-"`
}

func (u User) GetID() uint {
	return u.ID
}

func (u User) GetName() string {
	return u.Name
}

func (u User) GetPhone() string {
	return u.Phone
}
