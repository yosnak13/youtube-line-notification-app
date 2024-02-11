package model

type Hero struct {
	Type         string  `json:"type"`
	ThumbnailURL string  `json:"url"`
	Size         string  `json:"size"`
	AspectRatio  string  `json:"aspectRatio"`
	AspectMode   string  `json:"aspectMode"`
	Action       *Action `json:"action"`
}

func NewHero(typeOfHero string, thumbnailUrl string, size string, aspectRatio string, aspectMode string, action *Action) *Hero {
	return &Hero{
		Type:         typeOfHero,
		ThumbnailURL: thumbnailUrl,
		Size:         size,
		AspectRatio:  aspectRatio,
		AspectMode:   aspectMode,
		Action:       action,
	}
}
