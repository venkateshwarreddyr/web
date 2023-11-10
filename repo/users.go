package repo

import (
	"database/sql"
	"encoding/json"
	"log"
	"web/data"
)

func GetUsers(db *sql.DB) (*[]data.User, error) {
	rows, err := db.Query("SELECT * FROM users")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var users []data.User
	for rows.Next() {
		var user data.User
		var address string
		var company string

		if err := rows.Scan(
			&user.ID,
			&user.Name,
			&user.Username,
			&user.Email,
			&user.Phone,
			&user.Website,
			&address,
			&company,
		); err != nil {
			log.Fatal(err)
		}

		if err := json.Unmarshal([]byte(address), &user.Address); err != nil {
			log.Fatal(err)
		}

		if err := json.Unmarshal([]byte(company), &user.Company); err != nil {
			log.Fatal(err)
		}

		users = append(users, user)
	}
	return &users, nil
}
