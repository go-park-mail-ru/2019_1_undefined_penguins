package database

import (
	"encoding/json"
	"io/ioutil"
	"strings"

	"os"

	h "2019_1_undefined_penguins/internal/pkg/helpers"

	"github.com/jackc/pgx"
)

var connection *pgx.ConnPool = nil

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
		dirRep := strings.Replace(dir, "/internal/pkg/controllers", "", -1)
		file, err = os.Open(dirRep + "/configs/testbase.json")
		if err != nil {
			dirRep = strings.Replace(dir, "/internal/pkg/database", "", -1)
			file, err = os.Open(dirRep + "/configs/testbase.json")
			if err != nil {
				h.LogMsg("Open directory error: ", err)
				return err
			}
		}
	}

	body, _ := ioutil.ReadAll(file)
	err = json.Unmarshal(body, &connectionConfig)
	if err != nil {
		h.LogMsg("Init parse DB error: ", err)
		return err
	}

	connectionPoolConfig.ConnConfig = connectionConfig
	return nil
}

func Connect() error {
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

func Disconnect() {
	if connection != nil {
		connection.Close()
		connection = nil
	}
}
