
package models

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"-"`
	Role     string `json:"role"`
}

type ErrorResponse struct {
	Success      bool   `json:"success`
	ErrorMassage string `json:"error_msg"`
	ErrorCode    string `json:"error_code"`
}

type UserResponse struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Role  string `json:"role"`
}

type LoginRequest struct{
	Email string `json:"email"`
	Password string `json:"password"`
}