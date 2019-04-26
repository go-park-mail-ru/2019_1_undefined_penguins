package database

import (
	"database/sql"

	"github.com/jackc/pgx"
)

func Query(sql string, args ...interface{}) (*sql.Rows, error) {
	if connection == nil {
		return nil, pgx.ErrDeadConn
	}

	tx, err := connection.Begin()
	if err != nil {
		return nil, err
	}

	rows, err := tx.Query(sql, args...)
	return rows, err
}

func Exec(sql string, args ...interface{}) (commandTag sql.Result, err error) {
	// if connection == nil {
	// 	return "", nil
	// }

	tx, err := connection.Begin()
	// if err != nil {
	// 	return "", err
	// }
	defer tx.Rollback()

	tag, err := tx.Exec(sql, args...)
	if err != nil {
		return tag, err
	}
	err = tx.Commit()
	if err != nil {
		return tag, err
	}

	return tag, nil
}
