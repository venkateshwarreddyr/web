package repo

import (
	"database/sql"
	"fmt"
	"log"
	"web/data"
)

func GetComments(db *sql.DB) (*[]data.Comment, error) {
	rows, err := db.Query("SELECT * FROM comments")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var comments []data.Comment
	for rows.Next() {
		var comment data.Comment

		if err := rows.Scan(
			&comment.Id,
			&comment.PostId,
			&comment.Name,
			&comment.Email,
			&comment.Body,
		); err != nil {
			log.Fatal(err)
		}

		comments = append(comments, comment)
	}
	return &comments, nil
}

func GetCommentsById(db *sql.DB, postId int) (*[]data.Comment, error) {
	rows, err := db.Query("SELECT * FROM comments where postid = $1", postId)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var comments []data.Comment
	for rows.Next() {
		var comment data.Comment

		if err := rows.Scan(
			&comment.Id,
			&comment.PostId,
			&comment.Name,
			&comment.Email,
			&comment.Body,
		); err != nil {
			log.Fatal(err)
		}

		comments = append(comments, comment)
	}
	fmt.Println(comments)
	return &comments, nil
}
