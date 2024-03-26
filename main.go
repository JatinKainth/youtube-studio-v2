package main

import (
	"log"
	"time"

	"youtube-studio-v2/api"
	"youtube-studio-v2/config"

	"github.com/labstack/echo/v4"
)

const (
	fetchInterval = time.Second * 10
	fetchKeyword  = "cricket"
)

func main() {
	err := InitializeApplication()
	if err != nil {
		log.Fatalf("failed to initiailse the application, err: %s", err)
	}

	router := echo.New()
	router.GET("/videos", api.GetVideosHandler)
	router.GET("/search", api.SearchVideosHandler)

	go api.FetchVideosPeriodically(fetchInterval, fetchKeyword)
	log.Fatal(router.Start(":8080"))
}

func InitializeApplication() error {
	logger := config.InitializeLog()

	conf, err := config.InitializeAppConfig()
	if err != nil {
		logger.Fatalf("Failed to load configuration: %v", err)
		return err
	}

	_, err = config.InitializeDatabase(conf)
	if err != nil {
		logger.Fatalf("Failed to load configuration: %v", err)
		return err
	}

	return nil
}
