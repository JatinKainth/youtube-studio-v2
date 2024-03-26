package db

import (
	"context"
	"time"

	"youtube-studio-v2/config"
	"youtube-studio-v2/model"
)

func GetPaginatedVideos(ctx context.Context, limit int, offset int) ([]model.Video, error) {
	db := config.GetDB()

	var obj []model.Video
	err := db.WithContext(ctx).
		Model(&model.Video{}).
		Limit(limit).
		Offset(offset).
		Order("published_at DESC").
		Find(&obj).Error
	return obj, err
}

func Create(ctx context.Context, title, description, thumbnail string, publishedAt time.Time) (*model.Video, error) {
	db := config.GetDB()

	obj := model.Video{
		Title:       title,
		Description: description,
		PublishedAt: publishedAt,
		Thumbnail:   thumbnail,
	}
	err := db.WithContext(ctx).Create(&obj).Error
	return &obj, err
}

func SearchVideos(ctx context.Context, query string) ([]model.Video, error) {
	db := config.GetDB()
	query = "%" + query + "%"

	var obj []model.Video
	err := db.WithContext(ctx).
		Where("title like ? or description like ?", query, query).
		Find(&obj).Error

	return obj, err
}
