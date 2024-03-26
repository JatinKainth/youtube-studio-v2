package youtube

import "time"

type VideoDetails struct {
	Items []struct {
		Snippet struct {
			Title       string    `json:"title"`
			Description string    `json:"description"`
			PublishedAt time.Time `json:"publishedAt"`
			Thumbnails  struct {
				Default struct {
					URL string `json:"url"`
				} `json:"default"`
			} `json:"thumbnails"`
		} `json:"snippet"`
	} `json:"items"`
}
