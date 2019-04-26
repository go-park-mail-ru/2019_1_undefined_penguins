package database

import (
	"2019_1_undefined_penguins/internal/pkg/helpers"
	sq "database/sql"

	"github.com/jackc/pgx"
	_ "github.com/lib/pq"
)

var connection *sq.DB = nil

var connectionConfig pgx.ConnConfig
var connectionPoolConfig = pgx.ConnPoolConfig{
	MaxConnections: 8,
}

func initConfig() error {

	// dir, err := os.Getwd()
	// if err != nil {
	// 	h.LogMsg("Getting directory error: ", err)
	// 	return err

	// }

	// file, err := os.Open(dir + "/configs/database.json")
	// if err != nil {
	// 	h.LogMsg("Init parse DB error: ", err)
	// 	return err
	// }

	// body, _ := ioutil.ReadAll(file)
	// err = json.Unmarshal(body, &connectionConfig)
	// if err != nil {
	// 	h.LogMsg("Init parse DB error: ", err)
	// 	return err
	// }

	const psqlURI = "postgresql://iamfrommoscow@localhost:5432/penguins?sslmode=disable"
	// connectionConfig, err := pgx.ParseURI(psqlURI)
	// if err != nil {
	// 	h.LogMsg("Parsing error: ", err)
	// 	return err

	// }
	var err error
	connection, err = sq.Open("postgres", psqlURI)
	if err != nil {
		helpers.LogMsg("Can't connect to db: ", err)
		return err
	}
	// connectionPoolConfig.ConnConfig = connectionConfig
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
	// connection, err = pgx.NewConnPool(connectionPoolConfig)
	// if err != nil {
	// 	h.LogMsg("Connect DB error: " + err.Error())
	// 	return err
	// }
	return nil
}

func Disconnect() {
	if connection != nil {
		connection.Close()
		connection = nil
	}
}
