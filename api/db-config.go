package main

import (
	"database/sql"
	"fmt"
	"log"
)

type dbInfo struct {
	port         int
	host         string
	user         string
	password     string
	databaseName string
}

var DbInfo dbInfo = dbInfo{
	port:         5432,
	databaseName: "localstore",
	host:         "localhost",
	user:         "postgres",
	password:     "password",
}

var psqlInfo string = fmt.Sprintf(
	"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
	DbInfo.host, DbInfo.port,
	DbInfo.user, DbInfo.password,
	DbInfo.databaseName,
)

func openDB() *sql.DB {
	db, err := sql.Open("postgres", psqlInfo)
	err = db.Ping()
	if err != nil {
		return nil
	}
	log.Println("DB opened.......")

	return db
}
