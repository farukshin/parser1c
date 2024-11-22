package main

import (
	"fmt"
	"sync"
)

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
	MapFilesOffset    map[string]int64  `json:"mapFilesOffset"`
}

type JobInput struct {
	FileName string
	Offset   int64
}

type JobOutput struct {
	FileName   string
	Offset     int64
	OffsetLast int64
	Events     []*Event
}

func (p *Properties) String() string {
	return fmt.Sprintf("{\"%s\": \"%s\"}", p.Key, p.Value)
}

func (e *Event) String() string {

	return fmt.Sprintf("{\"time\":\"%s\",\"duration\":%d,\"name\":\"%s\",\"level\":\"%d\",\"log\":\"%s\"}",
		e.Time.Format("2006.01.02 15:04:05"),
		e.Duration,
		e.Name,
		e.EventLevel,
		e.Log)

}
