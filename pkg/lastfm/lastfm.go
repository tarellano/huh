package lastfm

import (
	"encoding/json"
	"log"
	"net/http"
)

const API_ROOT = "https://ws.audioscrobbler.com/2.0/"

type Client struct {
	ApiKey string
}

// user.getFriends provides a list of friends
func (c *Client) GetFriends(
	user string,
) (Friends, error) {
	url := buildFriendsRequestUrl(user, c.ApiKey)

	resp, err := http.Get(url)
	if err != nil {
		log.Println(err)
		return Friends{}, err
	}
	defer resp.Body.Close()

	var f FriendsRequest
	err = json.NewDecoder(resp.Body).Decode(&f)
	if err != nil {
		log.Println(err)
		return Friends{}, err
	}

	return f.Friends, nil
}