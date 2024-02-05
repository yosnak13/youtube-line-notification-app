package main

import (
	"flag"
	"fmt"
	"google.golang.org/api/youtube/v3"
	"log"
	"net/http"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
	"google.golang.org/api/googleapi/transport"
)

func handler() {
	var channelID string
	flag.StringVar(&channelID, "channel", "UCPVr7clenPjpD7WNsSI3UBQ", "YouTube Channel ID") // レトルトさんのチャンネルID
	flag.Parse()

	if channelID == "" {
		log.Fatal("Please provide a YouTube Channel ID using -channel flag")
	}

	apiKey := os.Getenv("API_KEY")
	client := &http.Client{
		Transport: &transport.APIKey{Key: apiKey},
	}

	service, err := youtube.New(client)
	if err != nil {
		log.Fatalf("Error creating new YouTube client: %v", err)
	}

	call := service.Search.List([]string{"snippet"}).ChannelId(channelID).Type("video").Order("date").MaxResults(5)
	response, err := call.Do()
	if err != nil {
		log.Fatalf("Error making search API call: %v", err)
	}

	for _, item := range response.Items {
		videoTitle := item.Snippet.Title
		thumbnails := item.Snippet.Thumbnails
		channelTitle := item.Snippet.ChannelTitle

		fmt.Printf("Title: %s\n", videoTitle)
		fmt.Printf("Thumbnail URL: %s\n", thumbnails.Default.Url)
		fmt.Printf("Channel Name: %s\n", channelTitle)
	}
}

func main() {
	lambda.Start(handler)
}
