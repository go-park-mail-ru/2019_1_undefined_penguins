package database

import (
	"2019_1_undefined_penguins/internal/pkg/models"
	"fmt"

	"github.com/jackc/pgx"
)

func RowsToUsers(rows *pgx.Rows) []models.User {
	users := []models.User{}
	for rows.Next() {
		entry := models.User{}
		if err := rows.Scan(&entry.Login, &entry.Score, &entry.Email); err == nil {
					fmt.Println(err)
				}
		users = append(users, entry)
	}
	return users
}
