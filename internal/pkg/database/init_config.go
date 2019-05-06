package database

import (
	"2019_1_undefined_penguins/internal/pkg/helpers"
	sq "database/sql"
	"strconv"

	"github.com/jackc/pgx"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

var connection *sq.DB = nil

var connectionConfig pgx.ConnConfig
var connectionPoolConfig = pgx.ConnPoolConfig{
	MaxConnections: 8,
}

var ImagesAddress string

func initConfig() error {
	viper.AddConfigPath("./configs")
	viper.SetConfigName("config")
	if err := viper.ReadInConfig(); err != nil {
		helpers.LogMsg("Can't find db config: ", err)
		return err
	}
	connectionConfig = pgx.ConnConfig{
		Host:     viper.GetString("db.host"),
		Port:     uint16(viper.GetInt("db.port")),
		Database: viper.GetString("db.database"),
		User:     viper.GetString("db.user"),
	}
	psqlURI := "postgresql://" + connectionConfig.User
	if len(connectionConfig.Password) > 0 {
		psqlURI += ":" + connectionConfig.Password
	}
	psqlURI += "@" + connectionConfig.Host + ":" + strconv.Itoa(int(connectionConfig.Port)) + "/" + connectionConfig.Database + "?sslmode=disable"
	var err error
	connection, err = sq.Open("postgres", psqlURI)
	if err != nil {
		helpers.LogMsg("Can't connect to db: ", err)
		return err
	}
	viper.SetConfigName("fileserver")
	if err := viper.ReadInConfig(); err != nil {
		helpers.LogMsg("Can't find images address: ", err)
		return err
	}
	ImagesAddress = viper.GetString("address")
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
