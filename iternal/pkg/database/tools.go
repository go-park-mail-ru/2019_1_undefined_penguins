package database

import (
	"2019_1_undefined_penguins/iternal/pkg/models"

	"github.com/jackc/pgx"
)

func RowsToUsers(rows *pgx.Rows) []models.User {
	users := []models.User{}
	for rows.Next() {
		var user models.User
		if err := rows.Scan(
			&user.Login,
			&user.Email,
			&user.Score); err != nil {
			return users
		}

		users = append(users, user)
	}
	// for rows.Next() {
	// 	entry := models.User{}
	// 	_ = rows.Scan(&entry.Login, &entry.Name, &entry.Score, &entry.Email, &entry.HashPassword)
	// 	users = append(users, entry)
	// }
	return users
}
