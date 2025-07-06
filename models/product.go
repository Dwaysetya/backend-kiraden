package models

type Products struct {
	Id			uint		`json:"id" gorm:"primaryKey"`
	Title		string		`json:"title"`
	Description	string		`json:"description"`
	Price		string		`json:"price"`
	Image		string		`json:"image"`
}