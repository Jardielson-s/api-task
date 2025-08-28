package entity

type Login struct {
	Email    string `json:"email" example:"manager@company.com"`
	Password string `json:"password" example:"password123"`
}
