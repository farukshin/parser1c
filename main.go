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

var app = &application{
	version: "v0.1.6",
}

func main() {

	app.infoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	app.errorLog = log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	app.parseArgs()
}
