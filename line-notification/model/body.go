package model

type Body struct {
	Type     string    `json:"type"`
	Layout   string    `json:"layout"`
	Contents []Content `json:"contents"`
}
