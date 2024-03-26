package config

import (
	"fmt"
	"time"
	"youtube-studio-v2/dtype"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var dbInstance *gorm.DB

func InitializeDatabase(conf *dtype.AppConfig) (*gorm.DB, error) {
	connString := fmt.Sprintf(
		"postgresql://%s:%s@%s/%s?sslmode=require",
		conf.DBUsername,
		conf.DBPassword,
		conf.DBHost,
		conf.DBName,
	)

	db, err := gorm.Open(postgres.Open(connString), &gorm.Config{
		SkipDefaultTransaction: true,
	})
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	sqlDB.SetMaxIdleConns(conf.DBMaxPools)
	sqlDB.SetMaxOpenConns(conf.DBMaxPools)
	sqlDB.SetConnMaxIdleTime(time.Duration(conf.DBMaxIdle * 1000_000_000))

	dbInstance = db
	return db, nil
}

func GetDB() *gorm.DB {
	return dbInstance
}
