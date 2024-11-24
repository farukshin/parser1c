package main

import (
	"errors"
	"io"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"time"
)

func (p *parser) worker(id int, jobs <-chan JobInput, results chan<- JobOutput) {
	for j := range jobs {
		out := p.checkFiles(j.FileName, j.Offset)
		results <- out
	}
}

func (p *parser) searchFiles() ([]string, error) {

	foundFiles := make([]string, 0)

	f, err := os.Stat(p.Input)
	if err != nil {
		return nil, err
	}
	if f.IsDir() {
		res := p.ListDir(p.Input)
		for _, file := range res {
			foundFiles = append(foundFiles, file)
		}
	} else if len(p.Input) >= 4 && p.Input[len(p.Input)-4:] == ".log" {
		foundFiles = append(foundFiles, p.Input)
	} else {
		return nil, errors.New("Не удалось прочитать из input = " + p.Input)
	}
	return foundFiles, nil

}
func (p *parser) run() error {

	if p.Output == "postgres" {
		err := db.Init()
		if err != nil {
			app.errorLog.Println(err)
			return err
		}
	}

	p.initMapFieldName()
	foundFiles, err := p.searchFiles()
	if err != nil {
		return err
	}
	p.loadMapFilesSeek(foundFiles)

	if p.CountRuner <= 0 {
		p.CountRuner = 1
	}
	cntWorkers := p.CountRuner
	jobs := make(chan JobInput, 1000) //todo
	results := make(chan JobOutput, 1000)
	for w := 1; w <= cntWorkers; w++ {
		go p.worker(w, jobs, results)
	}
	for filename, offset := range p.MapFilesOffset {
		jobs <- JobInput{FileName: filename, Offset: offset}
	}

	var cnt = 0
	for true {
		cnt++
		out := <-results
		if out.Offset != out.OffsetLast {
			db.updateTjFiles(out.FileName, out.Offset)
			p.MapFilesOffset[out.FileName] = out.Offset
		}
		if out.Events != nil && len(out.Events) > 0 {
			if p.Format == "postgres" {
				err := db.saveEvents(out.FileName, out.Events)
				if err != nil {
					return err
				}
			}
		}
		jobs <- JobInput{FileName: out.FileName, Offset: out.Offset}

		if cnt == 10 {
			cnt = 0
			foundFiles, err := p.searchFiles()
			if err == nil {
				for _, file := range foundFiles {
					if _, ok := p.MapFilesOffset[file]; !ok {
						jobs <- JobInput{FileName: file, Offset: 0}
						p.MapFilesOffset[file] = 0
					}
				}
			}
		}
		time.Sleep(1 * time.Second)
	}
	return nil
}

func (p *parser) loadMapFilesSeek(foundFiles []string) {
	p.MapFilesOffset = make(map[string]int64)

	var curMapFilesSeek = make(map[string]int64)
	TjFiles, err := db.loadTjFiles()
	if err != nil {
		panic(err)
	}
	for _, file := range TjFiles {
		curMapFilesSeek[file.Name] = file.Size
	}

	for _, file := range foundFiles {
		if _, ok := curMapFilesSeek[file]; !ok {
			p.MapFilesOffset[file] = 0
			db.addTjFiles(file, 0)
		} else {
			p.MapFilesOffset[file] = curMapFilesSeek[file]
		}
	}
}

func (p *parser) initMapFieldName() {
	p.MapFieldName = map[string]string{"ConnectString": "connectstring", "ServiceName": "servicename", "res": "res", "OSThread": "osthread", "ExtData": "extdata", "SESN1process": "sesn1process", "ClientID": "clientid", "Err": "err", "Appl": "appl", "DstId": "dstid", "p:processName": "pprocessname", "DataBase": "database", "Url": "url", "Event": "event", "SrcId": "srcid", "ID": "id", "Info": "info", "process": "process", "ATTN0process": "attn0process", "t:clientID": "tclientid", "IB": "ib", "TargetCall": "targetcall", "DBMS": "dbms", "Context": "context", "SrcName": "srcname", "t:applicationName": "tapplicationname", "ApplicationExt": "applicationext", "Data": "data", "Protected": "protected", "ProcessId": "processid", "t:computerName": "tcomputername", "DstAddr": "dstaddr", "SessionID": "sessionid", "txt": "txt", "AgentUrl": "agenturl", "CONN0process": "conn0process", "ClientComputerName": "clientcomputername", "DstPid": "dstpid", "DistribData": "distribdata", "RmngrURL": "rmngrurl", "CONN2process": "conn2process", "CallID": "callid", "Result": "result", "Request": "request", "Pid": "pid", "InfoBase": "infobase", "Message": "message", "ServerComputerName": "servercomputername", "t:connectID": "tconnectid", "Usr": "usr", "CONN1process": "conn1process", "Administrator": "administrator", "SrcAddr": "srcaddr", "MName": "mname", "EXCP0process": "excp0process", "Ref": "ref", "Nmb": "nmb", "UserName": "username", "Func": "func", "SrcPid": "srcpid", "Calls": "calls", "Txt": "txt", "Descr": "descr", "Exception": "exception"}
}

