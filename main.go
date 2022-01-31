package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type album struct {
	ID string `json:"id"`
	Title string `json:"title"`
	Artist string `json:"artist"`
	Price float64 `json:"price"`
}

// albums are mock albums data.
var albums = []album{
	{ID: "1", Title: "Metallica (Black Album)", Artist: "Metallica", Price: 60.0},
	{ID: "2", Title: "Piece Sells... but Who's Buying?", Artist: "Megadeth", Price: 59.99},
	{ID: "3", Title: "Cowboys From Hell", Artist: "Pantera", Price: 75.0 },
	{ID: "4", Title: "Reign in Blood", Artist: "Slayer", Price: 50.99},
	{ID: "5", Title: "Arise", Artist: "Sepultura", Price: 45.0 },
}

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumByID)
	router.POST("/albums", postAlbums)

	router.Run("localhost: 8080")
}

// getAlbums responds with the list of all albums as JSON.
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

// postAlbums adds an album from JSON received in the request body.
func postAlbums(c *gin.Context) {
	var newAlbum album

	// Call BindJSON to bind the received JSON to newAlbum.
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	// Add the new album to the slice.
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

// getAlbumByID locates the album whose ID value matches the id
// parameter sent by the client, then return that album as a response.
func getAlbumByID(c *gin.Context) {
	id := c.Param("id")

	// Loop over the list of albums, looking for
	// an album whose ID value matches the parameter.
	for _, album := range albums {
		if album.ID == id {
			c.IndentedJSON(http.StatusOK, album)
			return 
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}





