package model

type URL struct {
	ID         string `json:"id" gorm:"column:id;primaryKey;index;unique;not null;type:varchar(50)"`
	Code       string `json:"code" gorm:"column:code;index;unique;not null"`
	LongUrl    string `json:"long_url" gorm:"column:long_url;not null"`
	VisitCount int    `json:"visit_count" gorm:"column:visit_count"`
	UserId     string `json:"user_id" gorm:"column:user_id;index"`
	CreatedAt  int64  `json:"created_at" gorm:"column:created_at;index"`
}

type URLRequest struct {
	LongUrl string `json:"long_url" validate:"required,url"`
}
