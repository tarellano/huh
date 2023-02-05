package main

import (
	"flag"
	"fmt"
	"sort"

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

	if len(result) < 1 {
		fmt.Println("No results found...")
	}

	fmt.Printf("Have you heard %s (%s)?\nLet's see...\n", result[0].Title, result[0].Url)

	// Sort by most listened to to least
	sort.Slice(result, func(i, j int) bool {
		return result[i].PlayCount > result[j].PlayCount
	})

	for _, r := range result {
		fmt.Printf("%s listened to %s %d times.\n", r.User.Name, r.Title, r.PlayCount)
	}
}
