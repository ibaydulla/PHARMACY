package models

type Pharmacy struct {
	ID            int    `json:"id"`
	Name          string `json:"name"`
	Address       string `json:"address"`
	Pharmacyhours string `json:"pharmacy_hours"`
}
