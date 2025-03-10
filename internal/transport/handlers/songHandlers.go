package handlers

import (
	"Music_Library/internal/database/postgres"
	"Music_Library/internal/models"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
	"strconv"
)

// GetAllSongs godoc
//
//	@Summary		Get all songs
//	@Description	Fetch a list of all songs, with optional filters for group, title, release date, and link. Pagination is supported with offset and page_size parameters.
//	@Tags			songs
//	@Accept			json
//	@Produce		json
//	@Param			group			query		string					false	"Filter songs by group name"
//	@Param			title			query		string					false	"Filter songs by title"
//	@Param			release_date	query		string					false	"Filter songs by release date (YYYY-MM-DD)"
//	@Param			link			query		string					false	"Filter songs by associated link"
//	@Param			offset			query		int						false	"Pagination offset, starting from 0 (default: 0)"
//	@Param			page_size		query		int						false	"Number of items per page (default: 10)"
//	@Success		200				{object}	map[string]interface{}	"List of songs with pagination metadata"
//	@Failure		500				{object}	map[string]string		"Internal server error"
//	@Router			/songs [get]
func GetAllSongs(c *gin.Context, logger *slog.Logger) {
	group := c.Query("group")
	title := c.Query("title")
	releaseDate := c.Query("release_date")
	link := c.Query("link")

	offset, _ := strconv.Atoi(c.DefaultQuery("offset", "0"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))

	songs, total, err := postgres.GetAllSongs(group, title, releaseDate, link, offset, pageSize)
	if err != nil {
		logger.Error("Error fetching songs", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	logger.Info("Successfully fetched songs", "total", total)
	c.JSON(http.StatusOK, gin.H{
		"data": songs,
		"pagination": gin.H{
			"total":     total,
			"offset":    offset,
			"page_size": pageSize,
		},
	})
}

// GetSong godoc
//
//	@Summary		Get song by ID
//	@Description	Fetch details of a specific song by its ID
//	@Tags			songs
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int					true	"ID of the song"
//	@Success		200	{object}	models.Song			"Song details"
//	@Failure		400	{object}	map[string]string	"Invalid song ID"
//	@Failure		404	{object}	map[string]string	"Song not found"
//	@Router			/songs/{id} [get]
func GetSong(c *gin.Context, logger *slog.Logger) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logger.Warn("Invalid song ID", "id", c.Param("id"), "error", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	song, err := postgres.GetSong(uint(id))
	if err != nil {
		logger.Error("Error fetching song", "id", id, "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "song not found"})
		return
	}
	logger.Info("Successfully fetched song", "id", id)
	c.JSON(http.StatusOK, gin.H{"song": song})
}

// AddSong godoc
//
//	@Summary		Add a new song
//	@Description	Add a new song to the database
//	@Tags			songs
//	@Accept			json
//	@Produce		json
//	@Param			song	body		models.Song	true	"Song object"
//	@Success		201		{object}	models.Song
//	@Failure		400		{object}	map[string]string	"Invalid input"
//	@Router			/songs [post]
func AddSong(c *gin.Context, logger *slog.Logger) {
	var newSong models.Song
	if err := c.ShouldBindJSON(&newSong); err != nil {
		logger.Error("Invalid input for new song", "error", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	logger.Info("Received new song", "song", newSong)
	err := postgres.AddSong(&newSong)
	if err != nil {
		logger.Error("Error adding song", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	logger.Info("Successfully added new song", "song_id", newSong.ID)
	c.JSON(http.StatusOK, gin.H{"song": newSong})
}

// DeleteSong godoc
//
//	@Summary		Delete a song
//	@Description	Delete a song from the database using its ID
//	@Tags			songs
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int					true	"ID of the song to be deleted"
//	@Success		200	{object}	map[string]int		"ID of the deleted song"
//	@Failure		400	{object}	map[string]string	"Invalid song ID"
//	@Failure		404	{object}	map[string]string	"Song not found"
//	@Router			/songs/{id} [delete]
func DeleteSong(c *gin.Context, logger *slog.Logger) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logger.Warn("Invalid song ID for deletion", "id", c.Param("id"), "error", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	err = postgres.DeleteSong(uint(id))
	if err != nil {
		if err.Error() == "record not found" {
			logger.Warn("Song not found for deletion", "id", id)
			c.JSON(http.StatusNotFound, gin.H{"error": "song not found"})
		} else {
			logger.Error("Error deleting song", "id", id, "error", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}
	logger.Info("Successfully deleted song", "id", id)
	c.JSON(http.StatusOK, gin.H{"deleted_song_id": id})
}

// UpdateSong godoc
//
//	@Summary		Update an existing song
//	@Description	Update song details by its ID, such as title, group, or release date
//	@Tags			songs
//	@Accept			json
//	@Produce		json
//	@Param			id		path		int					true	"ID of the song to be updated"
//	@Param			song	body		models.Song			true	"Updated song details"
//	@Success		200		{object}	models.Song			"Updated song details"
//	@Failure		400		{object}	map[string]string	"Invalid input data"
//	@Failure		404		{object}	map[string]string	"Song not found"
//	@Router			/songs/{id} [put]
func UpdateSong(c *gin.Context, logger *slog.Logger) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logger.Warn("Invalid song ID for update", "id", c.Param("id"), "error", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	var updateSong models.Song
	if err = c.ShouldBindJSON(&updateSong); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	logger.Info("Received update data", "song", updateSong)

	song, err := postgres.UpdateSong(uint(id), &updateSong)
	if err != nil {
		if err.Error() == "record not found" {
			logger.Warn("Song not found for update", "id", id)
			c.JSON(http.StatusNotFound, gin.H{"error": "song not found"})
		} else {
			logger.Error("Error updating song", "id", id, "error", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}
	logger.Info("Successfully updated song", "id", id, "song", song)
	c.JSON(http.StatusOK, gin.H{"song": song})
}
