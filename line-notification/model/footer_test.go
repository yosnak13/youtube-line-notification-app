package model

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewFooter(t *testing.T) {

	typeOfAction := "uri"
	label := "Youtubeトップへ"
	uri := "https://youtube.com"
	action := NewAction(typeOfAction, label, uri)

	typeOfFooterContent := "button"
	style := "link"
	height := "sm"
	footerContent := NewFooterContent(typeOfFooterContent, style, height, action)

	typeOfFooter := "box"
	layout := "vertical"
	spacing := "sm"
	flex := 0

	expect := &Footer{
		Type:          typeOfFooter,
		Layout:        layout,
		Spacing:       spacing,
		FooterContent: []*FooterContent{footerContent},
		Flex:          flex,
	}
	actual := NewFooter(typeOfFooter, layout, spacing, []*FooterContent{footerContent}, flex)

	assert.Equal(t, expect, actual)
}
