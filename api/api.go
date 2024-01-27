package api

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/ssebs/padpal-server/data"
)

func APIFoo() string {
	return "APIFoo"
}

// HandleAndServe will handle the routes and serve HTTP
func HandleAndServe(host string, port int) {
	hostPort := fmt.Sprintf("%s:%d", host, port)

	router := gin.Default()
	// Handlers
	router.GET("/", rootHandler)
	router.GET("/notes", notesHandler)
	// Run & log the server once it errors
	log.Fatal(router.Run(hostPort))
}

func rootHandler(c *gin.Context) {
	// get contents of ./REST-API.md and return
	d, err := os.ReadFile("./REST-API.md")
	if err != nil {
		c.Error(err)
		c.JSON(500, err)
		return
	}

	c.String(200, string(d))
}

func notesHandler(c *gin.Context) {
	c.JSON(200, data.GetSampleNotes())
}
