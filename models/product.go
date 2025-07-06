package models

type Products struct {
	Id			uint		`json:"id" gorm:"primaryKey"`
	Title		string		`json:"title"`
	Description	string		`json:"description"`
	Price		float64		`json:"price"`
	Image		string		`json:"image"`
}