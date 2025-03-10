package handlers

import (
	"Music_Library/internal/database/postgres"
	"Music_Library/internal/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

func GetAllSongs(c *gin.Context) {
	group := c.Query("group")
	title := c.Query("title")
	releaseDate := c.Query("release_date")
	link := c.Query("link")

	offset, _ := strconv.Atoi(c.DefaultQuery("offset", "0"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))

	songs, total, err := postgres.GetAllSongs(group, title, releaseDate, link, offset, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": songs,
		"pagination": gin.H{
			"total":     total,
			"offset":    offset,
			"page_size": pageSize,
		},
	})
}

func GetSong(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	song, err := postgres.GetSong(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "song not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"song": song})
}

func AddSong(c *gin.Context) {
	var newSong models.Song
	if err := c.ShouldBindJSON(&newSong); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	log.Printf("Received update data: %+v", newSong)
	err := postgres.AddSong(&newSong)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"song": newSong})
}

func DeleteSong(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	err = postgres.DeleteSong(uint(id))
	if err != nil {
		if err.Error() == "record not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": "song not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}
	c.JSON(http.StatusOK, gin.H{"deleted_song_id": id})
}

func UpdateSong(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	var updateSong models.Song
	if err = c.ShouldBindJSON(&updateSong); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	log.Printf("Received update data: %+v", updateSong)
	song, err := postgres.UpdateSong(uint(id), &updateSong)
	if err != nil {
		if err.Error() == "record not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": "song not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}
	c.JSON(http.StatusOK, gin.H{"song": song})
}
