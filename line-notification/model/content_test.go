package model

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewContentOfBodyTitle(t *testing.T) {

	contentType := "text"
	text := "タイトル"
	weight := "bold"
	size := "xl"
	wrap := true

	expect := &Content{
		Type:   contentType,
		Text:   text,
		Weight: weight,
		Size:   size,
		Wrap:   wrap,
	}

	actual := NewContentOfBodyTitle(contentType, text, weight, size, wrap)

	assert.Equal(t, expect, actual)
}
