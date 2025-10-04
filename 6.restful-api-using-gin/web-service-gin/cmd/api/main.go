package main

import (
	"example/web-service-gin/internal/album"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/albums", album.GetAlbums)
	router.POST("/albums", album.PostAlbums)

	router.GET("/albums/:id", album.GetAlbumByID)

	router.Run("localhost:8080")
}
