package database

import (
	"encoding/json"
	"io/ioutil"
	"os"

	h "github.com/go-park-mail-ru/2019_1_undefined_penguins/iternal/pkg/helpers"
	"github.com/jackc/pgx"
)

var connection *pgx.ConnPool = nil

var connectionConfig pgx.ConnConfig
var connectionPoolConfig = pgx.ConnPoolConfig{
	MaxConnections: 8,
}
//TODO check connect
func initConfig() error {
	dir, err := os.Getwd()
	if err != nil {
		h.LogMsg("Init DB error: " + err.Error())
		return err
	}

	file, err := os.Open(dir + "/configs/database.json")
	if err != nil {
		h.LogMsg("Open DB error: " + err.Error())
		return err
	}

	body, _ := ioutil.ReadAll(file)
	err = json.Unmarshal(body, &connectionConfig)
	if err != nil {
		h.LogMsg("Init parse DB error: " + err.Error())
		return err
	}

	connectionPoolConfig.ConnConfig = connectionConfig
	return nil
}

func Connect() (error) {
	if connection != nil {
		return nil
	}
	err := initConfig()
	if err != nil {
		return err
	}
	connection, err = pgx.NewConnPool(connectionPoolConfig)
	if err != nil {
		h.LogMsg("Connect DB error: " + err.Error())
		return err
	}
	return nil
}

func Disconnect()  {
	if connection != nil {
		connection.Close()
		connection = nil
	}
}

