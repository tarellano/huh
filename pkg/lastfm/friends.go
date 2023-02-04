package lastfm

import (
	"fmt"
	"net/url"
)

type FriendsRequest struct {
	Friends Friends `json:"friends"`
}

type Friends struct {
	User []User `json:"user"`
}

func buildFriendsRequestUrl(
	user string,
	apiKey string,
) string {
	params := url.Values{}
	params.Add("method", "user.getfriends")
	params.Add("user", user)
	params.Add("api_key", apiKey)
	params.Add("format", "json")
	return fmt.Sprintf("%s?%s", API_ROOT, params.Encode())
}
