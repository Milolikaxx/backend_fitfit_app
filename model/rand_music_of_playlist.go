package model

type RandMusicOfPlaylist struct {
	PlaylistDetail []PlaylistDetail `json:"PlaylistDetail"`
	Index          int              `json:"Index"`
	Wpid           int              `json:"Wpid"`
}
