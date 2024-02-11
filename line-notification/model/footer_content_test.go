package model

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewFooterContent(t *testing.T) {

	typeOfAction := "uri"
	label := "Youtubeトップへ"
	uri := "https://youtube.com"
	action := NewAction(typeOfAction, label, uri) // ポインターを解除して渡すことで、

	typeOfFooterContent := "button"
	style := "link"
	height := "sm"

	expect := &FooterContent{
		Type:   typeOfFooterContent,
		Style:  style,
		Height: height,
		Action: action,
	}

	actual := NewFooterContent(typeOfFooterContent, style, height, action)

	assert.Equal(t, expect, actual)
}
