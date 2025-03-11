package models

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	ID      int    `json:"id"`
	Message string `json:"message"`
}

func NewResponse(c *gin.Context, id int, message string) {
	c.JSON(http.StatusOK, Response{
		ID:      id,
		Message: message,
	})
}
