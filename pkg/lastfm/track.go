package lastfm

import (
	"fmt"
	"net/url"
)

type TrackMatches struct {
	Track []SimpleTrack `json:"track"`
}

type SimpleTrack struct {
	Name   string `json:"name"`
	Artist string `json:"artist"`
	Url    string `json:"url"`
}

type TrackResult struct {
	Track Track `json:"track"`
}

type Track struct {
	Name          string `json:"name"`
	Url           string `json:"url"`
	UserPlayCount string `json:"userplaycount"`
}

func buildTrackRequestUrl(track string, artist string, username string, apiKey string) string {
	params := url.Values{}
	params.Add("method", "track.getInfo")
	params.Add("track", track)
	params.Add("artist", artist)
	params.Add("username", username)
	params.Add("api_key", apiKey)
	params.Add("format", "json")
	return fmt.Sprintf("%s?%s", API_ROOT, params.Encode())
}
