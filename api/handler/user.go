package handler

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github/meshachdamilare/trimly/service"
	"github/meshachdamilare/trimly/settings/constant"

	"github/meshachdamilare/trimly/model"
	"github/meshachdamilare/trimly/utils"
	"net/http"
)

type UserHandler struct {
	userService service.UserService
}

func NewUserHandler(userService service.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

func (handler *UserHandler) Register(c echo.Context) error {
	var payload model.SignUp
	if err := c.Bind(&payload); err != nil {
		rd := utils.ErrorResponse(http.StatusBadRequest, constant.StatusFailed, err.Error(), "Binding error")
		return c.JSON(http.StatusBadRequest, rd)
	}
	err := handler.userService.CreateUser(&payload)
	if err != nil {
		rd := utils.ErrorResponse(http.StatusInternalServerError, constant.StatusFailed, err.Error(), "Internal server error")
		return c.JSON(http.StatusNotFound, rd)
	}
	rd := utils.SuccessResponse(http.StatusOK, "user created")
	return c.JSON(http.StatusCreated, rd)
}

func (handler *UserHandler) Login(c echo.Context) error {
	req := new(model.SignIn)
	if err := c.Bind(&req); err != nil {
		rd := utils.ErrorResponse(http.StatusBadRequest, constant.StatusFailed, err.Error(), "Binding error")
		return c.JSON(http.StatusBadRequest, rd)
	}
	if validationErr := validator.New().Struct(req); validationErr != nil {
		rd := utils.ErrorResponse(http.StatusBadRequest, constant.StatusFailed, validationErr.Error(), "Validation error")
		return c.JSON(http.StatusBadRequest, rd)
	}
	login, err := handler.userService.Login(req)
	if err != nil {
		rd := utils.ErrorResponse(http.StatusBadRequest, constant.StatusFailed, err.Error(), "Internal server error")
		return c.JSON(http.StatusBadRequest, rd)
	}

	accessCookie := CreateCookie("access_token", login.AccessToken, 60, c, true)
	loggedInCookie := CreateCookie("logged_in", "true", 60, c, false)

	c.SetCookie(accessCookie)
	c.SetCookie(loggedInCookie)

	rd := utils.SuccessResponse(http.StatusCreated, login)
	return c.JSON(http.StatusOK, rd)

}

func (handler *UserHandler) Me(c echo.Context) error {
	userId := c.Get("userId").(string)
	fmt.Println("Reach here: ", userId)
	user, err := handler.userService.GetUserByIdOrEmail(userId)
	//fmt.Println("Reach here")
	if err != nil {
		rd := utils.ErrorResponse(http.StatusInternalServerError, constant.StatusFailed, err.Error(), "Internal server error")
		return c.JSON(http.StatusInternalServerError, rd)
	}
	rd := utils.SuccessResponse(http.StatusOK, model.FilteredUserResponse(user))
	return c.JSON(http.StatusOK, rd)
}

func (handler *UserHandler) GetUserURLs(c echo.Context) error {
	userId := c.Get("userId").(string)
	urls, err := handler.userService.GetUserURLs(userId)
	if err != nil {
		rd := utils.ErrorResponse(http.StatusInternalServerError, constant.StatusFailed, err.Error(), "Internal server error")
		return c.JSON(http.StatusInternalServerError, rd)
	}
	rd := utils.SuccessResponse(http.StatusOK, urls)
	return c.JSON(http.StatusOK, rd)
}

func (handler *UserHandler) LogoutUser(c echo.Context) error {
	accessCookie := CreateCookie("access_token", "", -1, c, true)
	loggedInCookie := CreateCookie("logged_in", "true", -1, c, true)

	c.SetCookie(accessCookie)
	c.SetCookie(loggedInCookie)

	rd := utils.SuccessResponse(http.StatusOK, nil)
	return c.JSON(http.StatusOK, rd)
}
