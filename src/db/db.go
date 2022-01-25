package db

import (
	"database/sql"
	"log"
	"todo_server/src/utils"

	_ "github.com/mattn/go-sqlite3"
)

type Connection struct {
	DbName string
}

func (conn *Connection) Open() *sql.DB {
	log.Print("Connecting db.")
	db, err := sql.Open("sqlite3", conn.DbName)
	utils.Check(err)
	return db
}
