package model

type Action struct {
	Type  string `json:"type"`
	Label string `json:"label.omitempty"`
	URL   string `json:"url"`
}

func NewAction(url string) *Action {
	return &Action{
		Type: "url",
		URL:  url,
	}
}

func NewActionForFooter() *Action {
	return &Action{
		Type:  "url",
		Label: "YouTubeトップへ",
		URL:   "https://youtube.com",
	}
}
