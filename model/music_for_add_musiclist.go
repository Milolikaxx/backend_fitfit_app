package model

type MusicForAddSong struct {
	MusicList      []Music `json:"MusicList"`
	MusicForAdd    []Music `json:"Music"`
	IndexMusiclist int     `json:"IndexML"`
	IndexMusicAdd  int     `json:"IndexMusic"`
}
