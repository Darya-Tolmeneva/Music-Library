package handlers

import (
	"Music_Library/internal/database/postgres"
	"Music_Library/internal/models"
	"github.com/gin-gonic/gin"
	"log"
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
func GetLyric(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	lyric, err := postgres.GetLyric(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "lyric not found"})
		return
	}
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
func UpdateLyric(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	var updateLyric models.Lyric
	if err = c.ShouldBindJSON(&updateLyric); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	log.Printf("Received update data: %+v", updateLyric)
	lyric, err := postgres.UpdateLyric(uint(id), &updateLyric)
	if err != nil {
		if err.Error() == "record not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": "lyric not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}
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
func DeleteLyric(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	err = postgres.DeleteLyric(uint(id))
	if err != nil {
		if err.Error() == "record not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": "song not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}
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
func AddLyric(c *gin.Context) {
	var newLyric models.Lyric
	if err := c.ShouldBindJSON(&newLyric); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	log.Printf("Received update data: %+v", newLyric)
	err := postgres.AddLyric(&newLyric)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"lyric": newLyric})

}
