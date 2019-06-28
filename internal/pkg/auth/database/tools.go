package database

import (
	"auth/helpers"
	"auth/models"
	"database/sql"
)

func RowsToUsers(rows *sql.Rows) []*models.User {
	users := []*models.User{}
	defer rows.Close()
	for rows.Next() {
		entry := new(models.User)
		if err := rows.Scan(&entry.Login, &entry.Score, &entry.Email, &entry.Count); err == nil {
			helpers.LogMsg(err)
		}
		users = append(users, entry)
	}
	return users
}
