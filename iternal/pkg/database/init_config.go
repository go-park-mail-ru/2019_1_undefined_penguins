package database

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"github.com/jackc/pgx"
)

type Database struct {
	Host     string
	Port     uint16
	Database string
	User     string
}

var connectionConfig pgx.ConnConfig

// var Connection *pgx.ConnPool
var connectionPoolConfig = pgx.ConnPoolConfig{
	MaxConnections: 8,
}

func InitConfig() {

	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	file, err := os.Open(dir + "/configs/database.json")
	if err != nil {
		//remove panic
		panic(err)
	}
	body, _ := ioutil.ReadAll(file)
	err = json.Unmarshal(body, &connectionConfig)
	connectionPoolConfig.ConnConfig = connectionConfig

}

func Connect() *pgx.ConnPool {
	connectionPool, err := pgx.NewConnPool(connectionPoolConfig)
	if err != nil {
		//remove FATAL, panic, use fmt 
		log.Fatal(err)
	}
	return connectionPool
}

func StartTransaction() *pgx.Tx {
	connection := Connect()
	if transaction, err := connection.Begin(); err != nil {
		log.Fatal(err)
		//return nil
		return transaction
	} else {

		return transaction
	}
}

func CommitTransaction(transaction *pgx.Tx) error {
	if err := transaction.Commit(); err != nil {
		transaction.Rollback()
		return err
	}
	return nil
}
