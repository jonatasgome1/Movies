package main

import (
	"github.com/gin-gonic/gin"
)

type Movie struct {
	ID         int    `json:"id"`
	Title      string `json:"tilte"`
	Descripion string `json:"descripion"`
	Version    string `json:"version"`

	
}
var movies = []Movie{}

func CreateMovie(c *gin.Context) {
	var newMovie Movie
	if err := c.BindJSON(&newMovie); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})

	}

	newMovie.ID = len(movies) + 1
	movies = append(movies, newMovie)
	c.JSON(200., newMovie)

}
func GetMovies(c *gin.Context) {
	c.JSON(200, movies)
}

func main() {
	r := gin.Default()
	r.POST("/movie", CreateMovie)
	r.GET("/movies", GetMovies)
	r.Run(":8080")
}
