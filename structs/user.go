package structs

type UserResponse struct {
	Id			uint		`json:"id"`
	Name		string		`json:"name"`
	Email		string		`json:"email"`
	CreatedAt	string		`json:"createdat"`
	UpdatedAt	string		`json:"updatedat"`
	Token		*string		`json:"token,omitempty"`
}

type UserCreateRequest struct {
	Name		string		`json:"name" binding:"required" `
	Email		string		`json:"email" binding:"required" gorm:"unique;not null"`
	Password	string		`json:"password" binding:"required"`
}

type UserLoginRequest struct {
	Name		string		`json:"name" binding:"required"`
	Password	string		`json:"password" binding:"required"`
}