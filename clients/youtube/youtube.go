package youtube

import (
	"encoding/json"
	"net/http"
	"time"

	"youtube-studio-v2/config"
)

type YoutubeVideo struct {
	Title       string    `json:"title"`
	Description string    `json:"description"`
	PublishedAt time.Time `json:"publishedAt"`
	Thumbnail   string    `json:"thumbnail"`
}

const YouTubeAPIURL = "https://www.googleapis.com/youtube/v3"

func GetVideoDetails(query string) ([]YoutubeVideo, error) {
	url := YouTubeAPIURL + "/search?part=snippet&q=" + query + "&type=video&order=date&key=" + config.GetNextAPIKey()

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var data VideoDetails
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}

	if len(data.Items) == 0 {
		return nil, nil
	}

	var res []YoutubeVideo
	for _, item := range data.Items {
		res = append(res, YoutubeVideo{
			Title:       item.Snippet.Title,
			Description: item.Snippet.Description,
			PublishedAt: item.Snippet.PublishedAt,
			Thumbnail:   item.Snippet.Thumbnails.Default.URL,
		})
	}

	return res, nil
}
