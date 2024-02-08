package model

type Action struct {
	Type  string `json:"type"`
	Label string `json:"label,omitempty"`
	Uri   string `json:"uri"`
}

func NewAction(uri string) *Action {
	return &Action{
		Type: "uri",
		Uri:  uri,
	}
}

func NewActionForFooter() *Action {
	return &Action{
		Type:  "uri",
		Label: "YouTubeトップへ",
		Uri:   "https://youtube.com",
	}
}
