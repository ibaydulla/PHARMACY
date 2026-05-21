package models

type Medicines struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       string `json:"price"`
	NewPrice    string `json:"new_price"`
	CategoryID  int    `json:"category_id"`
}

type MedicinesErrorResponse struct {
	Message string `json:"message"`
	Code    string `json:"code"`
}
