package repo

import (
	"database/sql"
	"log"
	"web/data"
)

func GetPosts(db *sql.DB) (*[]data.Post, error) {
	rows, err := db.Query("SELECT * FROM posts")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var posts []data.Post
	for rows.Next() {
		var post data.Post

		if err := rows.Scan(
			&post.Id,
			&post.UserId,
			&post.Title,
			&post.Body,
		); err != nil {
			log.Fatal(err)
		}

		posts = append(posts, post)
	}
	return &posts, nil
}
