package model

type Action struct {
	Type  string `json:"type"`
	Label string `json:"label,omitempty"`
	Uri   string `json:"uri"`
}

func NewAction(typeForAction string, label string, uri string) *Action {
	if label == "" {
		return &Action{
			Type: typeForAction,
			Uri:  uri,
		}
	}
	return &Action{
		Type:  typeForAction,
		Label: label,
		Uri:   uri,
	}
}
