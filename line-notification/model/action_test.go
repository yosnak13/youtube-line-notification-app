package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewAction(t *testing.T) {
	testCases := []struct {
		name           string
		typeForAction  string
		label          string
		uri            string
		expectedAction *Action
	}{
		{
			name:          "Type and URI only",
			typeForAction: "url",
			label:         "",
			uri:           "https://example.com",
			expectedAction: &Action{
				Type: "url",
				Uri:  "https://example.com",
			},
		},
		{
			name:          "All fields provided",
			typeForAction: "uri",
			label:         "Example",
			uri:           "https://example.com",
			expectedAction: &Action{
				Type:  "uri",
				Label: "Example",
				Uri:   "https://example.com",
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			action := NewAction(tc.typeForAction, tc.label, tc.uri)
			assert.Equal(t, tc.expectedAction, action)
		})
	}
}
