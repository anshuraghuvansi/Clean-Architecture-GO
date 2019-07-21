package models

import "time"

//User :
type User struct {
	ID        int64     `json:"id" gorm:"column:id"`
	Name      string    `json:"name,omitempty" gorm:"column:name"`
	Email     string    `json:"email,omitempty" gorm:"column:email"`
	Password  string    `json:"password,omitempty" gorm:"column:password"`
	CreatedAt time.Time `json:"createdAt,omitempty" gorm:"column:created_at"`
	UpdatedAt time.Time `json:"updatedAt,omitempty" gorm:"column:updated_at"`
}

//TableName :
func (User) TableName() string {
	return "users"
}
