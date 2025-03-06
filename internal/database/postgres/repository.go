package postgres

import (
	"Music_Library/internal/models"
)

// GetSong возвращает песню по её ID.
func GetSong(id uint) (*models.Song, error) {
	var song models.Song
	result := db.First(&song, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &song, nil
}

// AddSong добавляет новую песню в базу данных.
func AddSong(song *models.Song) error {
	result := db.Create(song)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// GetAllSongs возвращает список песен с фильтрацией и пагинацией.
func GetAllSongs(group, title string, page, pageSize int) ([]models.Song, int64, error) {
	var songs []models.Song
	query := db.Model(&models.Song{})

	if group != "" {
		query = query.Where("group = ?", group)
	}
	if title != "" {
		query = query.Where("title = ?", title)
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
func UpdateSong(id uint, updatedSong *models.Song) error {
	result := db.Model(&models.Song{}).Where("id = ?", id).Updates(updatedSong)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// DeleteSong удаляет песню по её ID.
func DeleteSong(id uint) error {
	result := db.Delete(&models.Song{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
