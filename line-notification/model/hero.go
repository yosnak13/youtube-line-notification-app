package model

type Hero struct {
	Type         string `json:"type"`
	ThumbnailURL string `json:"url"`
	Size         string `json:"size"`
	AspectRatio  string `json:"aspectRatio"`
	AspectMode   string `json:"aspectMode"`
	Action       Action `json:"action"`
}

func NewHero(thumbnail string, url string) *Hero {
	return &Hero{
		Type:         "image",
		ThumbnailURL: thumbnail,
		Size:         "full",
		AspectRatio:  "20:30",
		AspectMode:   "cover",
		Action:       *NewAction(url), // MEMO: Actionを知らなければならず、結合度が高くなっている
	}
}
