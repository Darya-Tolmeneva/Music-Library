package router

import (
	"Music_Library/docs"
	"Music_Library/internal/transport/handlers"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log/slog"
)

func NewRouter(log *slog.Logger) *gin.Engine {
	router := gin.Default()
	docs.SwaggerInfo.BasePath = "/"
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	songRouter := router.Group("/songs")
	{
		songRouter.GET("/", func(c *gin.Context) {
			handlers.GetAllSongs(c, log)
		})
		songRouter.POST("/", func(c *gin.Context) {
			handlers.AddSong(c, log)
		})
		songRouter.GET("/:id", func(c *gin.Context) {
			handlers.GetSong(c, log)
		})
		songRouter.PUT("/:id", func(c *gin.Context) {
			handlers.UpdateSong(c, log)
		})
		songRouter.DELETE("/:id", func(c *gin.Context) {
			handlers.DeleteSong(c, log)
		})
	}

	lyricsRouter := router.Group("/lyrics")
	{
		lyricsRouter.GET("/:id", func(c *gin.Context) {
			handlers.GetLyric(c, log)
		})
		lyricsRouter.POST("/", func(c *gin.Context) {
			handlers.AddLyric(c, log)
		})
		lyricsRouter.PUT("/:id", func(c *gin.Context) {
			handlers.UpdateLyric(c, log)
		})
		lyricsRouter.DELETE("/:id", func(c *gin.Context) {
			handlers.DeleteLyric(c, log)
		})
	}

	return router
}
