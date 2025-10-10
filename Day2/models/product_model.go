package models

type Product struct {
	ID           int     `json:"id"`
	Namee        string  `json:"namee"`
	Descriptionn string  `json:"descriptionn"`
	Price        float64 `json:"price"`
	Quantity     int     `json:"quantity"`
}
