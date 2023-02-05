package lastfm

import (
	"fmt"
	"net/url"
)

type AlbumMatches struct {
	Album []SimpleAlbum `json:"album"`
}

type SimpleAlbum struct {
	Name   string `json:"name"`
	Artist string `json:"artist"`
	Url    string `json:"url"`
}

type AlbumResult struct {
	Album Album `json:"album"`
}

type Album struct {
	Name          string `json:"name"`
	Artist        string `json:"artist"`
	Url           string `json:"url"`
	UserPlayCount int    `json:"userplaycount"`
}

func buildAlbumRequestUrl(album string, artist string, username string, apiKey string) string {
	params := url.Values{}
	params.Add("method", "album.getInfo")
	params.Add("album", album)
	params.Add("artist", artist)
	params.Add("username", username)
	params.Add("api_key", apiKey)
	params.Add("format", "json")
	return fmt.Sprintf("%s?%s", API_ROOT, params.Encode())
}
