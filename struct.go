package main

import (
	"fmt"
	"time"
)

type Event struct {
	Time       time.Time     `json:"time"`
	Duration   uint64        `json:"duration"`
	EventName  string        `json:"eventName"`
	EventNum   string        `json:"eventNum"`
	Properties []*Properties `json:"properties"`
	Log        string        `json:"log"`
}

type Properties struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type parserImp interface {
	setCatalog(cat string)
	getAllFiles() []string
	parse(filename string, text string) ([]*Event, error)
}

type parser struct {
	input  string
	format string
	debug  string
	events []*Event
}

func (e *Properties) String() string {
	return fmt.Sprintf("{\"%s\": \"%s\"}", e.Key, e.Value)
}
