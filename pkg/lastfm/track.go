package lastfm

type TrackMatches struct {
	Track []SimpleTrack `json:"track"`
}

type SimpleTrack struct {
	Name   string `json:"name"`
	Artist string `json:"artist"`
	Url    string `json:"url"`
}
