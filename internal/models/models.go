package models

import (
	"github.com/jinzhu/gorm"
)

type Song struct {
	gorm.Model
	ID          uint    `gorm:"primaryKey"`
	Group       string  `json:"group"`
	Title       string  `json:"title"`
	ReleaseDate string  `json:"release_date"`
	Link        string  `json:"link"`
	Lyrics      []Lyric `json:"lyrics" gorm:"foreignKey:SongID"`
}

type Lyric struct {
	gorm.Model
	ID          uint   `gorm:"primaryKey"`
	SongID      uint   `json:"song_id"`
	VerseNumber int    `json:"verse_number"`
	Text        string `json:"text"`
}
