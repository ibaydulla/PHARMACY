package models

type Category struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	
}

type CategoryErrorResponse struct {
	Message string `json:"message"`
	Code    string `json:"code"`
}