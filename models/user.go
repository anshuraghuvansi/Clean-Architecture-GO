package models

//User :
type User struct {
	ID       int64  `json:"id" gorm:"column:id"`
	Name     string `json:"name,omitempty" gorm:"column:name"`
	Email    string `json:"email,omitempty" gorm:"column:email"`
	Password string `json:"password,omitempty" gorm:"column:password"`
}

//TableName :
func (User) TableName() string {
	return "users"
}
