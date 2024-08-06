package model

type MusicForAddSongOfPlaylist struct {
	PlaylistDetail []PlaylistDetail `json:"PlaylistDetail"`
	MusicForAdd    []Music          `json:"Music"`
	IndexPL        int              `json:"IndexPL"`
	IndexMusic     int              `json:"IndexMusic"`
}
