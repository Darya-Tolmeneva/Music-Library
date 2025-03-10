package handlers

import (
	"Music_Library/internal/database/postgres"
	"Music_Library/internal/models"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
	"strconv"
)

// GetLyric godoc
//
//	@Summary		Retrieve lyrics by ID
//	@Description	Fetches the lyrics of a song using its unique ID.
//	@Tags			Lyrics
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int					true	"Lyric ID"	Format(int)
//	@Success		200	{object}	models.Lyric		"Successfully retrieved lyrics"
//	@Failure		400	{object}	map[string]string	"Invalid ID format"
//	@Failure		404	{object}	map[string]string	"Lyrics not found"
//	@Router			/lyrics/{id} [get]
func GetLyric(c *gin.Context, logger *slog.Logger) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logger.Warn("Invalid ID format", "error", err, "id", c.Param("id"))
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	lyric, err := postgres.GetLyric(uint(id))
	if err != nil {
		logger.Error("Invalid input format", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "lyric not found"})
		return
	}
	logger.Info("Successfully fetched lyric", "id", id)
	c.JSON(http.StatusOK, gin.H{"lyric": lyric})

}

// UpdateLyric godoc
//
//	@Summary		Update lyrics information
//	@Description	Updates an existing lyric entry using its unique ID.
//	@Tags			Lyrics
//	@Accept			json
//	@Produce		json
//	@Param			id		path		int					true	"Lyric ID"	Format(int)
//	@Param			lyric	body		models.Lyric		true	"Updated lyric object"
//	@Success		200		{object}	models.Lyric		"Successfully updated lyrics"
//	@Failure		400		{object}	map[string]string	"Invalid input format"
//	@Failure		404		{object}	map[string]string	"Lyrics not found"
//	@Router			/lyrics/{id} [put]
func UpdateLyric(c *gin.Context, logger *slog.Logger) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logger.Warn("Invalid ID format", "error", err, "id", c.Param("id"))
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	var updateLyric models.Lyric
	if err = c.ShouldBindJSON(&updateLyric); err != nil {
		logger.Error("Invalid input format", "error", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	logger.Info("Received update data", "lyric", updateLyric)
	lyric, err := postgres.UpdateLyric(uint(id), &updateLyric)
	if err != nil {
		if err.Error() == "record not found" {
			logger.Warn("Lyric not found", "id", id)
			c.JSON(http.StatusNotFound, gin.H{"error": "lyric not found"})
		} else {
			logger.Error("Failed to update lyric", "error", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}
	logger.Info("Successfully updated lyric", "id", id, "lyric", lyric)
	c.JSON(http.StatusOK, gin.H{"lyric": lyric})

}

// DeleteLyric godoc
//
//	@Summary		Delete a lyric entry
//	@Description	Removes a lyric entry from the database using its ID.
//	@Tags			Lyrics
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int					true	"Lyric ID"	Format(int)
//	@Success		200	{object}	map[string]int		"Successfully deleted lyric ID"
//	@Failure		400	{object}	map[string]string	"Invalid lyric ID format"
//	@Failure		404	{object}	map[string]string	"Lyrics not found"
//	@Router			/lyrics/{id} [delete]
func DeleteLyric(c *gin.Context, logger *slog.Logger) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logger.Warn("Invalid song ID for deletion", "id", c.Param("id"), "error", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	err = postgres.DeleteLyric(uint(id))
	if err != nil {
		if err.Error() == "record not found" {
			logger.Warn("Song not found for deletion", "id", id)
			c.JSON(http.StatusNotFound, gin.H{"error": "song not found"})
		} else {
			logger.Error("Error deleting lyric", "id", id, "error", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}
	logger.Info("Successfully deleted lyric", "id", id)
	c.JSON(http.StatusOK, gin.H{"deleted_lyric_id": id})

}

// AddLyric godoc
//
//	@Summary		Create a new lyric entry
//	@Description	Adds a new lyric entry for a specific song in the database.
//	@Tags			Lyrics
//	@Accept			json
//	@Produce		json
//	@Param			lyric	body		models.Lyric		true	"Lyric object containing song ID and text"
//	@Success		201		{object}	models.Lyric		"Successfully created lyric entry"
//	@Failure		400		{object}	map[string]string	"Invalid input format"
//	@Router			/lyrics [post]
func AddLyric(c *gin.Context, logger *slog.Logger) {
	var newLyric models.Lyric
	if err := c.ShouldBindJSON(&newLyric); err != nil {
		logger.Error("Invalid input for new lyric", "error", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	logger.Info("Received new song", "song", newLyric)
	err := postgres.AddLyric(&newLyric)
	if err != nil {
		logger.Error("Error adding lyric", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	logger.Info("Successfully added new lyric", "song_id", newLyric.ID)
	c.JSON(http.StatusOK, gin.H{"lyric": newLyric})

}
