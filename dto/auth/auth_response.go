package authdto

type LoginResponse struct {
	Email string `json:"email" gorm:"type:varchar(255)"`
	Token string `json:"token" gorm:"type:varchar(255)"`
}

type RegisterResponse struct {
	Fullname string `json:"fullname" gorm:"type:varchar(255)"`
	Email    string `json:"email" gorm:"type:varchar(255)"`
}

type CheckAuthResponse struct {
	ID       int    `json:"id"`
	Fullname string `json:"fullname" gorm:"type:varchar(255)"`
	Email    string `json:"email" gorm:"type:varchar(255)"`
}
