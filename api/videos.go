package api

import (
	"context"
	"time"

	"youtube-studio-v2/clients/youtube"
	"youtube-studio-v2/config"
	"youtube-studio-v2/db"
)

func FetchVideosPeriodically(interval time.Duration, keyword string) {
	logger := config.GetLogger()

	for {
		logger.Info("starting fetching videos ...")
		ctx := context.Background()

		// fetch latest youtube video by keyword
		youtubeVideos, err := youtube.GetVideoDetails(keyword)
		if err != nil {
			logger.Errorf("error while fetching videos from youtube, err: %s", err)
			continue
		}

		// save the info in DB to be used by the APIs
		for _, youtubeVideo := range youtubeVideos {
			_, err := db.Create(
				ctx,
				youtubeVideo.Title,
				youtubeVideo.Description,
				youtubeVideo.Thumbnail,
				youtubeVideo.PublishedAt,
			)
			if err != nil {
				logger.Errorf("error while saving youtube video to db, err: %s", err)
				continue
			}
		}

		logger.Infof("videos fetched, going to sleep for %d seconds", int(interval.Seconds()))
		time.Sleep(interval)
	}
}
