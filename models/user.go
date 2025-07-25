package models

import "time"

type User struct {
	Id			uint		`json:"id" gorm:"primaryKey"`
	Name		string		`json:"name"`
	Email		string		`json:"email" gorm:"unique;not null"`
	Password	string		`json:"password"`
	CreatedAt	time.Time	`json:"createdat"`
}