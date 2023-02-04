package lastfm

type TrackMatches struct {
	Track []Track `json:"track"`
}

type Track struct {
	Name   string `json:"name"`
	Artist string `json:"artist"`
	Url    string `json:"url"`
}
