package UIModels

type User struct {
	ID        uint   `json:"id"`
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	UserType  int    `json:"user_type"`
}

func (u User) GetID() uint {
	return u.ID
}

func (u User) GetEmail() string {
	return u.Email
}

func (u User) GetFirstName() string {
	return u.FirstName
}

func (u User) GetLastName() string {
	return u.LastName
}
