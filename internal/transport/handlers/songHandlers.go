package handlers

import (
	"Music_Library/internal/database/postgres"
	"Music_Library/internal/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetAllSongs(c *gin.Context) {
	group := c.Query("group")
	song := c.Query("song")
	offset, _ := strconv.Atoi(c.DefaultQuery("offset", "0"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))

	songs, total, err := postgres.GetAllSongs(group, song, offset, pageSize)
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
	id, _ := strconv.Atoi(c.Param("id"))
	song, err := postgres.GetSong(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Song not found"})
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

	err := postgres.AddSong(&newSong)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"song": newSong})
}

func DeleteSong(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	err := postgres.DeleteSong(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"song_id": id})

}
func UpdateSong(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var updateSong models.Song
	if err := c.ShouldBindJSON(&updateSong); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := postgres.UpdateSong(uint(id), &updateSong)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"song": updateSong})
}
