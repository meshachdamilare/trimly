package model

import (
	"time"
)

type User struct {
	ID        string    `json:"id,omitempty" gorm:"primaryKey;index;unique;not null;type:varchar(50)"`
	Name      string    `json:"name" gorm:"not null;type:varchar(100)"`
	Email     string    `json:"email" gorm:"index;unique;not null;type:varchar(100)"`
	Password  string    `json:"password" gorm:"type:varchar(100);not null"`
	Role      string    `json:"role" gorm:"not null;type:varchar(10)"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	URLs      []URL     `json:"URLs" gorm:"foreignKey:UserId"`
}

type SignUp struct {
	Name            string `json:"name" validate:"required"`
	Email           string `json:"email" validate:"required,email"`
	Password        string `json:"password" validate:"required,min=6"`
	PasswordConfirm string `json:"passwordConfirm" validate:"required,min=6"`
}

type SignIn struct {
	Email    string `json:"email" bson:"email" validate:"required,email"`
	Password string `json:"password" bson:"password" validate:"required"`
}

type UserResponse struct {
	ID        string    `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Email     string    `json:"email,omitempty"`
	Role      string    `json:"role,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type LoginResponse struct {
	AccessToken string
	User        UserResponse
}

func FilteredUserResponse(user *User) UserResponse {
	return UserResponse{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Role:      user.Role,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}
