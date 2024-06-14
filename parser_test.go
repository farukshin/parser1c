package main

import (
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
