package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type DataBase struct {
	DB      *sql.DB   `json:"time"`
	TjFiles []TjFiles `json:"tj_files"`
}

type TjFiles struct {
	Name string `json:"name"`
	Size int    `json:"size"`
	Time string `json:"time"`
}

var db = &DataBase{}

const (
	pg_host     = "localhost"
	pg_port     = 5432
	pg_user     = "postgres"
	pg_password = "postgres"
	pg_dbname   = "alsu"
)

func (database *DataBase) openConnection() (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		pg_host, pg_port, pg_user, pg_password, pg_dbname)
	db, err := sql.Open("postgres", psqlInfo)
	return db, err
}

func (database *DataBase) Init() error {
	db, err := database.openConnection()
	if err != nil {
		return err
	}
	database.DB = db
	err = database.createTable()
	if err != nil {
		return err
	}
	return err
}

func (dbs *DataBase) createTable() error {
	_, err := dbs.DB.Exec(`CREATE TABLE IF NOT EXISTS tj_files (name VARCHAR(255) PRIMARY KEY, size INT, time TIMESTAMP)`)
	if err != nil {
		return err
	}
	_, err = dbs.DB.Exec(`CREATE TABLE IF NOT EXISTS events (id SERIAL PRIMARY KEY, time TIMESTAMP, duration INT, event_name VARCHAR(255), event_level VARCHAR(255), log TEXT)`)
	if err != nil {
		return err
	}
	return err
}

func loadTjFiles() error {
	rows, err := db.DB.Query("SELECT name, size, time FROM tj_files")
	if err != nil {
		return err
	}
	defer rows.Close()
	for rows.Next() {
		var tjFile TjFiles
		err := rows.Scan(&tjFile.Name, &tjFile.Size, &tjFile.Time)
		if err != nil {
			return err
		}
		db.TjFiles = append(db.TjFiles, tjFile)
	}
	return nil
}
