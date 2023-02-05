package lastfm

import (
	"fmt"
	"net/url"
)

type ArtistMatches struct {
	Artist []SimpleArtist `json:"artist"`
}

type SimpleArtist struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type ArtistResult struct {
	Artist Artist `json:"artist"`
}

type Artist struct {
	Name        string      `json:"name"`
	Url         string      `json:"url"`
	ArtistStats ArtistStats `json:"stats"`
}

type ArtistStats struct {
	Listeners     string `json:"listeners"`
	PlayCount     string `json:"playcount"`
	UserPlayCount string `json:"userplaycount"`
}

func buildArtistRequestUrl(artist string, username string, apiKey string) string {
	params := url.Values{}
	params.Add("method", "artist.getInfo")
	params.Add("artist", artist)
	params.Add("username", username)
	params.Add("api_key", apiKey)
	params.Add("format", "json")
	return fmt.Sprintf("%s?%s", API_ROOT, params.Encode())
}
