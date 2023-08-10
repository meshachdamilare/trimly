package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/teris-io/shortid"
	"github/meshachdamilare/trimly/model"
	"github/meshachdamilare/trimly/repository/storage"
	"gorm.io/gorm"
	"time"
)

type UrlService interface {
	Find(code string) (*model.URL, error)
	Store(urlRequest *model.URLRequest, userId string) (*model.URL, error)
	GetAllURLs(userId string) ([]model.URL, error)
}

type urlService struct {
	dbRepo storage.RepoInterface
}

func NewUrlService(dbRepo storage.RepoInterface) UrlService {
	return &urlService{
		dbRepo: dbRepo,
	}
}

func (u *urlService) Find(code string) (*model.URL, error) {
	url, err := u.dbRepo.GetUrl(context.TODO(), code)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("url-code not found")
		}
		return nil, fmt.Errorf("could not fetch the url-code, got %w", err)
	}

	url.VisitCount++
	err = u.dbRepo.UpdateUrlVisitCount(context.TODO(), url)
	if err != nil {
		return nil, err
	}

	return url, nil
}

func (u *urlService) Store(urlRequest *model.URLRequest, userId string) (*model.URL, error) {

	url := new(model.URL)
	code := shortid.MustGenerate()

	{
		url.ID = uuid.New().String()
		url.Code = code
		url.LongUrl = urlRequest.LongUrl
		url.CreatedAt = time.Now().UTC().Unix()
		url.UserId = userId
	}

	err := u.dbRepo.Store(context.TODO(), url)
	if err != nil {
		return nil, err
	}
	return url, err
}

func (u *urlService) GetAllURLs(userId string) ([]model.URL, error) {
	urls, err := u.dbRepo.GetUrls(context.TODO(), userId)
	if err != nil {
		return nil, err
	}
	return urls, nil
}
