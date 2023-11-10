package services

import (
	"database/sql"
	"fmt"
	"net/http"
	"web/repo"

	"github.com/gin-gonic/gin"
)

func GetPosts(c *gin.Context, db *sql.DB) {
	posts, err := repo.GetPosts(db)

	if err != nil {
		return
	}

	fmt.Println("hello12")
	c.JSON(http.StatusOK, posts)
}
