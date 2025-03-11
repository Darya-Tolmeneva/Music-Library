package models

import "github.com/gin-gonic/gin"

type ErrorResponse struct {
	Error string `json:"error"`
}

func NewErrorResponse(c *gin.Context, status int, err string) {
	c.AbortWithStatusJSON(status, ErrorResponse{Error: err})
}
