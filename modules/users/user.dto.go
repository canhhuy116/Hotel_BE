package users

type UserCreate struct {
	Username string `json:"username" `
	Password string `json:"password" `
	Name     string `json:"name"`
}

func (UserCreate) TableName() string {
	return User{}.TableName()
}
