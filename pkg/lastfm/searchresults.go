package lastfm

import (
	"fmt"
	"net/url"
)

type SearchResultsWrapper struct {
	SearchResults SearchResults `json:"results"`
}

type SearchResults struct {
	TrackMatches TrackMatches `json:"trackmatches"`
}

func buildTrackSearchRequestUrl(title string, apiKey string) string {
	params := url.Values{}
	params.Add("method", "track.search")
	params.Add("track", title)
	params.Add("api_key", apiKey)
	params.Add("format", "json")
	return fmt.Sprintf("%s?%s", API_ROOT, params.Encode())
}
