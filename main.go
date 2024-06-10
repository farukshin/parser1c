package main

import (
	"fmt"
	"time"
)

type Event struct {
	Time       time.Time
	Duration   uint64
	EventName  string
	EventNum   string
	Properties []*Properties
	Log        string
}

type Properties struct {
	key   string
	value string
}

type parser interface {
	parse(filename string, text string) ([]*Event, error)
}

func main() {
	fmt.Println("Hello parser")
}
