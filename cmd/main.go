package main

import (
	"flag"
	"fmt"

	"tarellano.com/huh/cmd/huh"
)

func main() {
	var trackName, albumName, artistName, username, apiKey string
	flag.StringVar(&trackName, "track", "", "a track name")
	flag.StringVar(&albumName, "album", "", "an album name")
	flag.StringVar(&artistName, "artist", "", "an artist name")
	flag.StringVar(&username, "username", "", "your username")
	flag.StringVar(&apiKey, "api-key", "", "your api key")

	flag.Parse()

	if apiKey == "" {
		fmt.Println("api-key is required")
		return
	}

	if username == "" {
		fmt.Println("username is required")
		return
	}

	result, err := huh.GetStats(trackName, albumName, artistName, username, apiKey)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, r := range result {
		fmt.Printf("%s listened to %s %d times.\n", r.User.Name, r.Title, r.PlayCount)
	}
}
