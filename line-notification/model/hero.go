package model

type Hero struct {
	Type         string `json:"type"`
	ThumbnailURL string `json:"url"`
	Size         string `json:"size"`
	AspectRatio  string `json:"aspectRatio"`
	AspectMode   string `json:"aspectMode"`
	Action       Action `json:"action"`
}

func NewHero(thumbnailUrl string, url string) *Hero {
	return &Hero{
		Type:         "image",
		ThumbnailURL: thumbnailUrl,
		Size:         "full",
		AspectRatio:  "20:30",
		AspectMode:   "cover",
		Action:       *NewAction("uri", "", url),
	}
}
