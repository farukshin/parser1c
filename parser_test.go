package main

import (
	"io/ioutil"
	"log"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestIsNumber_ps(t *testing.T) {
	var n byte = '1'
	expected := true
	actual := isNumber(n)
	assert.Equal(t, expected, actual)
}
func TestIsNumber_ng(t *testing.T) {
	var n byte = 'a'
	expected := false
	actual := isNumber(n)
	assert.Equal(t, expected, actual)
}

func TestIsBeginEvent_ps(t *testing.T) {
	var s string = "00:00.470006-0,CONN,0,process="
	expected := true
	actual := isBeginEvent(s)
	assert.Equal(t, expected, actual)
}

func TestIsBeginEvent_ng1(t *testing.T) {
	var s string = "0a:00.470006-0,CONN,0,process="
	expected := false
	actual := isBeginEvent(s)
	assert.Equal(t, expected, actual)
}

func TestIsBeginEvent_ng(t *testing.T) {
	var s string = "123"
	expected := false
	actual := isBeginEvent(s)
	assert.Equal(t, expected, actual)
}

func TestIsLowerChar_ps(t *testing.T) {
	var n byte = 'p'
	expected := true
	actual := isLowerChar(n)
	assert.Equal(t, expected, actual)
}
func TestIsLowerChar_ng(t *testing.T) {
	var n byte = 'P'
	expected := false
	actual := isLowerChar(n)
	assert.Equal(t, expected, actual)
}

func TestIsChar_ps(t *testing.T) {
	var n byte = 'G'
	expected := true
	actual := isChar(n)
	assert.Equal(t, expected, actual)
}
func TestIsChar_ng(t *testing.T) {
	var n byte = '#'
	expected := false
	actual := isChar(n)
	assert.Equal(t, expected, actual)
}

func TestGetTimeFromFileName_ps(t *testing.T) {
	p := parser{}
	s := "24051505.log"
	actual, err := p.GetTimeFromFileName(s)
	expected := time.Date(2024, time.May, 15, 05, 0, 0, 0, time.UTC)
	assert.Nil(t, err)
	assert.Equal(t, expected, actual)
}

func TestGetTimeFromFileName_ng(t *testing.T) {
	p := parser{}
	s := "24051505.log"
	actual, err := p.GetTimeFromFileName(s)
	expected := time.Date(2024, time.May, 15, 0, 0, 0, 0, time.UTC)
	assert.Nil(t, err)
	assert.NotEqual(t, expected, actual)
}

func TestInitMapFieldName_ps(t *testing.T) {
	p := parser{}
	p.initMapFieldName()
	actual, ok := p.MapFieldName["OSThread"]
	expected := "osthread"
	assert.Equal(t, true, ok)
	assert.Equal(t, expected, actual)
}

func TestInitMapFieldName_ng(t *testing.T) {
	p := parser{}
	p.initMapFieldName()
	actual, _ := p.MapFieldName["array"]
	expected := "Arrya"
	assert.NotEqual(t, expected, actual)
}

func TestParseLogLine_ps(t *testing.T) {
	p := parser{}
	p.initMapFieldName()
	tTime := time.Date(2024, time.May, 15, 0, 0, 0, 0, time.UTC)
	line := "00:00.673002-0,CONN,1,process=ragent,OSThread=8776,Txt=Clnt: MyUserName2: "
	ev, err := p.parseLogLine(line, tTime)
	assert.Nil(t, err)
	assert.Equal(t, len(ev.Properties), 3)
}

func TestParseLogLine_ps1(t *testing.T) {
	p := parser{}
	p.initMapFieldName()
	tTime := time.Date(2024, time.May, 15, 0, 0, 0, 0, time.UTC)
	line := "00:00.673002-0,CONN,1,process=ragent,OSThread=8776,Txt116=Clnt: MyUserName2: "
	ev, err := p.parseLogLine(line, tTime)
	assert.Nil(t, err)
	assert.Equal(t, len(ev.Properties), 3)
}

func TestParseRun_ps(t *testing.T) {
	dname, err := tmpEventDir()
	defer os.RemoveAll(dname)
	assert.Nil(t, err)
	p := parser{Input: dname, Output: "./log.log", Format: "", Debug: "0", CountRuner: 4}
	p.initMapFieldName()
	err = p.run()
	assert.Nil(t, err)
}

func TestParseSaerchFile_ps(t *testing.T) {

	f, err := ioutil.TempFile("", "*.log")
	assert.Nil(t, err)
	for i := 0; i < 10; i++ {
		f.WriteString("00:00.673002-0,CONN,1,process=ragent,OSThread=8776,Txt116=Clnt: MyUserName2:")
	}
	f.Close()

	p := parser{Input: f.Name(), Output: "./log.log", Format: "", Debug: "0", CountRuner: 4}
	p.initMapFieldName()
	err = p.run()
	assert.Nil(t, err)
}

func tmpEventDir() (string, error) {
	arr := []string{"00:00.673002-0,CONN,1,process=ragent,OSThread=8776,Txt116=Clnt: MyUserName2: ",
		"00:11.658012-0,CLSTR,0,process=rmngr,OSThread=1780,Event=Performance update,Data='process=tcp://s-msk-p-csd-as1:1541,pid=5256,cpu=0,queue_length=0,queue_length/cpu_num=0,memory_performance=22,disk_performance=17,response_time=39,average_response_time=30.78'",
		"10:54.696021-0,EXCP,0,process=rphost,OSThread=25616,Exception=dd149677-3d47-4e05-a55f-4e75b13a441f,Descr='src\rserver\\src\\RMngrCalls.cpp(516):\n		dd149677-3d47-4e05-a55f-4e75b13a441f: Процесс завершается. Исходящий вызов запрещен.'"}

	dname, err := ioutil.TempDir("", "")
	if err != nil {
		return "", err
	}

	for _, ev := range arr {
		f, err := ioutil.TempFile(dname, "*.log")
		if err != nil {
			return "", err
		}
		for i := 0; i < 10; i++ {
			f.WriteString(ev)
		}
		f.Close()
	}
	return dname, nil
}

func TestAppRun_ps(t *testing.T) {
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
		version:  "v0.1.4",
	}

	dname, err := tmpEventDir()
	defer os.RemoveAll(dname)
	assert.Nil(t, err)
	os.Args = append(os.Args, "--input="+dname)
	os.Args = append(os.Args, "--output=log.log")
	os.Args = append(os.Args, "--format=json")
	app.parseArgs()
}

func TestHelpHomeStr_ps(t *testing.T) {
	str := helpHomeStr()
	assert.Equal(t, str[:30], "Приложение: parser1c")
}
