package model

type RandMusic struct {
	MusicList []Music `json:"MusicList"`
	Index     int     `json:"Index"`
	Wpid      int     `json:"Wpid"`
}
