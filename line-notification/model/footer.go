package model

type Footer struct {
	Type          string        `json:"type"`
	Layout        string        `json:"layout"`
	Spacing       string        `json:"spacing"`
	FooterContent FooterContent `json:"contents"` // []にする必要がありそう
	Flex          int           `json:"flex"`
}

func NewFooter() *Footer {
	return &Footer{
		Type:          "box",
		Layout:        "vertical",
		Spacing:       "sm",
		FooterContent: *NewFooterContent(),
		Flex:          0,
	}
}
