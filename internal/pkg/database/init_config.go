package database

import (
	"2019_1_undefined_penguins/internal/pkg/helpers"
	h "2019_1_undefined_penguins/internal/pkg/helpers"
	sq "database/sql"
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/jackc/pgx"
	_ "github.com/lib/pq"
)

var connection *sq.DB = nil

var connectionConfig pgx.ConnConfig
var connectionPoolConfig = pgx.ConnPoolConfig{
	MaxConnections: 8,
}

func initConfig() error {
	dir, err := os.Getwd()
	if err != nil {
		h.LogMsg("Getting directory error: ", err)
		return err

	}
	file, err := os.Open(dir + "/configs/database.json")
	if err != nil {
		h.LogMsg("Init parse DB error: ", err)
		return err
	}
	body, _ := ioutil.ReadAll(file)
	err = json.Unmarshal(body, &connectionConfig)
	if err != nil {
		h.LogMsg("Init parse DB error: ", err)
		return err
	}
	psqlURI := "postgresql://" + connectionConfig.User
	if len(connectionConfig.Password) > 0 {
		psqlURI += ":" + connectionConfig.Password
	}
	psqlURI += "@" + connectionConfig.Host + ":" + connectionConfig.Host + "/" + connectionConfig.Database + "?sslmode=disable"
	connection, err = sq.Open("postgres", psqlURI)
	if err != nil {
		helpers.LogMsg("Can't connect to db: ", err)
		return err
	}
	return nil
}

func SetMock(databaseMock *sq.DB) {
	connection = databaseMock
}

func Connect() error {
	if connection != nil {
		return nil
	}
	err := initConfig()
	if err != nil {
		return err
	}
	return nil
}

func Disconnect() {
	if connection != nil {
		connection.Close()
		connection = nil
	}
}
