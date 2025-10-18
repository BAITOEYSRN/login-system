package controllers

import (
	"errors"
	"fmt"
	"login-system/models"
	"login-system/utils"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRequest struct {
	FullName string `json:"full_name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}
type RegisterUserResponse struct {
	UserID uuid.UUID `json:"user_id"`
}

func RegisterUser(c *gin.Context, db *gorm.DB) (*uuid.UUID, error) {
	user, err := validation(c, new(UserRequest))
	if err != nil {
		return nil, err
	}

	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return nil, err
	}

	user.Password = hashedPassword

	existingUser, err := models.GetUser(db, models.GetUserRequest{Email: user.Email})
	if err != nil {
		return nil, errors.New("failed to get user")
	}
	if existingUser != nil {
		return nil, errors.New("email already exists")
	}

	userID, err := models.CreateUser(db, &models.User{
		FullName: user.FullName,
		Email:    user.Email,
		Password: user.Password,
	})
	if err != nil {
		return nil, err
	}

	return userID, nil
}

func validation(c *gin.Context, user *UserRequest) (*UserRequest, error) {
	if err := c.ShouldBindJSON(&user); err != nil {
		return nil, fmt.Errorf("invalid request body: %v", err)
	}

	validate := validator.New()
	if err := validate.Struct(user); err != nil {
		return nil, fmt.Errorf("invalid request body: %v", err)
	}

	return &UserRequest{
		FullName: user.FullName,
		Email:    user.Email,
		Password: user.Password,
	}, nil
}
