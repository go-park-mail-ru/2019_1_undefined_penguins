package database

import (
	"database/sql"

	"github.com/jackc/pgx"
)

func Query(sql string, args ...interface{}) (*sql.Rows, error) {
	if connection == nil {
		return nil, pgx.ErrDeadConn
	}
	rows, err := connection.Query(sql, args...)
	return rows, err
}

func Exec(sql string, args ...interface{}) (commandTag sql.Result, err error) {

	tx, err := connection.Begin()
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
