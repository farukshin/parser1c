package main

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

type Event struct {
	Time       time.Time     `json:"time"`
	Duration   uint64        `json:"duration"`
	EventName  string        `json:"eventName"`
	EventLevel string        `json:"eventLevel"`
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
	Input             string            `json:"input"`
	Output            string            `json:"output"`
	Format            string            `json:"format"`
	Debug             string            `json:"debug"`
	CountRuner        int               `json:"countRuner"`
	MapFieldName      map[string]string `json:"mapFieldName"`
	mapFieldNameMutex sync.RWMutex      `json:"mapFieldNameMutex"`
	Files             []string          `json:"files"`
	Events            []*Event          `json:"events"`
}

func (p *Properties) String() string {
	return fmt.Sprintf("{\"%s\": \"%s\"}", p.Key, p.Value)
}

func (e *Event) String() string {
	var l = make([]string, len(e.Properties))
	for i, p := range e.Properties {
		l[i] = p.String()
	}
	return fmt.Sprintf("{\"time\":\"%s\",\"duration\":%d,\"eventName\":\"%s\",\"eventLevel\":\"%s\",\"properties\":[%s],\"log\":\"%s\"}",
		e.Time.Format("2006.01.02 15:04:05"),
		e.Duration,
		e.EventName,
		e.EventLevel,
		strings.Join(l, ","),
		e.Log)

}
