package services

import (
	"database/sql"
	"net/http"
	"web/repo"

	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context, db *sql.DB) {
	users, err := repo.GetUsers(db)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, *users)
}
