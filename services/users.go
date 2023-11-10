package services

import (
	"database/sql"
	"fmt"
	"net/http"
	"web/repo"

	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context, db *sql.DB) {
	users, err := repo.GetUsers(db)

	if err != nil {
		return
	}

	fmt.Println("hello12")
	c.JSON(http.StatusOK, *users)
}
