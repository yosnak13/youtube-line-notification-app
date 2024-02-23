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

	/*
		Content構造体はBody構造体の部品に使われており、下記のようなContentの配列をいくつも抱える構造をしている。
		同名プロパティが存在するため、omitemptyを活用し、必要なプロパティのみ引数で必要とするコンストラクタを必要なだけ作成し、
		それで生成されたインスタンスを配列に格納することを繰り返して、以下の構造体を作る。
	*/

	/*
		Body{
			Type:   "box",
			Layout: "vertical",
			Contents: []Content{
				{ * NewContentMoviePropertyで生成
					Type:   "text",
					Text:   "タイトル",
					Weight: "bold",
					Size:   "xl",
					Wrap:   true,
				},
				{ * NewContentMovieValueで生成
					Type:    "box",
					Layout:  "vertical",
					Margin:  "lg",
					Spacing: "sm",
					Contents: []Content{
						{ * NewContentBodyBlockChannelRootで生成
							Type:   "box",
							Layout: "baseline",
							Contents: []Content{
								{ * NewContentBodyBlockChannelPropertyValueで生成
									Type:  "text",
									Text:  "ch",
									Flex:  1,
									Wrap:  true,
									Size:  "sm",
									Color: "#aaaaaa",
								},
								{ * NewContentBodyBlockChannelPropertyValueで生成
									Type:  "text",
									Text:  channelName,
									Flex:  5,
									Wrap:  true,
									Size:  "sm",
									Color: "#aaaaaa",
								},
							},
						},
						{ * NewContentBodyBlockUrlRootで生成
							Type:    "box",
							Layout:  "baseline",
							Spacing: "sm",
							Contents: []Content{
								{ * NewContentBodyBlockUrlPropertyで生成
									Type:  "text",
									Text:  "URL",
									Color: "#aaaaaa",
									Size:  "sm",
									Flex:  1,
								},
								{ * NewContentBodyBlockUrlValueで生成
									Type:  "text",
									Text:  "動画はこちらをタップ",
									Wrap:  true,
									Color: "#666666",
									Size:  "sm",
									Flex:  5,
									Action: &model.Action{
										Type: "uri",
										Uri:  videoURL,
									},
								},
							},
						},
					},
				},
			},
		}
	*/

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
