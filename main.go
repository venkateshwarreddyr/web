package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"go-posts-api/services"

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

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		getEnv("DB_HOST", "localhost"),
		getEnv("DB_PORT", "5432"),
		getEnv("DB_USER", "postgres"),
		getEnv("DB_PASSWORD", ""),
		getEnv("DB_NAME", "postgres"),
		getEnv("DB_SSLMODE", "disable"),
	)
	db, err := sql.Open("postgres", dsn)
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

func getEnv(key, fallback string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return fallback
}
