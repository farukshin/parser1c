package main

import (
	"log"
	"os"
)

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
	version  string
}

func main() {

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
		version:  "v.0.1.0",
	}

	app.parseArgs()
}
