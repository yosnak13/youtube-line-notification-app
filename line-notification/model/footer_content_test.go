package model

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewFooterContent(t *testing.T) {

	typeOfAction := "uri"
	label := "Youtubeトップへ"
	uri := "https://youtube.com"
	action := *NewAction(typeOfAction, label, uri)

	typeOfFooterContent := "button"
	style := "link"
	height := "sm"

	expect := &FooterContent{
		Type:   typeOfFooterContent,
		Style:  style,
		Height: height,
		Action: &action,
	}

	actual := NewFooterContent(typeOfFooterContent, style, height, &action)

	assert.Equal(t, expect.getHeight(), actual.getHeight())
	assert.Equal(t, expect.getStyle(), actual.getStyle())
	assert.Equal(t, expect.getType(), actual.getType())
	assert.Equal(t, expect.getAction(), actual.getAction())
}
