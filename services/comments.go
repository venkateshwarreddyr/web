package services

import (
	"database/sql"
	"fmt"
	"net/http"
	"web/repo"

	"github.com/gin-gonic/gin"
)

func GetComments(c *gin.Context, db *sql.DB) {
	comments, err := repo.GetComments(db)

	if err != nil {
		return
	}

	c.JSON(http.StatusOK, comments)
}

func GetCommentsById(c *gin.Context, db *sql.DB, postId int) {
	comments, err := repo.GetCommentsById(db, postId)
	fmt.Println(comments, postId, err)
	if err != nil {
		return
	}

	c.JSON(http.StatusOK, comments)
}
