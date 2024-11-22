package main

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

type Event struct {
	Time               time.Time     `json:"time"`
	Duration           uint64        `json:"duration"`
	Name               string        `json:"name"`
	Level              string        `json:"level"`
	Properties         []*Properties `json:"properties"`
	Log                string        `json:"log"`
	ConnectString      string        `json:"connectstring"`
	ServiceName        string        `json:"servicename"`
	res                string        `json:"res"`
	OSThread           string        `json:"osthread"`
	ExtData            string        `json:"extdata"`
	SESN1process       string        `json:"sesn1process"`
	ClientID           string        `json:"clientid"`
	Err                string        `json:"err"`
	Appl               string        `json:"appl"`
	DstId              string        `json:"dstid"`
	pprocessName       string        `json:"pprocessname"`
	DataBase           string        `json:"database"`
	Url                string        `json:"url"`
	Event              string        `json:"event"`
	SrcId              string        `json:"srcid"`
	ID                 string        `json:"id"`
	Info               string        `json:"info"`
	process            string        `json:"process"`
	ATTN0process       string        `json:"attn0process"`
	tclientID          string        `json:"tclientid"`
	IB                 string        `json:"ib"`
	TargetCall         string        `json:"targetcall"`
	DBMS               string        `json:"dbms"`
	Context            string        `json:"context"`
	SrcName            string        `json:"srcname"`
	tapplicationName   string        `json:"tapplicationname"`
	ApplicationExt     string        `json:"applicationext"`
	Data               string        `json:"data"`
	Protected          string        `json:"protected"`
	ProcessId          string        `json:"processid"`
	tcomputerName      string        `json:"tcomputername"`
	DstAddr            string        `json:"dstaddr"`
	SessionID          string        `json:"sessionid"`
	txt                string        `json:"txt"`
	AgentUrl           string        `json:"agenturl"`
	CONN0process       string        `json:"conn0process"`
	ClientComputerName string        `json:"clientcomputername"`
	DstPid             string        `json:"dstpid"`
	DistribData        string        `json:"distribdata"`
	RmngrURL           string        `json:"rmngrurl"`
	CONN2process       string        `json:"conn2process"`
	CallID             string        `json:"callid"`
	Result             string        `json:"result"`
	Request            string        `json:"request"`
	Pid                string        `json:"pid"`
	InfoBase           string        `json:"infobase"`
	Message            string        `json:"message"`
	ServerComputerName string        `json:"servercomputername"`
	tconnectID         string        `json:"tconnectid"`
	Usr                string        `json:"usr"`
	CONN1process       string        `json:"conn1process"`
	Administrator      string        `json:"administrator"`
	SrcAddr            string        `json:"srcaddr"`
	MName              string        `json:"mname"`
	EXCP0process       string        `json:"excp0process"`
	Ref                string        `json:"ref"`
	Nmb                string        `json:"nmb"`
	UserName           string        `json:"username"`
	Func               string        `json:"func"`
	SrcPid             string        `json:"srcpid"`
	Calls              string        `json:"calls"`
	Txt                string        `json:"txt"`
	Descr              string        `json:"descr"`
	Exception          string        `json:"exception"`
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

	var l = make([]string, len(e.Properties))
	for i, p := range e.Properties {
		l[i] = p.String()
	}
	return fmt.Sprintf("{\"time\":\"%s\",\"duration\":%d,\"eventName\":\"%s\",\"eventLevel\":\"%s\",\"properties\":[%s],\"log\":\"%s\"}",
		e.Time.Format("2006.01.02 15:04:05"),
		e.Duration,
		e.Name,
		e.Level,
		strings.Join(l, ","),
		e.Log)

}
