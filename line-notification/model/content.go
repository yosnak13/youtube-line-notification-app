package model

type Content struct {
	Type     string     `json:"type"`
	Text     string     `json:"text,omitempty"`
	Flex     int        `json:"flex,omitempty"`
	Weight   string     `json:"weight,omitempty"`
	Size     string     `json:"size,omitempty"`
	Wrap     bool       `json:"wrap,omitempty"`
	Layout   string     `json:"layout,omitempty"`
	Margin   string     `json:"margin,omitempty"`
	Spacing  string     `json:"spacing,omitempty"`
	Color    string     `json:"color,omitempty"`
	Action   *Action    `json:"action,omitempty"`
	Contents []*Content `json:"contents,omitempty"`
	Style    string     `json:"style,omitempty"`
	Height   string     `json:"height,omitempty"`
}

func NewContentMovieProperty(contentType string, movieTitle string, weight string, size string, wrap bool) *Content {
	return &Content{
		Type:   contentType,
		Text:   movieTitle,
		Weight: weight,
		Size:   size,
		Wrap:   wrap,
	}
}

func NewContentMovieValue(contentType string, layout string, margin string, spacing string, contents []*Content) *Content {
	return &Content{
		Type:     contentType,
		Layout:   layout,
		Margin:   margin,
		Spacing:  spacing,
		Contents: contents,
	}
}

func NewContentBodyBlockChannelRoot(contentType string, layout string, contents []*Content) *Content {
	return &Content{
		Type:     contentType,
		Layout:   layout,
		Contents: contents,
	}
}

func NewContentBodyBlockChannelPropertyValue(contentType string, text string, flex int, wrap bool, size string, color string) *Content {
	return &Content{
		Type:  contentType,
		Text:  text,
		Flex:  flex,
		Wrap:  wrap,
		Size:  size,
		Color: color,
	}
}

func NewContentBodyBlockUrlRoot(contentType string, layout string, spacing string, contents []*Content) *Content {
	return &Content{
		Type:     contentType,
		Layout:   layout,
		Spacing:  spacing,
		Contents: contents,
	}
}

func NewContentBodyBlockUrlProperty(contentType string, movieUrl string, color string, size string, flex int) *Content {
	return &Content{
		Type:  contentType,
		Text:  movieUrl,
		Color: color,
		Size:  size,
		Flex:  flex,
	}
}

func NewContentBodyBlockUrlValue(contentType string, text string, wrap bool, color string, size string, flex int, action *Action) *Content {
	return &Content{
		Type:   contentType,
		Text:   text,
		Wrap:   wrap,
		Color:  color,
		Size:   size,
		Flex:   flex,
		Action: action,
	}
}
