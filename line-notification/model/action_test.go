package model

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestNewAction(t *testing.T) {
	uri := "https://example.com"
	expectedAction := &Action{
		Type: "uri",
		Uri:  uri,
	}

	action := NewAction(uri)

	if !reflect.DeepEqual(action, expectedAction) {
		t.Errorf("NewAction() = %v, want %v", action, expectedAction)
	}
}

func TestNewActionForFooter(t *testing.T) {
	expectedAction := &Action{
		Type:  "uri",
		Label: "YouTubeトップへ",
		Uri:   "https://youtube.com",
	}

	action := NewActionForFooter()

	if !reflect.DeepEqual(action, expectedAction) {
		t.Errorf("NewActionForFooter() = %v, want %v", action, expectedAction)
	}
}

func TestActionJSONMarshal(t *testing.T) {
	action := &Action{
		Type:  "uri",
		Label: "Example",
		Uri:   "https://example.com",
	}
	expectedJSON := `{"type":"uri","label":"Example","uri":"https://example.com"}`

	actionJSON, err := json.Marshal(action)
	if err != nil {
		t.Errorf("Error marshalling Action: %v", err)
	}

	if string(actionJSON) != expectedJSON {
		t.Errorf("Action JSON = %s, want %s", string(actionJSON), expectedJSON)
	}
}
