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
	appid              string    `json:"appid"`
	trans              int       `json:"trans"`
	rows               int       `json:"rows"`
	dstclientid        int       `json:"dstclientid"`
	_interface         string    `json:"interface"`
	iname              string    `json:"iname"`
	method             int       `json:"method"`
	memory             int       `json:"memory"`
	memorypeak         int       `json:"memorypeak"`
	inbytes            int       `json:"inbytes"`
	outbytes           int       `json:"outbytes"`
	cputime            int       `json:"cputime"`
	waitconnections    int       `json:"waitconnections"`
	dbpid              int       `json:"dbpid"`
	rowsaffected       int       `json:"rowsaffected"`
	body               int       `json:"body"`
	status             int       `json:"status"`
	callwait           int       `json:"callwait"`
	regions            string    `json:"regions"`
	locks              string    `json:"locks"`
	sql                string    `json:"sql"`
	uri                string    `json:"uri"`
	headers            string    `json:"headers"`
	phrase             string    `json:"phrase"`
	first              string    `json:"first"`
	ablename           string    `json:"ablename"`
	prm                string    `json:"prm"`
	processname        string    `json:"processname"`
	srcprocessname     string    `json:"srcprocessname"`
	tablename          string    `json:"tablename"`
	callQlevel         string    `json:"callQlevel"`
	retexcp            string    `json:"retexcp"`
	scallPlevel        string    `json:"scallPlevel"`
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
		ev.tclientID = setIntFromString(prop.Value)
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
		ev.tconnectID = setIntFromString(prop.Value)
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
	case "appid":
		ev.appid = prop.Value
	case "trans":
		ev.trans = setIntFromString(prop.Value)
	case "rows":
		ev.rows = setIntFromString(prop.Value)
	case "dstclientid":
		ev.dstclientid = setIntFromString(prop.Value)
	case "interface":
		ev._interface = prop.Value
	case "iname":
		ev.iname = prop.Value
	case "method":
		ev.method = setIntFromString(prop.Value)
	case "memory":
		ev.memory = setIntFromString(prop.Value)
	case "memorypeak":
		ev.memorypeak = setIntFromString(prop.Value)
	case "inbytes":
		ev.inbytes = setIntFromString(prop.Value)
	case "outbytes":
		ev.outbytes = setIntFromString(prop.Value)
	case "cputime":
		ev.cputime = setIntFromString(prop.Value)
	case "waitconnections":
		ev.waitconnections = setIntFromString(prop.Value)
	case "dbpid":
		ev.dbpid = setIntFromString(prop.Value)
	case "rowsaffected":
		ev.rowsaffected = setIntFromString(prop.Value)
	case "body":
		ev.body = setIntFromString(prop.Value)
	case "status":
		ev.status = setIntFromString(prop.Value)
	case "callwait":
		ev.callwait = setIntFromString(prop.Value)
	case "regions":
		ev.regions = prop.Value
	case "locks":
		ev.locks = prop.Value
	case "sql":
		ev.sql = prop.Value
	case "uri":
		ev.uri = prop.Value
	case "headers":
		ev.headers = prop.Value
	case "phrase":
		ev.phrase = prop.Value
	case "first":
		ev.first = prop.Value
	case "ablename":
		ev.ablename = prop.Value
	case "prm":
		ev.prm = prop.Value
	case "processname":
		ev.processname = prop.Value
	case "srcprocessname":
		ev.srcprocessname = prop.Value
	case "tablename":
		ev.tablename = prop.Value
	case "callQlevel":
		ev.callQlevel = prop.Value
	case "retexcp":
		ev.retexcp = prop.Value
	case "scallPlevel":
		ev.scallPlevel = prop.Value

	default:
		if _, ok := notFoundProp[prop.Key]; !ok {
			notFoundProp[prop.Key] = true
		}
	}
}

func setIntFromString(val string) int {
	res, err := getIntFromString(val)
	if err == nil {
		return res
	}
	return 0
}

var notFoundProp = make(map[string]bool)
