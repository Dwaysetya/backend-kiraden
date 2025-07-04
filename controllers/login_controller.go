package controllers

import (
	"backend/database"
	"backend/helpers"
	"backend/models"
	"backend/structs"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Login(c *gin.Context) {

	var req = structs.UserLoginRequest{}
	var user = models.User{}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, structs.ErrorResponse{
			Success: false,
			Message: "Validation Errors",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	if err := database.DB.Where("name = ?", req.Name).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, structs.ErrorResponse{
			Success: false,
			Message: "User Not Found",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, structs.ErrorResponse{
			Success: false,
			Message: "Invalid Password",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	token := helpers.GenerateToken(user.Name)

	c.JSON(http.StatusOK, structs.SuccessResponse{
		Success: true,
		Message: "Login Success",
		Data: structs.UserResponse{
			Id:        user.Id,
			Name:      user.Name,
			Email:     user.Email,
			CreatedAt: user.CreatedAt.String(),
			Token:     &token,
		},
	})
}