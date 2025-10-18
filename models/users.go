package models

import (
	"fmt"
	"login-system/utils/constants"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

const (
	TableNameUsers = "users"
	TableAsUsers   = "u"
)

type User struct {
	ID        uuid.UUID `gorm:"column:id;type:uuid;default:uuid_generate_v4();primaryKey"`
	FullName  string    `gorm:"column:full_name;not null"`
	Email     string    `gorm:"column:email;not null;unique"`
	Password  string    `gorm:"column:password;not null"`
	CreatedAt time.Time `gorm:"column:created_at;not null;default:now()"`
}

type GetUserRequest struct {
	ID    uuid.UUID
	Email string
}

func (user *User) TableName() string {
	return fmt.Sprintf("%s.%s", constants.Schema, TableNameUsers)
}

func (user *User) TableAs() string {
	return fmt.Sprintf("%s.%s", constants.Schema, TableAsUsers)
}

func CreateUser(db *gorm.DB, user *User) (*uuid.UUID, error) {
	var newUser uuid.UUID
	err := db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Table(user.TableName()).Create(&user).Error; err != nil {
			return fmt.Errorf("failed to create user: %v", err)
		}

		newUser = user.ID
		return nil
	})
	if err != nil {
		return &uuid.UUID{}, fmt.Errorf("failed to create user: %v", err)
	}
	return &newUser, nil
}

func GetUser(db *gorm.DB, req GetUserRequest) (*User, error) {
	var user User
	query := db.Table(user.TableName())
	if req.ID != uuid.Nil {
		query.Where("id = ?", req.ID)
	}

	if req.Email != "" {
		query.Where("email = ?", req.Email)
	}

	if err := query.Scan(&user).Error; err != nil {
		return nil, fmt.Errorf("failed to get user: %v", err)
	}
	return &user, nil
}
