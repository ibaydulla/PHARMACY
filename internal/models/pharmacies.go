package models

type Pharmacy struct {
	ID            int    `json:"id"`
	Name          string `json:"name"`
	Address       string `json:"address"`
	Pharmacyhours string `json:"pharmacy_hours"`
}

type PharmacyErrorResponse struct {
	Message string `json:"message"`
	Code    string `json:"code"`
}
