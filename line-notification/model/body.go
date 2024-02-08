package model

type Body struct {
	Type     string   `json:"type"`
	Layout   string   `json:"layout"`
	Contents []string `json:"contents"` // Content構造体をあとで作成
}
