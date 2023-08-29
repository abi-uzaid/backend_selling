package models

type User struct {
	ID int `json:"id"`
	Fullname string `json:"fullname" gorm:"type: varchar(255)"`
	Email string `json:"email" gorm:"type: varchar(255)"`
	Password string `json:"password" gorm:"type: varchar(255)"`
}

type UserResponse struct {
	ID int `json:"id"`
	Fullname string `json:"fullname" gorm:"type: varchar(255)"`
}

func (UserResponse) TableName() string {
	return "users"
}