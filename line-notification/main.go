package main

import (
	"fmt"
	"google.golang.org/api/youtube/v3"
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
}

func main() {
	lambda.Start(handler)
}
