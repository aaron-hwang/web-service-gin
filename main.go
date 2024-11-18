package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type album struct {
	ID          string  `json:"id"`
	Title       string  `json:"title"`
	Artist      string  `json:"artist"`
	Price       float64 `json:"price"`
	ReleaseYear int     `json:"releaseyear"`
}

var exampleAlbums = []album{
	{"1", "Test", "Brian", 12.0, 1200},
}

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, exampleAlbums)
}

func getAlbumById(c *gin.Context) {
	id := c.Param("id")

	for _, a := range exampleAlbums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func postAlbums(c *gin.Context) {
	var newAlbum album
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	exampleAlbums = append(exampleAlbums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumById)
	router.POST("/albums", postAlbums)

	router.Run("localhost:8080")
}
