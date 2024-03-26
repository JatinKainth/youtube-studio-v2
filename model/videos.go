package model

import "time"

type Video struct {
	ID          int       `gorm:"primary_key;column:id;type:int" json:"id"`
	Title       string    `gorm:"not null;column:title;type:varchar(255)" json:"title"`
	Description string    `gorm:"column:description;type:text" json:"description"`
	PublishedAt time.Time `gorm:"not null;column:published_at;type:timestamp" json:"published_at"`
	Thumbnail   string    `gorm:"column:thumbnail;type:varchar(255)" json:"thumbnail"`
}

func (Video) TableName() string {
	return "video"
}
