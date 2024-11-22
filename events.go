package main

import "time"

type Event struct {
	Time               time.Time `json:"time"`
	Duration           uint64    `json:"duration"`
	Name               string    `json:"name"`
	EventLevel         int       `json:"eventlevel"`
	Log                string    `json:"log"`
	ConnectString      string    `json:"connectstring"`
	ServiceName        string    `json:"servicename"`
	res                string    `json:"res"`
	OSThread           string    `json:"osthread"`
	ExtData            string    `json:"extdata"`
	SESN1process       string    `json:"sesn1process"`
	ClientID           string    `json:"clientid"`
	Err                string    `json:"err"`
	Appl               string    `json:"appl"`
	DstId              string    `json:"dstid"`
	pprocessName       string    `json:"pprocessname"`
	DataBase           string    `json:"database"`
	Url                string    `json:"url"`
	Event              string    `json:"event"`
	SrcId              string    `json:"srcid"`
	ID                 string    `json:"id"`
	Info               string    `json:"info"`
	process            string    `json:"process"`
	ATTN0process       string    `json:"attn0process"`
	tclientID          int       `json:"tclientid"`
	IB                 string    `json:"ib"`
	TargetCall         string    `json:"targetcall"`
	DBMS               string    `json:"dbms"`
	Context            string    `json:"context"`
	SrcName            string    `json:"srcname"`
	tapplicationName   string    `json:"tapplicationname"`
	ApplicationExt     string    `json:"applicationext"`
	Data               string    `json:"data"`
	Protected          string    `json:"protected"`
	ProcessId          string    `json:"processid"`
	tcomputerName      string    `json:"tcomputername"`
	DstAddr            string    `json:"dstaddr"`
	SessionID          string    `json:"sessionid"`
	AgentUrl           string    `json:"agenturl"`
	CONN0process       string    `json:"conn0process"`
	ClientComputerName string    `json:"clientcomputername"`
	DstPid             string    `json:"dstpid"`
	DistribData        string    `json:"distribdata"`
	RmngrURL           string    `json:"rmngrurl"`
	CONN2process       string    `json:"conn2process"`
	CallID             string    `json:"callid"`
	Result             string    `json:"result"`
	Request            string    `json:"request"`
	Pid                string    `json:"pid"`
	InfoBase           string    `json:"infobase"`
	Message            string    `json:"message"`
	ServerComputerName string    `json:"servercomputername"`
	tconnectID         int       `json:"tconnectid"`
	Usr                string    `json:"usr"`
	CONN1process       string    `json:"conn1process"`
	Administrator      string    `json:"administrator"`
	SrcAddr            string    `json:"srcaddr"`
	MName              string    `json:"mname"`
	EXCP0process       string    `json:"excp0process"`
	Ref                string    `json:"ref"`
	Nmb                string    `json:"nmb"`
	UserName           string    `json:"username"`
	Func               string    `json:"func"`
	SrcPid             string    `json:"srcpid"`
	Calls              string    `json:"calls"`
	Txt                string    `json:"txt"`
	Descr              string    `json:"descr"`
	Exception          string    `json:"exception"`
	Level              string    `json:"level"`
	SDBL               string    `json:"sdbl"`
}

func (ev *Event) setProrerites(prop *Properties) {

	switch prop.Key {
	case "connectstring":
		ev.ConnectString = prop.Value
	case "servicename":
		ev.ServiceName = prop.Value
	case "res":
		ev.res = prop.Value
	case "osthread":
		ev.OSThread = prop.Value
	case "extdata":
		ev.ExtData = prop.Value
	case "sesn1process":
		ev.SESN1process = prop.Value
	case "clientid":
		ev.ClientID = prop.Value
	case "err":
		ev.Err = prop.Value
	case "appl":
		ev.Appl = prop.Value
	case "dstid":
		ev.DstId = prop.Value
	case "pprocessname":
		ev.pprocessName = prop.Value
	case "database":
		ev.DataBase = prop.Value
	case "url":
		ev.Url = prop.Value
	case "event":
		ev.Event = prop.Value
	case "srcid":
		ev.SrcId = prop.Value
	case "id":
		ev.ID = prop.Value
	case "info":
		ev.Info = prop.Value
	case "process":
		ev.process = prop.Value
	case "attn0process":
		ev.ATTN0process = prop.Value
	case "tclientid":
		val, err := getIntFromString(prop.Value)
		if err == nil {
			ev.tclientID = val
		}
	case "ib":
		ev.IB = prop.Value
	case "targetcall":
		ev.TargetCall = prop.Value
	case "dbms":
		ev.DBMS = prop.Value
	case "context":
		ev.Context = prop.Value
	case "srcname":
		ev.SrcName = prop.Value
	case "tapplicationname":
		ev.tapplicationName = prop.Value
	case "applicationext":
		ev.ApplicationExt = prop.Value
	case "data":
		ev.Data = prop.Value
	case "protected":
		ev.Protected = prop.Value
	case "processid":
		ev.ProcessId = prop.Value
	case "tcomputername":
		ev.tcomputerName = prop.Value
	case "dstaddr":
		ev.DstAddr = prop.Value
	case "sessionid":
		ev.SessionID = prop.Value
	case "agenturl":
		ev.AgentUrl = prop.Value
	case "conn0process":
		ev.CONN0process = prop.Value
	case "clientcomputername":
		ev.ClientComputerName = prop.Value
	case "dstpid":
		ev.DstPid = prop.Value
	case "distribdata":
		ev.DistribData = prop.Value
	case "rmngrurl":
		ev.RmngrURL = prop.Value
	case "conn2process":
		ev.CONN2process = prop.Value
	case "callid":
		ev.CallID = prop.Value
	case "result":
		ev.Result = prop.Value
	case "request":
		ev.Request = prop.Value
	case "pid":
		ev.Pid = prop.Value
	case "infobase":
		ev.InfoBase = prop.Value
	case "message":
		ev.Message = prop.Value
	case "servercomputername":
		ev.ServerComputerName = prop.Value
	case "tconnectid":
		val, err := getIntFromString(prop.Value)
		if err == nil {
			ev.tconnectID = val
		}
	case "usr":
		ev.Usr = prop.Value
	case "conn1process":
		ev.CONN1process = prop.Value
	case "administrator":
		ev.Administrator = prop.Value
	case "srcaddr":
		ev.SrcAddr = prop.Value
	case "mname":
		ev.MName = prop.Value
	case "excp0process":
		ev.EXCP0process = prop.Value
	case "ref":
		ev.Ref = prop.Value
	case "nmb":
		ev.Nmb = prop.Value
	case "username":
		ev.UserName = prop.Value
	case "func":
		ev.Func = prop.Value
	case "srcpid":
		ev.SrcPid = prop.Value
	case "calls":
		ev.Calls = prop.Value
	case "txt":
		ev.Txt = prop.Value
	case "descr":
		ev.Descr = prop.Value
	case "exception":
		ev.Exception = prop.Value
	case "level":
		ev.Level = prop.Value
	case "sdbl":
		ev.SDBL = prop.Value
	default:
		ev.Exception = ev.Exception
	}
}
