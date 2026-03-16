package services

import (
	"database/sql"
	"net/http"
	"web/repo"

	"github.com/gin-gonic/gin"
)

func GetPosts(c *gin.Context, db *sql.DB) {
	posts, err := repo.GetPosts(db)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, posts)
}
