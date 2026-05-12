package models

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

type UserErrorResponse struct {
	Message string `json:"message"`
	Code    string `json:"code"`
}
