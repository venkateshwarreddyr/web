package main

import (
	"database/sql"
	"log"
	"net/http"
	"strconv"
	"web/services"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	r := gin.Default()
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"} // Use "*" to allow any origin (not recommended for production)
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Authorization"}

	r.Use(cors.New(config)) // Use the CORS middleware

	db, err := sql.Open("postgres", "user=postgres dbname=postgres sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	r.GET("/users", func(c *gin.Context) {
		services.GetUsers(c, db)
	})

	r.GET("/posts", func(c *gin.Context) {
		services.GetPosts(c, db)
	})

	r.GET("/comments", func(c *gin.Context) {
		services.GetComments(c, db)
	})

	r.GET("/posts/:postId/comments", func(c *gin.Context) {
		postId := c.Param("postId")

		postIdInt, err := strconv.Atoi(postId)
		if err != nil {
			// Handle the error if the conversion fails
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid postId parameter"})
			return
		}
		services.GetCommentsById(c, db, int(postIdInt))
	})

	r.Run()
}
