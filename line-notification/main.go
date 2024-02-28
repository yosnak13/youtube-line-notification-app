package main

import (
	"encoding/json"
	"fmt"
	"github.com/line/line-bot-sdk-go/v7/linebot"
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

	var bubbles []*model.Bubble

	for _, channelID := range channelIDs {
		call := service.Search.List([]string{"snippet"}).ChannelId(channelID).Type("video").Order("date").MaxResults(1)
		response, err := call.Do()
		if err != nil {
			log.Fatalf("Error making search API call for channel %s: %v", channelID, err)
		}

		item := response.Items[0] // MaxResultsが1なので、配列の最初の要素でよい
		movieTitle := item.Snippet.Title
		thumbnail := item.Snippet.Thumbnails
		channelName := item.Snippet.ChannelTitle
		videoID := item.Id.VideoId

		thumbnailURL := thumbnail.Default.Url
		movieURL := fmt.Sprintf("https://www.youtube.com/watch?v=%s", videoID)

		bubble := buildBubble(movieTitle, thumbnailURL, channelName, movieURL)
		bubbles = append(bubbles, bubble)
	}

	carousel := *model.NewCarousel("carousel", bubbles)
	flexMessage := *model.NewFlexMessage("flex", "本日の動画です。", &carousel)

	messageJSON, err := json.Marshal(flexMessage)
	if err != nil {
		fmt.Println("JSON marshal error:", err)
		return
	}
	fmt.Println(string(messageJSON))

	if err := sendMessage(messageJSON); err != nil {
		log.Fatal(err)
	}
}

func main() {
	lambda.Start(handler)
}

func buildBubble(movieTitle string, thumbnailURL string, channelTitle string, movieURL string) *model.Bubble {
	hero := buildHero(thumbnailURL, movieURL)
	body := buildBody(movieTitle, channelTitle, movieURL)
	footer := buildFooter()

	bubble := *model.NewBubble("bubble", hero, body, footer)
	return &bubble
}

func buildHero(thumbnailURL string, movieURL string) *model.Hero {
	action := *model.NewAction("uri", "", movieURL)
	hero := *model.NewHero("image", thumbnailURL, "full", "20:30", "cover", &action)
	return &hero
}

func buildBody(movieTitle string, channelTitle string, movieURL string) *model.Body {
	urlProperty := *model.NewContentBodyBlockUrlProperty("text", "URL", "#aaaaaa", "sm", 1)
	urlValueAction := *model.NewAction("url", "", movieURL)
	urlValue := *model.NewContentBodyBlockUrlValue("text", "動画はこちらをタップ", true, "#666666", "sm", 5, &urlValueAction)
	urlComponents := []*model.Content{&urlProperty, &urlValue}
	urlRootComponent := *model.NewContentBodyBlockUrlRoot("box", "baseline", "sm", urlComponents)

	channelProperty := *model.NewContentBodyBlockChannelPropertyValue("text", "ch", 1, true, "sm", "#aaaaaa")
	channelValue := *model.NewContentBodyBlockChannelPropertyValue("text", channelTitle, 5, true, "sm", "#aaaaaa")
	channelNameComponents := []*model.Content{&channelProperty, &channelValue}
	channelRootComponent := *model.NewContentBodyBlockChannelRoot("box", "baseline", channelNameComponents)

	movieInfo := []*model.Content{&channelRootComponent, &urlRootComponent}

	movieComponent := *model.NewContentMovieValue("box", "vertical", "lg", "sm", movieInfo)
	titleComponent := *model.NewContentMovieProperty("text", movieTitle, "bold", "xl", true)

	bodyComponents := []*model.Content{&titleComponent, &movieComponent}

	body := *model.NewBody("box", "vertical", bodyComponents)
	return &body
}

func buildFooter() *model.Footer {
	action := *model.NewAction("uri", "Youtubeトップへ", "https://youtube.com")
	footerContent := *model.NewFooterContent("button", "link", "sm", &action)
	content := []*model.FooterContent{&footerContent}
	footer := *model.NewFooter("box", "vertical", "sm", content, 0)
	return &footer
}

func sendMessage(messageJSON []byte) error {
	bot, err := linebot.New(os.Getenv("LineBotChannelSecret"), os.Getenv("LineBotChannelToken"))
	if err != nil {
		log.Fatal(err)
		return err
	}

	flexContainer, err := linebot.UnmarshalFlexMessageJSON(messageJSON)
	if err != nil {
		log.Fatal(err)
		return err
	}

	flexMessage := linebot.NewFlexMessage("Flex Message", flexContainer)
	if _, err := bot.BroadcastMessage(flexMessage).Do(); err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}
