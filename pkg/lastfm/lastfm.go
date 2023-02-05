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

func (c *Client) SearchTracks(
	title string,
) (SearchResults, error) {
	url := buildTrackSearchRequestUrl(title, c.ApiKey)

	resp, err := http.Get(url)
	if err != nil {
		log.Println(err)
		return SearchResults{}, err
	}
	defer resp.Body.Close()

	var srw SearchResultsWrapper
	err = json.NewDecoder(resp.Body).Decode(&srw)
	if err != nil {
		log.Println(err)
		return SearchResults{}, err
	}

	return srw.SearchResults, nil
}

func (c *Client) GetTrack(
	title string,
	artist string,
	username string,
) (Track, error) {
	url := buildTrackRequestUrl(title, artist, username, c.ApiKey)

	resp, err := http.Get(url)
	if err != nil {
		log.Println(err)
		return Track{}, err
	}
	defer resp.Body.Close()

	var tr TrackResult
	err = json.NewDecoder(resp.Body).Decode(&tr)
	if err != nil {
		log.Println(err)
		return Track{}, err
	}

	return tr.Track, nil
}

func (c *Client) SearchAlbums(
	album string,
) (SearchResults, error) {
	url := buildAlbumSearchRequestUrl(album, c.ApiKey)

	resp, err := http.Get(url)
	if err != nil {
		log.Println(err)
		return SearchResults{}, err
	}
	defer resp.Body.Close()

	var srw SearchResultsWrapper
	err = json.NewDecoder(resp.Body).Decode(&srw)
	if err != nil {
		log.Println(err)
		return SearchResults{}, err
	}

	return srw.SearchResults, nil
}

func (c *Client) GetAlbum(
	album string,
	artist string,
	username string,
) (Album, error) {
	url := buildAlbumRequestUrl(album, artist, username, c.ApiKey)

	resp, err := http.Get(url)
	if err != nil {
		log.Println(err)
		return Album{}, err
	}
	defer resp.Body.Close()

	var ar AlbumResult
	err = json.NewDecoder(resp.Body).Decode(&ar)
	if err != nil {
		log.Println(err)
		return Album{}, err
	}

	return ar.Album, nil
}

func (c *Client) SearchArtists(
	title string,
) (SearchResults, error) {
	url := buildArtistSearchRequestUrl(title, c.ApiKey)

	resp, err := http.Get(url)
	if err != nil {
		log.Println(err)
		return SearchResults{}, err
	}
	defer resp.Body.Close()

	var srw SearchResultsWrapper
	err = json.NewDecoder(resp.Body).Decode(&srw)
	if err != nil {
		log.Println(err)
		return SearchResults{}, err
	}

	return srw.SearchResults, nil
}

func (c *Client) GetArtist(
	artist string,
	username string,
) (Artist, error) {
	url := buildArtistRequestUrl(artist, username, c.ApiKey)

	resp, err := http.Get(url)
	if err != nil {
		log.Println(err)
		return Artist{}, err
	}
	defer resp.Body.Close()

	var ar ArtistResult
	err = json.NewDecoder(resp.Body).Decode(&ar)
	if err != nil {
		log.Println(err)
		return Artist{}, err
	}

	return ar.Artist, nil
}
