package postgres

import (
	"context"
	"github/meshachdamilare/trimly/model"
)

func (p *Postgres) Store(ctx context.Context, url *model.URL) error {
	db, cancel := p.DBWithTimeout(ctx)
	defer cancel()
	err := db.Create(url).Error
	user := model.User{ID: url.UserId}
	// this line of code below only associate the new url created to the user with the given id
	//and does not necessarily update the URLs column in the User table
	err = db.Model(&user).Association("URLs").Append(url)
	return err
}

func (p *Postgres) GetUrl(ctx context.Context, code string) (*model.URL, error) {
	db, cancel := p.DBWithTimeout(ctx)
	defer cancel()
	var url model.URL
	err := db.First(&url, "code = ?", code).Error
	return &url, err
}

func (p *Postgres) UpdateUrlVisitCount(ctx context.Context, url *model.URL) error {
	db, cancel := p.DBWithTimeout(ctx)
	defer cancel()
	return db.Model(url).
		Where("id = ?", url.ID).
		UpdateColumn("visit_count", url.VisitCount).Error
}

func (p *Postgres) GetUrls(ctx context.Context, userId string) ([]model.URL, error) {
	db, cancel := p.DBWithTimeout(ctx)
	defer cancel()

	var urls []model.URL
	user := model.User{ID: userId}
	err := db.Model(&user).Association("URLs").Find(&urls)
	return urls, err
}
