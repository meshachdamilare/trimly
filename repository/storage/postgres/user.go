package postgres

import (
	"context"
	"github/meshachdamilare/trimly/model"
	"gorm.io/gorm/clause"
	"strings"
)

// CreateUser store 'user' in the database
func (p *Postgres) CreateUser(ctx context.Context, user *model.User) error {
	db, cancel := p.DBWithTimeout(ctx)
	defer cancel()
	return db.Create(user).Error
}

// GetUserByIdOrEmail fetches a user from the database using its 'id' or 'email'
func (p *Postgres) GetUserByIdOrEmail(ctx context.Context, idOrEmail string) (*model.User, error) {
	db, cancel := p.DBWithTimeout(ctx)
	defer cancel()

	// check if 'idOrEmail' is an email or an id
	var cond string
	if strings.Contains(idOrEmail, "@") {
		cond = "email = ?"
	} else {
		cond = "id = ?"
	}
	var user model.User
	err := db.First(&user, cond, idOrEmail).Error
	return &user, err
}

// GetAllUsers gets all users in the database
func (p *Postgres) GetAllUsers(ctx context.Context) ([]model.User, error) {
	db, cancel := p.DBWithTimeout(ctx)
	defer cancel()

	var users []model.User
	err := db.Find(&users).Error
	return users, err
}

// UpdateUser updates some specific user fields
func (p *Postgres) UpdateUser(ctx context.Context, id string, user *model.User) error {
	db, cancel := p.DBWithTimeout(ctx)
	defer cancel()

	// Ensure 'is_verified', 'password', 'email' cannot be updated using this function
	err := db.Model(user).Clauses(clause.Returning{}).
		Omit("password", "email", "role").
		Where("id = ?", id).Updates(user).Error

	return err
}

func (p *Postgres) GetUserUrls(ctx context.Context, id string) ([]model.URL, error) {
	db, cancel := p.DBWithTimeout(ctx)
	defer cancel()

	var user model.User
	err := db.Preload("URLs").First(&user, "id = ?", id).Error
	return user.URLs, err
}
