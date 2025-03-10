package models

// Song represents a song
// @Description Song model
type Song struct {
	ID          uint    `gorm:"primaryKey"`
	Group       string  `json:"group"`
	Title       string  `json:"title"`
	ReleaseDate string  `json:"release_date"`
	Link        string  `json:"link"`
	Lyrics      []Lyric `json:"lyrics" gorm:"foreignKey:SongID"`
}

// Lyric represents a song lyric
// @Description Song lyrics model
type Lyric struct {
	ID          uint   `gorm:"primaryKey"`
	SongID      uint   `json:"song_id"`
	VerseNumber int    `json:"verse_number"`
	Text        string `json:"text"`
}
