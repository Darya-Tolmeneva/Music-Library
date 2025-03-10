package postgres

import (
	"Music_Library/internal/models"
	"errors"
)

// GetSong возвращает песню по её ID.
func GetSong(id uint) (*models.Song, error) {
	var song models.Song
	query := DB.Model(&models.Song{}).Preload("Lyrics")
	result := query.First(&song, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &song, nil
}

// AddSong добавляет новую песню в базу данных.
func AddSong(song *models.Song) error {
	result := DB.Create(song)
	if result.Error != nil {
		return result.Error
	}
	for lyric := range song.Lyrics {
		DB.Create(lyric)
	}
	return nil
}

// GetAllSongs возвращает список песен с фильтрацией и пагинацией.
func GetAllSongs(group, title, releaseDate, link string, page, pageSize int) ([]models.Song, int64, error) {
	var songs []models.Song
	query := DB.Model(&models.Song{}).Preload("Lyrics")

	if group != "" {
		query = query.Where(`"group" = ?`, group)
	}
	if title != "" {
		query = query.Where("title = ?", title)
	}
	if releaseDate != "" {
		query = query.Where("release_date = ?", releaseDate)
	}
	if link != "" {
		query = query.Where("link = ?", link)
	}

	var total int64
	query.Count(&total)

	query = query.Offset(page).Limit(pageSize)

	result := query.Find(&songs)
	if result.Error != nil {
		return nil, 0, result.Error
	}

	return songs, total, nil
}

// UpdateSong обновляет данные песни.
func UpdateSong(id uint, updatedSong *models.Song) (*models.Song, error) {
	result := DB.Model(&models.Song{}).Where("id = ?", id).Updates(updatedSong)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, errors.New("record not found")
	}
	song, err := GetSong(id)
	return song, err
}

// DeleteSong удаляет песню по её ID.
func DeleteSong(id uint) error {
	if err := DB.Where("song_id = ?", id).Delete(&models.Lyric{}).Error; err != nil {
		return errors.New("failed to delete lyrics")
	}
	result := DB.Delete(&models.Song{}, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("record not found")
	}
	return nil
}

// GetLyric возвращает куплет по его ID
func GetLyric(id uint) (*models.Lyric, error) {
	lyric := &models.Lyric{}
	result := DB.First(&lyric, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return lyric, nil
}

// AddLyric добавляет новый куплет
func AddLyric(lyric *models.Lyric) error {
	result := DB.Create(lyric)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// UpdateLyric обновляет куплет по id
func UpdateLyric(id uint, updateLyric *models.Lyric) (*models.Lyric, error) {
	result := DB.Model(&models.Lyric{}).Where("id = ?", id).Updates(updateLyric)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, errors.New("record not found")
	}
	lyric, err := GetLyric(id)
	return lyric, err

}

// DeleteLyric удаляет куплет по ID
func DeleteLyric(id uint) error {
	result := DB.Delete(&models.Lyric{}, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("record not found")
	}
	return nil
}
