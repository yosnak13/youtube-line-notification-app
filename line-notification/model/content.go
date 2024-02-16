package model

type Content struct {
	Type     string     `json:"type"`
	Text     string     `json:"text,omitempty"`
	Weight   string     `json:"weight,omitempty"`
	Size     string     `json:"size,omitempty"`
	Wrap     bool       `json:"wrap,omitempty"`
	Layout   string     `json:"layout,omitempty"`
	Margin   string     `json:"margin,omitempty"`
	Spacing  string     `json:"spacing,omitempty"`
	Color    string     `json:"color,omitempty"`
	Flex     int        `json:"flex,omitempty"`
	Action   *Action    `json:"action,omitempty"`
	Contents []*Content `json:"contents,omitempty"`
	Style    string     `json:"style,omitempty"`
	Height   string     `json:"height,omitempty"`
}

func NewContentMovieTitle(contentType string, movieTitle string, weight string, size string, wrap bool) *Content {
	return &Content{
		Type:   contentType,
		Text:   movieTitle,
		Weight: weight,
		Size:   size,
		Wrap:   wrap,
	}
}

func NewContentBodyContainer(contentType string, layout string, margin string, spacing string, contents []*Content) *Content {
	// Bodyを構成する4つのContentを格納する
	return &Content{
		Type:     contentType,
		Layout:   layout,
		Margin:   margin,
		Spacing:  spacing,
		Contents: contents,
	}
}
