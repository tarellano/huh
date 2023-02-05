package huh

import (
	"errors"
	"fmt"
	"strconv"

	"tarellano.com/huh/pkg/lastfm"
)

type UserStats struct {
	User      lastfm.User
	Title     string
	Url       string
	PlayCount int
}

func GetStats(
	trackName string,
	albumName string,
	artistName string,
	username string,
	apiKey string,
) ([]UserStats, error) {
	client := &lastfm.Client{
		ApiKey: apiKey,
	}

	friends, err := client.GetFriends(username)
	if err != nil {
		fmt.Println(err.Error())
	}

	if len(trackName) > 0 {
		track := findTrack(trackName, *client)
		return getTrackStats(track, friends.User, client)
	}

	if len(albumName) > 0 {
		album := findAlbum(albumName, *client)
		return getAlbumStats(album, friends.User, client)
	}

	if len(artistName) > 0 {
		artist := findArtist(artistName, *client)
		return getArtistStats(artist, friends.User, client)
	}

	return []UserStats{}, errors.New("could not determine music entity")
}

func findTrack(
	trackName string,
	client lastfm.Client,
) lastfm.SimpleTrack {
	results, err := client.SearchTracks(trackName)
	if err != nil {
		fmt.Println(err.Error())
	}

	if len(results.TrackMatches.Track) > 0 {
		return results.TrackMatches.Track[0]
	}
	return lastfm.SimpleTrack{}
}

func getTrackStats(
	track lastfm.SimpleTrack,
	friends []lastfm.User,
	client *lastfm.Client,
) ([]UserStats, error) {
	stats := make([]UserStats, len(friends))
	for i, friend := range friends {
		trackInfo, err := client.GetTrack(track.Name, track.Artist, friend.Name)
		if err != nil {
			fmt.Println(err.Error())
			return []UserStats{}, err
		}
		playCount, err := strconv.Atoi(trackInfo.UserPlayCount)
		if err != nil {
			fmt.Println(err.Error())
			return []UserStats{}, err
		}
		stats[i] = UserStats{
			friend,
			track.Name,
			track.Url,
			playCount,
		}
	}

	return stats, nil
}

func findAlbum(
	album string,
	client lastfm.Client,
) lastfm.SimpleAlbum {
	results, err := client.SearchAlbums(album)
	if err != nil {
		fmt.Println(err.Error())
	}

	if len(results.AlbumMatches.Album) > 0 {
		return results.AlbumMatches.Album[0]
	}
	return lastfm.SimpleAlbum{}
}

func getAlbumStats(
	album lastfm.SimpleAlbum,
	friends []lastfm.User,
	client *lastfm.Client,
) ([]UserStats, error) {
	stats := make([]UserStats, len(friends))
	for i, friend := range friends {
		albumInfo, err := client.GetAlbum(album.Name, album.Artist, friend.Name)
		if err != nil {
			fmt.Println(err.Error())
			return []UserStats{}, err
		}
		stats[i] = UserStats{
			friend,
			album.Name,
			album.Url,
			albumInfo.UserPlayCount,
		}
	}

	return stats, nil
}

func findArtist(
	artist string,
	client lastfm.Client,
) lastfm.SimpleArtist {
	results, err := client.SearchArtists(artist)
	if err != nil {
		fmt.Println(err.Error())
	}

	if len(results.ArtistMatches.Artist) > 0 {
		return results.ArtistMatches.Artist[0]
	}
	return lastfm.SimpleArtist{}
}

func getArtistStats(
	artist lastfm.SimpleArtist,
	friends []lastfm.User,
	client *lastfm.Client,
) ([]UserStats, error) {
	stats := make([]UserStats, len(friends))
	for i, friend := range friends {
		artistInfo, err := client.GetArtist(artist.Name, friend.Name)
		if err != nil {
			fmt.Println(err.Error())
			return []UserStats{}, err
		}
		playCount, err := strconv.Atoi(artistInfo.ArtistStats.UserPlayCount)
		if err != nil {
			fmt.Println(err.Error())
			return []UserStats{}, err
		}
		stats[i] = UserStats{
			friend,
			artist.Name,
			artist.Url,
			playCount,
		}
	}

	return stats, nil
}
