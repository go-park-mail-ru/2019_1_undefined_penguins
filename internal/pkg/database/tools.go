package database

import (
	"2019_1_undefined_penguins/internal/pkg/helpers"
	"2019_1_undefined_penguins/internal/pkg/models"
	"database/sql"
)

func RowsToUsers(rows *sql.Rows) []models.User {
	users := []models.User{}
	for rows.Next() {
		entry := models.User{}
		if err := rows.Scan(&entry.Login, &entry.Score, &entry.Email); err == nil {
			helpers.LogMsg(err)
		}
		users = append(users, entry)
	}
	return users
}
