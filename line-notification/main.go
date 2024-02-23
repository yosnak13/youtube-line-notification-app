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
			videoID := item.Id.VideoId

			fmt.Printf("Channel ID: %s\n", channelID)
			fmt.Printf("Title: %s\n", videoTitle)
			fmt.Printf("Thumbnail URL: %s\n", thumbnails.Default.Url)
			fmt.Printf("Channel Name: %s\n", channelTitle)
			videoURL := fmt.Sprintf("https://www.youtube.com/watch?v=%s", videoID)
			fmt.Printf("Video URL: %s\n", videoURL)

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
	Type   string        `json:"type"`
	Hero   *model.Hero   `json:"hero"`
	Body   *Body         `json:"body"`
	Footer *model.Footer `json:"footer"`
}

type Body struct {
	Type     string    `json:"type"`
	Layout   string    `json:"layout"`
	Contents []Content `json:"contents"`
}

type Content struct {
	Type     string        `json:"type"`
	Text     string        `json:"text,omitempty"`
	Weight   string        `json:"weight,omitempty"`
	Size     string        `json:"size,omitempty"`
	Wrap     bool          `json:"wrap,omitempty"`
	Layout   string        `json:"layout,omitempty"`
	Margin   string        `json:"margin,omitempty"`
	Spacing  string        `json:"spacing,omitempty"`
	Color    string        `json:"color,omitempty"`
	Flex     int           `json:"flex,omitempty"`
	Action   *model.Action `json:"action,omitempty"`
	Contents []Content     `json:"contents,omitempty"`
	Style    string        `json:"style,omitempty"`
	Height   string        `json:"height,omitempty"`
}

func buildJson() *Message {

	hero := buildHero()
	body := buildBody()
	footer := buildFooter()

	return &Message{
		Type:   "bubble",
		Hero:   hero,
		Body:   body,
		Footer: footer,
	}
}

func buildHero() *model.Hero {
	thumbnailURL := "https://example.com"

	typeOfAction := "uri"
	uri := "https://youtube.com"
	action := *model.NewAction(typeOfAction, "", uri)
	hero := *model.NewHero("image", thumbnailURL, "full", "20:30", "cover", &action)
	return &hero
}

func buildBody() *Body {

	videoURL := "https://www.youtube.com"
	channelName := "channelTitle"

	urlProperty := *model.NewContentBodyBlockUrlProperty("text", "URL", "#aaaaaa", "sm", 1)
	urlValueAction := *model.NewAction("url", "", "movieのURL")
	urlValue := *model.NewContentBodyBlockUrlValue("text", "動画はこちらをタップ", true, "#666666", "sm", 5, &urlValueAction)

	urlComponents := []*model.Content{&urlProperty, &urlValue}
	urlRootComponent := *model.NewContentBodyBlockUrlRoot("box", "baseline", "sm", urlComponents)

	channelProperty := *model.NewContentBodyBlockChannelPropertyValue("text", "ch", 1, true, "sm", "#aaaaaa")
	channelValue := *model.NewContentBodyBlockChannelPropertyValue("text", "チャンネル名", 5, true, "sm", "#aaaaaa")
	channelNameComponents := []*model.Content{&channelProperty, &channelValue}
	channelRootComponent := *model.NewContentBodyBlockChannelRoot("box", "baseline", channelNameComponents)

	// titleContainer := *model.NewContentMovieTitle("text", "動画のタイトル", "bold", "xl", true)
	// bodyParentContainer := *model.NewContentBodyContainer("box", "vertical", "lg", "sm", )
	// var contents []model.Content
	// contents = append(contents, titleContainer)

	var movieComponents []*model.Content
	movieComponents = append(movieComponents, &channelRootComponent)
	movieComponents = append(movieComponents, &urlRootComponent)
	movieValue := *model.NewContentMovieValue("box", "vertical", "lg", "sm", movieComponents)

	return &Body{
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
								Action: &model.Action{
									Type: "uri",
									Uri:  videoURL,
								},
							},
						},
					},
				},
			},
		},
	}
}

func buildFooter() *model.Footer {
	typeOfAction := "uri"
	label := "Youtubeトップへ"
	uri := "https://youtube.com"
	action := *model.NewAction(typeOfAction, label, uri)

	typeOfFooterContent := "button"
	style := "link"
	height := "sm"
	footerContent := *model.NewFooterContent(typeOfFooterContent, style, height, &action)

	typeOfFooter := "box"
	layout := "vertical"
	spacing := "sm"
	content := []*model.FooterContent{&footerContent}
	flex := 0

	footer := *model.NewFooter(typeOfFooter, layout, spacing, content, flex)
	return &footer
}
