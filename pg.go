package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
)

type DataBase struct {
	DB *sql.DB `json:"time"`
	//TjFiles []TjFiles `json:"tj_files"`
}

type TjFiles struct {
	Name string `json:"name"`
	Size int64  `json:"size"`
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
	_, err = dbs.DB.Exec(`CREATE TABLE IF NOT EXISTS events (time TIMESTAMP, duration INT, name VARCHAR(255), level VARCHAR(255), log TEXT)`)
	if err != nil {
		return err
	}
	return err
}

func (dbs *DataBase) loadTjFilesStat() ([]TjFiles, error) {
	rows, err := dbs.DB.Query("SELECT name, size, time FROM tj_files")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var tjFiles []TjFiles
	for rows.Next() {
		var tjFile TjFiles
		err := rows.Scan(&tjFile.Name, &tjFile.Size, &tjFile.Time)
		if err != nil {
			return nil, err
		}
		tjFiles = append(tjFiles, tjFile)
	}
	return tjFiles, nil
}

func (dbs *DataBase) updateTjFiles(fileName string, seed int64) error {
	_, err := dbs.DB.Exec("UPDATE tj_files SET size = $1, time = $3 WHERE name = $2", seed, fileName, time.Now())
	return err
}

func (dbs *DataBase) addTjFiles(fileName string, seed int64) error {
	_, err := dbs.DB.Exec("INSERT INTO tj_files (name, size, time) VALUES ($1, $2, $3)", fileName, seed, time.Now())
	return err
}

func (dbs *DataBase) loadTjFiles() ([]TjFiles, error) {

	curTjFiles := make([]TjFiles, 0)
	rows, err := db.DB.Query("SELECT name, size, time FROM tj_files")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var tjFile TjFiles
		err := rows.Scan(&tjFile.Name, &tjFile.Size, &tjFile.Time)
		if err != nil {
			return nil, err
		}
		curTjFiles = append(curTjFiles, tjFile)
	}
	return curTjFiles, nil
}

func (dbs *DataBase) saveEvents(fileName string, events []*Event) error {
	for _, e := range events {
		_, err := dbs.DB.Exec("INSERT INTO events (time, duration, name, level, log) VALUES ($1, $2, $3, $4, $5)",
			e.Time, e.Duration, e.Name, e.Level, e.Log)
		if err != nil {
			return err
		}
	}
	return nil
}
