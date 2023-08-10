package model

type URL struct {
	ID         string `json:"id" gorm:"primaryKey;index;unique;not null;type:varchar(50)"`
	Code       string `json:"code" gorm:"index;unique;not null"`
	LongUrl    string `json:"long_url" gorm:"not null"`
	VisitCount int    `json:"visit_count" gorm:"not null"`
	UserId     string `json:"user_id" gorm:"index"`
	CreatedAt  int64  `json:"created_at" gorm:"index"`
}

type URLRequest struct {
	LongUrl string `json:"long_url" validate:"required,url"`
}
