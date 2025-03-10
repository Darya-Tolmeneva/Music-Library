package router

import (
	"Music_Library/docs"
	"Music_Library/internal/transport/handlers"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewRouter() *gin.Engine {
	router := gin.Default()
	docs.SwaggerInfo.BasePath = "/"
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	songRouter := router.Group("/songs")
	{
		songRouter.GET("/", handlers.GetAllSongs)
		songRouter.POST("/", handlers.AddSong)
		songRouter.GET("/:id", handlers.GetSong)
		songRouter.PUT("/:id", handlers.UpdateSong)
		songRouter.DELETE("/:id", handlers.DeleteSong)
	}
	lyricsRouter := router.Group("/lyrics")
	{
		lyricsRouter.GET("/:id", handlers.GetLyric)
		lyricsRouter.POST("/", handlers.AddLyric)
		lyricsRouter.PUT("/:id", handlers.UpdateLyric)
		lyricsRouter.DELETE("/:id", handlers.DeleteLyric)
	}

	return router
}
