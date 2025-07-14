package models

import "time"

type User struct {
	Id			uint		`json:"id" gorm:"primaryKey"`
	Name		string		`json:"name"`
	Email		string		`json:"email" gorm:"unique;not null"`
	Gender 		string		`json:"gender"`
	Phone		string		`json:"phone"`
	Address		string		`json:"address"`
	Role		string		`json:"role"`
	Status		string		`json:"status"`
	Password	string		`json:"password"`
	CreatedAt	time.Time	`json:"createdat"`
}