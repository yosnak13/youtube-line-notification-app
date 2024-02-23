package model

type Body struct {
	Type     string     `json:"type"`
	Layout   string     `json:"layout"`
	Contents []*Content `json:"contents"`
}

func NewBody(contentType string, layout string, contents []*Content) *Body {
	return &Body{
		Type:     contentType,
		Layout:   layout,
		Contents: contents,
	}
}
