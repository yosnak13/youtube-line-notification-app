package model

type Footer struct {
	Type          string           `json:"type"`
	Layout        string           `json:"layout"`
	Spacing       string           `json:"spacing"`
	FooterContent []*FooterContent `json:"contents"`
	Flex          int              `json:"flex"`
}

func NewFooter(typeOfFooter string, layout string, spacing string, content []*FooterContent, flex int) *Footer {
	return &Footer{
		Type:          typeOfFooter,
		Layout:        layout,
		Spacing:       spacing,
		FooterContent: content,
		Flex:          flex,
	}
}