func (p *parser) ListDir(path string) []string {

	res := make([]string, 0)
	if len(path) > 0 && path[len(path)-1] != '/' {
		path = path + "/"
	}

	lst, err := ioutil.ReadDir(path)
	if err != nil {
		return make([]string, 0)
	}
	for _, val := range lst {
		name := val.Name()
		if val.IsDir() {
			cur := p.ListDir(path + val.Name())
			for _, file := range cur {
				res = append(res, file)
			}
		} else if len(name) > 4 && name[len(name)-4:] == ".log" {
			res = append(res, path+name)
		}
	}
	return res
}

func (p *parser) checkFiles(fileName string, pos int64) JobOutput {

	res := JobOutput{}
	eventsString, newpos, _ := parseFile(fileName, pos)
	time, _ := p.GetTimeFromFileName(fileName)
	events, _ := p.getEvetsFromString(eventsString, time)

	res.FileName = fileName
	res.Offset = newpos
	res.OffsetLast = pos
	res.Events = events

	return res
}

func (p *parser) GetTimeFromFileName(fileName string) (time.Time, error) {
	if len(fileName) < 12 {
		return time.Time{}, errors.New("Некорректный формат файла")
	}
	str := fileName[len(fileName)-12 : len(fileName)-4]

	year_str := str[0:2]
	month_str := str[2:4]
	day_str := str[4:6]
	hour_str := str[6:8]

	year := 0
	month := 0
	day := 0
	hour := 0
	err := errors.New("")

	if year, err = strconv.Atoi(year_str); err != nil {
		return time.Time{}, errors.New("Некорректный формат файла")
	}
	if month, err = strconv.Atoi(month_str); err != nil {
		return time.Time{}, errors.New("Некорректный формат файла")
	}
	if day, err = strconv.Atoi(day_str); err != nil {
		return time.Time{}, errors.New("Некорректный формат файла")
	}
	if hour, err = strconv.Atoi(hour_str); err != nil {
		return time.Time{}, errors.New("Некорректный формат файла")
	}
	year = year + 2000
	then := time.Date(year, time.Month(month), day, hour, 00, 00, 0, time.UTC)
	return then, nil
}

func isNumber(ch byte) bool {
	return ch >= '0' && ch <= '9'
}

func isLowerChar(ch byte) bool {
	return ch >= 'a' && ch <= 'z'
}
func isChar(ch byte) bool {
	return ch >= 'A' && ch <= 'Z' || ch >= 'a' && ch <= 'z'
}

func isBeginEvent(str string) bool {
	success := len(str) >= 15 && str[2] == ':' && str[5] == '.' && str[12] == '-'
	if !success {
		return success
	}
	success = success && isNumber(str[0]) && isNumber(str[1]) &&
		isNumber(str[3]) && isNumber(str[4]) &&
		isNumber(str[6]) && isNumber(str[7]) && isNumber(str[8]) && isNumber(str[9]) &&
		isNumber(str[10]) && isNumber(str[11]) && isNumber(str[13])
	if !success {
		return success
	}
	isComma := false
	for i := 13; i < len(str); i++ {
		if !isNumber(str[i]) {
			isComma = str[i] == ','
			break
		}
	}
	return isComma
}

func (p *parser) parseLogLine(log string, timebase time.Time) (*Event, error) {
	if !isBeginEvent(log) {
		return nil, errors.New("Не удалось распарсить лог '" + log + "'")
	}
	tire1 := strings.Index(log, "-")
	firstComma := strings.Index(log, ",")
	secondComma := firstComma + 1 + strings.Index(log[firstComma+1:], ",")
	comma3 := secondComma + 1 + strings.Index(log[secondComma+1:], ",")

	ev := Event{}
	if p.Debug == "1" || p.Debug == "true" {
		ev.Log = log
	}
	ev.Name = log[firstComma+1 : secondComma]

	eventlevel, err2 := getIntFromString(log[secondComma+1 : comma3])
	if err2 != nil {
		return nil, errors.New("Не удалось распарсить лог '" + log + "'")
	}
	ev.EventLevel = eventlevel

	duration, err := getUint64FromString(log[tire1+1 : firstComma])
	if err != nil {
		return nil, errors.New("Не удалось распарсить лог '" + log + "'")
	}
	ev.Duration = duration
	curtime := getCurTime(log, timebase)
	ev.Time = curtime
	log = log[comma3+1:]

	prop := strings.Split(log, ",")
	ost := ""
	for _, pr := range prop {
		properties, success := p.checkProperties(ost + pr)
		if success {
			//ev.Properties = append(ev.Properties, properties)
			ev.setProrerites(properties)
			ost = ""
		} else {
			ost = ost + pr
		}
	}
	if ost != "" {
		return nil, errors.New("Не удалось распарсить лог '" + log + "'")
	}
	return &ev, nil
}

