package database

import (
	"github.com/jackc/pgx"
	"io/ioutil"
)


var config = pgx.ConnConfig{
	Host:     "localhost",
	Port:     5432,
	Database: "postgres",
	User:     "postgres",
	Password: "123",
}

var Connection, _ = pgx.NewConnPool(
	pgx.ConnPoolConfig{
		ConnConfig:     config,
		MaxConnections: 50,
	})

func InitDataBase() error {
	initFile, _ := ioutil.ReadFile("init.sql")
	tx, err := Connection.Begin()
	if err != nil {
		return err
	}
	_, err = tx.Exec(string(initFile))
	if err != nil {
		tx.Rollback()
	}
	tx.Commit()
	Connection.Reset()
	return nil
}
