package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github/meshachdamilare/trimly/api/middleware"
	"github/meshachdamilare/trimly/model"
	"github/meshachdamilare/trimly/repository/storage"
	"github/meshachdamilare/trimly/settings/constant"
	"github/meshachdamilare/trimly/utils"
	"gorm.io/gorm"
	"time"
)

type UserService interface {
	CreateUser(user *model.SignUp) error
	Login(req *model.SignIn) (*model.LoginResponse, error)
	GetUserByIdOrEmail(idOrEmail string) (*model.User, error)
	GetUserURLs(userId string) ([]model.URL, error)
	GetAllUsers() ([]model.User, error)
	UpdateUser(sid string, user *model.User) error
}

type userService struct {
	dbRepo storage.RepoInterface
}

func NewUserService(dbRepo storage.RepoInterface) UserService {
	return &userService{
		dbRepo: dbRepo,
	}
}

func (u *userService) CreateUser(signinReq *model.SignUp) error {
	user := new(model.User)
	if signinReq.Password != signinReq.PasswordConfirm {
		return fmt.Errorf("passwords not match")
	}
	passwordHash, err := utils.HashPassword(signinReq.Password)
	if err != nil {
		return fmt.Errorf("couldn't hash the password: %v", err)
	}

	{
		user.ID = uuid.New().String()
		user.Name = signinReq.Name
		user.Email = signinReq.Email
		user.Password = passwordHash
		user.CreatedAt = time.Now()
		user.UpdatedAt = time.Now()
		user.Role = constant.User
	}
	err = u.dbRepo.CreateUser(context.TODO(), user)
	if err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return fmt.Errorf("duplicate key found: %v", err)
		}
		return fmt.Errorf("error ocuured: %v", err)
	}
	return nil
}

func (u *userService) Login(req *model.SignIn) (*model.LoginResponse, error) {
	user, err := u.dbRepo.GetUserByIdOrEmail(context.TODO(), req.Email)
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("could not get user, got error %w", err)
		}
		return nil, fmt.Errorf("invalid user")
	}

	if err := utils.VerifyPassword(user.Password, req.Password); err != nil {
		return nil, fmt.Errorf("invalid password")
	}

	accessToken, err := middleware.CreateToken(30, user.ID, user.Email, user.Role)

	if err != nil {
		return nil, fmt.Errorf("could not create token")
	}

	// Omit user Password
	user.Password = ""

	// Ensure that user is verified

	// return users credentials
	return &model.LoginResponse{
		AccessToken: accessToken,
		User:        model.FilteredUserResponse(user),
	}, nil

}

func (u *userService) GetUserByIdOrEmail(idOrEmail string) (*model.User, error) {
	user, err := u.dbRepo.GetUserByIdOrEmail(context.TODO(), idOrEmail)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("email or id not found")
		}
		return nil, err
	}
	return user, nil
}

func (u *userService) GetUserURLs(userId string) ([]model.URL, error) {
	urls, err := u.dbRepo.GetUserUrls(context.TODO(), userId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("user not found")
		}
		return nil, err
	}
	return urls, nil
}

func (u *userService) GetAllUsers() ([]model.User, error) {
	users, err := u.dbRepo.GetAllUsers(context.TODO())
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (u *userService) UpdateUser(id string, user *model.User) error {
	err := u.dbRepo.UpdateUser(context.TODO(), id, user)
	if err != nil {
		return err
	}
	return nil
}
