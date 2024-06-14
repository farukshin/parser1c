package main

import (
	"log"
	"os"
	"runtime/pprof"
)

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
	version  string
}

func main() {

	f, err := os.Create("cpu.prof")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	err = pprof.StartCPUProfile(f)
	if err != nil {
		panic(err)
	}
	defer pprof.StopCPUProfile()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
		version:  "v.0.1.3",
	}

	app.parseArgs()
}