func getCurTime(log string, timebase time.Time) time.Time {
	if len(log) < 13 {
		return timebase
	}

	tm := timebase
	var min, sec, nano int
	var err error

	if min, err = strconv.Atoi(log[0:2]); err != nil {
		return timebase
	}
	if sec, err = strconv.Atoi(log[3:5]); err != nil {
		return timebase
	}
	if nano, err = strconv.Atoi(log[6:12]); err != nil {
		return timebase
	}
	return tm.Add(time.Minute*time.Duration(int64(min)) + time.Second*time.Duration(int64(sec)) + time.Nanosecond*time.Duration(int64(nano)))
}

func getUint64FromString(str string) (uint64, error) {
	var res uint64 = 0
	if len(str) > 0 {
		for _, ch := range str {
			//if !isNumber(ch)
			if !(ch >= '0' && ch <= '9') {
				return 0, errors.New("Некорректная строка")
			}
			res = res*10 + uint64(ch-'0')
		}
	}
	return res, nil
}

func getIntFromString(str string) (int, error) {
	var res int = 0
	if len(str) > 0 {
		for _, ch := range str {
			if !(ch >= '0' && ch <= '9') {
				return 0, errors.New("Некорректная строка")
			}
			res = res*10 + int(ch-'0')
		}
	}
	return res, nil
}

func (p *parser) getEvetsFromString(events []string, time time.Time) ([]*Event, error) {
	var res = make([]*Event, len(events))
	for i, ev := range events {
		e, err := p.parseLogLine(ev, time)
		if err != nil {
			//continue
		}
		//res = append(res, e)
		res[i] = e
	}
	return res, nil
}

func parseFile(fileName string, offset int64) ([]string, int64, error) {

	const portion = 1024

	var res []string
	f, err := os.Open(fileName)
	defer f.Close()
	check(err)

	ByteOrderMarkAsString := string('\uFEFF')

	pos, err := f.Seek(int64(offset), io.SeekCurrent)
	b1 := make([]byte, portion)
	var ost string
	for true {
		n1, err := f.Read(b1)
		if n1 == 0 {
			break
		}
		cur := string(b1[:n1])
		check(err)

		if strings.HasPrefix(cur, ByteOrderMarkAsString) {
			cur = strings.TrimPrefix(cur, ByteOrderMarkAsString)
		}
		var curres []string
		curres, ost = ostUpdate(ost + cur)
		res = append(res, curres...)
	}
	if len(ost) != 0 {
		res = append(res, ost)
	}
	pos, err = f.Seek(0, io.SeekCurrent)
	return res, pos, nil
}

func ostUpdate(cur string) ([]string, string) {

	var res []string
	lines := strings.Split(cur, "\r\n")
	if lines[len(lines)-1] == "" {
		lines = lines[0 : len(lines)-1]
	}
	pre := 0
	for idx, line := range lines {
		if idx == 0 {
			continue
		}
		if isBeginEvent(line) {
			res = append(res, strings.Join(lines[pre:idx], "\r\n"))
			pre = idx
		} else {
			a := 1
			if a == 1 {

			}
		}
	}
	cur = strings.Join(lines[pre:], "\r\n")
	return res, cur
}

func check(e error) {
	if e != nil {
		app.errorLog.Println(e)
		//panic(e)
	}
}

func (p *parser) checkProperties(str string) (*Properties, bool) {

	pos_ravno := strings.Index(str, "=")
	if pos_ravno > 0 && len(str)-1 > pos_ravno && pos_ravno != 0 &&
		(str[pos_ravno+1] != '\'' ||
			str[pos_ravno+1] == '\'' && len(str)-2 > pos_ravno && str[len(str)-1] == '\'') || pos_ravno > 0 && len(str)-1 == pos_ravno && pos_ravno != 0 {
		res := Properties{}
		res.Key = p.strToFieldName(str[0:pos_ravno])
		res.Value = str[pos_ravno+1:]
		return &res, true
	}
	return nil, false
}

func (p *parser) strToFieldName(str string) string {

	val, ok := p.MapFieldName[str]
	if ok {
		return val
	}

	var cur byte
	var res []byte
	for i := 0; i < len(str); i++ {
		if isChar(str[i]) || isNumber(str[i]) {
			cur = str[i]
			if !isLowerChar(str[i]) {
				cur = 'a' + (str[i] - 'A')
			}
			res = append(res, cur)
		}
	}
	return string(res)
}
