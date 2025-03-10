package handlers

import (
	"Music_Library/internal/database/postgres"
	"Music_Library/internal/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

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
