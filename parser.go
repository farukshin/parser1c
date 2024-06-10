package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func (p *parser) run() {

	p.checkFiles(p.input, 0)

}

func (p *parser) checkFiles(fileName string, pos int) {

	eventsString, _, _ := parseFile(fileName, pos)
	time, _ := p.GetTimeFromFileName(fileName)
	events, _ := p.getEvetsFromString(eventsString, time)

	var sb strings.Builder
	sb.WriteString("{\"events\":[\n")

	for i, s := range events {
		if i != 0 {
			sb.WriteString(",\n")
		}
		res1B, _ := json.Marshal(s)
		sb.WriteString(string(res1B))
	}
	sb.WriteString("]}")
	fmt.Print(sb.String())
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
	if p.debug == "1" || p.debug == "true" {
		ev.Log = log
	}
	ev.EventName = log[firstComma+1 : secondComma]
	ev.EventNum = log[secondComma+1 : comma3]
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
	for _, p := range prop {
		properties, success := checkProperties(ost + p)
		if success {
			ev.Properties = append(ev.Properties, properties)
			ost = ""
		} else {
			ost = ost + p
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

func (p *parser) getEvetsFromString(events []string, time time.Time) ([]*Event, error) {
	var res []*Event

	for _, ev := range events {
		e, err := p.parseLogLine(ev, time)
		if err != nil {
			continue
		}
		res = append(res, e)
	}
	return res, nil
}

func parseFile(fileName string, offset int) ([]string, uint64, error) {

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
	return res, uint64(pos), nil
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
		panic(e)
	}
}

func checkProperties(p string) (*Properties, bool) {

	pos_ravno := strings.Index(p, "=")
	if pos_ravno > 0 && len(p)-1 > pos_ravno && pos_ravno != 0 &&
		(p[pos_ravno+1] != '\'' ||
			p[pos_ravno+1] == '\'' && len(p)-2 > pos_ravno && p[len(p)-1] == '\'') {
		res := Properties{}
		res.Key = strToFieldName(p[0:pos_ravno])
		res.Value = p[pos_ravno+1:]
		return &res, true
	} else if pos_ravno > 0 && len(p)-1 == pos_ravno && pos_ravno != 0 {
		res := Properties{}
		res.Key = strToFieldName(p[0:pos_ravno])
		res.Value = p[pos_ravno+1:]
		return &res, true
	}
	return nil, false
}

func strToFieldName(str string) string {
	res := strings.ToLower(str)
	re := regexp.MustCompile("[a-z0-9]+")
	res = strings.Join(re.FindAllString(res, -1), "")
	return res
}
