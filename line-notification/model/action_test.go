package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewAction(t *testing.T) {
	typeForAction := "uri"
	label := "YouTubeトップへ"
	uri := "https://example.com"
	expectedAction := &Action{
		Type:  typeForAction,
		Label: label,
		Uri:   uri,
	}

	action := NewAction(typeForAction, label, uri)

	assert.Equal(t, expectedAction, action)
}
