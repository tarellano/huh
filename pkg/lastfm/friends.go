package lastfm

import "fmt"

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
	return fmt.Sprintf("%s?method=user.getfriends&user=%s&api_key=%s&format=json", API_ROOT, user, apiKey)
}
