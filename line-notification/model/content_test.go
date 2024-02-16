package model

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewContentMovieTitle(t *testing.T) {

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

	actual := NewContentMovieTitle(contentType, text, weight, size, wrap)

	assert.Equal(t, expect, actual)
}

func TestNewContentBodyContainer(t *testing.T) {

}
