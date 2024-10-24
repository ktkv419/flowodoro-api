package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type session struct {
	ID    string `json:"id"`
	Start string `json:"start"`
	End   string `json:"end"`
}

// Mockdata
var sessions = []session{
	{ID: "1", Start: "2024-10-24T17:29:44Z", End: "2024-10-24T17:33:58Z"},
	{ID: "2", Start: "2024-10-24T17:29:44Z", End: "2024-10-24T17:33:58Z"},
	{ID: "3", Start: "2024-10-24T17:29:44Z", End: "2024-10-24T17:33:58Z"},
}

func getSessions(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, sessions)
}

func getSessionByID(c *gin.Context) {
	id := c.Param("id")

	for _, a := range sessions {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "session not found"})
}

func postAlbums(c *gin.Context) {
	var newSession session

	if err := c.BindJSON(&newSession); err != nil {
		return
	}

	sessions = append(sessions, newSession)
	c.IndentedJSON(http.StatusCreated, newSession)
}

func main() {
	router := gin.Default()
	router.GET("/sessions", getSessions)
	router.POST("/sessions", postAlbums)
	router.GET("/sessions/:id", getSessionByID)
	router.Run("localhost:8080")
}
