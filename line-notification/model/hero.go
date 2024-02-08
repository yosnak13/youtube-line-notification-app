package model

type Hero struct {
	Type        string `json:"type"`
	URL         string `json:"url"`
	Size        string `json:"size"`
	AspectRatio string `json:"aspectRatio"`
	AspectMode  string `json:"aspectMode"`
	Action      Action `json:"action"` // Action型をあとで作る
}

func NewHero(thumbnails string, url string) *Hero {
	return &Hero{
		Type:        "image",
		URL:         thumbnails,
		Size:        "full",
		AspectRatio: "20:30",
		AspectMode:  "cover",
		Action:      *NewAction(url), // MEMO: Actionを知らなければならず、結合度が高くなっている
	}
}

func (h *Hero) getThumbnail() string {
	return h.URL
}
