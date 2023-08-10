package handler

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"github/meshachdamilare/trimly/service"
	"github/meshachdamilare/trimly/settings/constant"

	"github/meshachdamilare/trimly/model"
	"github/meshachdamilare/trimly/utils"
	"net/http"
)

type UrlHandler struct {
	urlService service.UrlService
}

func NewUrlHandler(urlService service.UrlService) *UrlHandler {
	return &UrlHandler{
		urlService: urlService,
	}
}

func (handler *UrlHandler) Redirect(c echo.Context) error {
	code := c.Param("code")

	url, err := handler.urlService.Find(code)
	if err != nil {

		if errors.Is(errors.Cause(err), errors.New(constant.ErrRedirectNotFound)) {
			rd := utils.ErrorResponse(http.StatusNotFound, constant.StatusFailed, err.Error(), nil)
			return c.JSON(http.StatusNotFound, rd)
		}
		rd := utils.ErrorResponse(http.StatusInternalServerError, constant.StatusFailed, err.Error(), nil)
		return c.JSON(http.StatusInternalServerError, rd)
	}
	// to debug the visit count
	//fmt.Println(url)
	return c.Redirect(http.StatusMovedPermanently, url.LongUrl)
}

func (handler *UrlHandler) TrimUrl(c echo.Context) error {
	req := new(model.URLRequest)
	if err := c.Bind(&req); err != nil {
		rd := utils.ErrorResponse(http.StatusBadRequest, constant.StatusFailed, err.Error(), nil)
		return c.JSON(http.StatusBadRequest, rd)
	}

	if validationErr := validator.New().Struct(req); validationErr != nil {
		rd := utils.ErrorResponse(http.StatusBadRequest, constant.StatusFailed, validationErr.Error(), nil)
		return c.JSON(http.StatusBadRequest, rd)
	}
	userId := c.Get("userId").(string)
	url, err := handler.urlService.Store(req, userId)
	if err != nil {
		rd := utils.ErrorResponse(http.StatusInternalServerError, constant.StatusFailed, err.Error(), nil)
		return c.JSON(http.StatusInternalServerError, rd)
	}
	rd := utils.SuccessResponse(http.StatusCreated, url)
	return c.JSON(http.StatusOK, rd)
}

func (handler *UrlHandler) GetAllUrls(c echo.Context) error {
	userId := c.Get("userId").(string)
	urls, err := handler.urlService.GetAllURLs(userId)
	if err != nil {
		rd := utils.ErrorResponse(http.StatusInternalServerError, constant.StatusFailed, err.Error(), nil)
		return c.JSON(http.StatusInternalServerError, rd)
	}
	rd := utils.SuccessResponse(http.StatusOK, urls)

	return c.JSON(http.StatusOK, rd)
}
