package main

import (
	"encoding/json"
	"fmt"
	"google.golang.org/api/youtube/v3"
	"line-notification/model"
	"log"
	"net/http"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
	"google.golang.org/api/googleapi/transport"
)

func handler() {
	channelIDs := []string{
		"UCPVr7clenPjpD7WNsSI3UBQ", // レトルト
		"UCZMRuagdTBKmmrFtSMN48Xw", // 牛沢
		"UCWcEgYIOqq1BVr4Qm1sPuVg", // ガッチマン
		// 追加のチャンネルIDをここに追加
	}

	apiKey := os.Getenv("API_KEY")
	client := &http.Client{
		Transport: &transport.APIKey{Key: apiKey},
	}

	service, err := youtube.New(client)
	if err != nil {
		log.Fatalf("Error creating new YouTube client: %v", err)
	}

	for _, channelID := range channelIDs {
		call := service.Search.List([]string{"snippet"}).ChannelId(channelID).Type("video").Order("date").MaxResults(1)
		response, err := call.Do()
		if err != nil {
			log.Fatalf("Error making search API call for channel %s: %v", channelID, err)
		}

		for _, item := range response.Items {
			videoTitle := item.Snippet.Title
			thumbnails := item.Snippet.Thumbnails
			channelTitle := item.Snippet.ChannelTitle

			fmt.Printf("Channel ID: %s\n", channelID)
			fmt.Printf("Title: %s\n", videoTitle)
			fmt.Printf("Thumbnail URL: %s\n", thumbnails.Default.Url)
			fmt.Printf("Channel Name: %s\n", channelTitle)
			fmt.Println()
		}
	}

	sampleMessage := buildJson()
	messageJSON, err := json.Marshal(sampleMessage)
	if err != nil {
		fmt.Println("JSON marshal error:", err)
		return
	}

	fmt.Println(string(messageJSON))
}

func main() {
	lambda.Start(handler)
}

// MEMO: 徐々に切り出していく

type Message struct {
	Type   string       `json:"type"`
	Hero   Hero         `json:"hero"`
	Body   Body         `json:"body"`
	Footer model.Footer `json:"footer"`
}

type Hero struct {
	Type        string `json:"type"`
	URL         string `json:"url"`
	Size        string `json:"size"`
	AspectRatio string `json:"aspectRatio"`
	AspectMode  string `json:"aspectMode"`
	Action      Action `json:"action"`
}

type Action struct {
	Type  string `json:"type"`
	URI   string `json:"uri"`
	Label string `json:"label,omitempty"`
}

type Body struct {
	Type     string    `json:"type"`
	Layout   string    `json:"layout"`
	Contents []Content `json:"contents"`
}

type Content struct {
	Type     string    `json:"type"`
	Text     string    `json:"text,omitempty"`
	Weight   string    `json:"weight,omitempty"`
	Size     string    `json:"size,omitempty"`
	Wrap     bool      `json:"wrap,omitempty"`
	Layout   string    `json:"layout,omitempty"`
	Margin   string    `json:"margin,omitempty"`
	Spacing  string    `json:"spacing,omitempty"`
	Color    string    `json:"color,omitempty"`
	Flex     int       `json:"flex,omitempty"`
	Action   Action    `json:"action,omitempty"`
	Contents []Content `json:"contents,omitempty"`
	Style    string    `json:"style,omitempty"`
	Height   string    `json:"height,omitempty"`
}

func buildJson() *Message {
	thumbnailURL := "thumbnails"
	channelName := "channelTitle"
	videoURL := "https://www.youtube.com"

	footer := *model.NewFooter()
	hero := *model.NewHero(thumbnailURL, videoURL)

	return &Message{
		Type: "bubble",
		Hero: Hero{
			Type:        "image",
			URL:         thumbnailURL,
			Size:        "full",
			AspectRatio: "20:13",
			AspectMode:  "cover",
			Action: Action{
				Type: "uri",
				URI:  videoURL,
			},
		},
		Body: Body{
			Type:   "box",
			Layout: "vertical",
			Contents: []Content{
				{
					Type:   "text",
					Text:   "タイトル",
					Weight: "bold",
					Size:   "xl",
					Wrap:   true,
				},
				{
					Type:    "box",
					Layout:  "vertical",
					Margin:  "lg",
					Spacing: "sm",
					Contents: []Content{
						{
							Type:   "box",
							Layout: "baseline",
							Contents: []Content{
								{
									Type:  "text",
									Text:  "ch",
									Flex:  1,
									Wrap:  true,
									Size:  "sm",
									Color: "#aaaaaa",
								},
								{
									Type:  "text",
									Text:  channelName,
									Flex:  5,
									Wrap:  true,
									Size:  "sm",
									Color: "#aaaaaa",
								},
							},
						},
						{
							Type:    "box",
							Layout:  "baseline",
							Spacing: "sm",
							Contents: []Content{
								{
									Type:  "text",
									Text:  "URL",
									Color: "#aaaaaa",
									Size:  "sm",
									Flex:  1,
								},
								{
									Type:  "text",
									Text:  "動画はこちらをタップ",
									Wrap:  true,
									Color: "#666666",
									Size:  "sm",
									Flex:  5,
									Action: Action{
										Type: "uri",
										URI:  videoURL,
									},
								},
							},
						},
					},
				},
			},
		},
		Footer: footer,
	}
}
