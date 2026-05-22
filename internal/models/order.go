package models

type Order struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Price    string `json:"price"`
	Description string `json:"description"`
}

type OrderErrorResponse struct {
	Message string `json:"message"`
	Code    string `json:"code"`
}
