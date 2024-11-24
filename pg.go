package main

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"
	"time"

	_ "github.com/lib/pq"
)

type DataBase struct {
	DB *sql.DB `json:"time"`
	//TjFiles []TjFiles `json:"tj_files"`
}

type TjFiles struct {
	Name string `json:"name"`
	Size int64  `json:"size"`
	Time string `json:"time"`
}

var db = &DataBase{}

const (
	pg_host     = "localhost"
	pg_port     = 5432
	pg_user     = "postgres"
	pg_password = "postgres"
	pg_dbname   = "alsu"
)

func initArgs(str string, env string) (string, error) {

	val, errh := getArgs(str)
	if errh != nil {
		val = os.Getenv(env)
		if val == "" {
			return "", fmt.Errorf("Не задан параметр %s. Задайте параметр %s в команде запуска или иницализируйте переменную окружения %s", str, str, env)
		}
	}
	return val, nil
}

func (database *DataBase) openConnection() (*sql.DB, error) {

	host, err := initArgs("--host", "PG_HOST")
	if err != nil {
		return nil, err
	}
	port, err := initArgs("--port", "PG_PORT")
	if err != nil {
		return nil, err
	}
	portInt, err := strconv.Atoi(port)
	user, err := initArgs("--user", "PG_USER")
	if err != nil {
		return nil, err
	}
	password, err := initArgs("--password", "PG_PASSWORD")
	if err != nil {
		return nil, err
	}
	dbname, err := initArgs("--dbname", "PG_DBNAME")
	if err != nil {
		return nil, err
	}

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, portInt, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	return db, err
}

func (database *DataBase) Init() error {
	db, err := database.openConnection()
	if err != nil {
		return err
	}
	database.DB = db
	err = database.createTable()
	if err != nil {
		return err
	}
	return err
}

func (dbs *DataBase) createTable() error {
	_, err := dbs.DB.Exec(`CREATE TABLE IF NOT EXISTS tj_files (name VARCHAR(255) PRIMARY KEY, size INT, time TIMESTAMP)`)
	if err != nil {
		return err
	}
	_, err = dbs.DB.Exec(`CREATE TABLE IF NOT EXISTS events (time TIMESTAMP, 
		duration INT, 
		name VARCHAR(255), 
		eventlevel INT,
		log TEXT,
		ConnectString TEXT,
		ServiceName TEXT,
		res TEXT,
		OSThread TEXT,
		ExtData TEXT,
		SESN1process TEXT,
		ClientID TEXT,
		Err TEXT,
		Appl TEXT,
		DstId TEXT,
		pprocessName TEXT,
		DataBase TEXT,
		Url TEXT,
		Event TEXT,
		SrcId TEXT,
		ID TEXT,
		Info TEXT,
		process TEXT,
		ATTN0process TEXT,
		tclientID INT,
		IB TEXT,
		TargetCall TEXT,
		DBMS TEXT,
		Context TEXT,
		SrcName TEXT,
		tapplicationName TEXT,
		ApplicationExt TEXT,
		Data TEXT,
		Protected TEXT,
		ProcessId TEXT,
		tcomputerName TEXT,
		DstAddr TEXT,
		SessionID TEXT,
		AgentUrl TEXT,
		CONN0process TEXT,
		ClientComputerName TEXT,
		DstPid TEXT,
		DistribData TEXT,
		RmngrURL TEXT,
		CONN2process TEXT,
		CallID TEXT,
		Result TEXT,
		Request TEXT,
		Pid TEXT,
		InfoBase TEXT,
		Message TEXT,
		ServerComputerName TEXT,
		tconnectID INT,
		Usr TEXT,
		CONN1process TEXT,
		Administrator TEXT,
		SrcAddr TEXT,
		MName TEXT,
		EXCP0process TEXT,
		Ref TEXT,
		Nmb TEXT,
		UserName TEXT,
		Func TEXT,
		SrcPid TEXT,
		Calls TEXT,
		Txt TEXT,
		Descr TEXT,
		Exception TEXT,
		Level TEXT,
		SDBL TEXT,
		appid TEXT,
		trans INT,
		rows INT,
		dstclientid INT,
		interface TEXT,
		iname  TEXT,
		method INT,
		memory INT,
		memorypeak INT,
		inbytes INT,
		outbytes INT,
		cputime INT,
		waitconnections INT,
		dbpid INT,
		rowsaffected INT,
		body INT,
		status INT,
		callwait INT,
		regions TEXT,
		locks TEXT,
		sql TEXT,
		uri TEXT,
		headers TEXT,
		phrase TEXT,
		first TEXT,
		ablename TEXT,
		prm TEXT,
		processname TEXT,
		srcprocessname TEXT,
		tablename TEXT,
		callQlevel TEXT,
		retexcp TEXT,
		scallPlevel TEXT
		)`)
	if err != nil {
		return err
	}
	return err
}

func (dbs *DataBase) loadTjFilesStat() ([]TjFiles, error) {
	rows, err := dbs.DB.Query("SELECT name, size, time FROM tj_files")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var tjFiles []TjFiles
	for rows.Next() {
		var tjFile TjFiles
		err := rows.Scan(&tjFile.Name, &tjFile.Size, &tjFile.Time)
		if err != nil {
			return nil, err
		}
		tjFiles = append(tjFiles, tjFile)
	}
	return tjFiles, nil
}

func (dbs *DataBase) updateTjFiles(fileName string, seed int64) error {
	_, err := dbs.DB.Exec("UPDATE tj_files SET size = $1, time = $3 WHERE name = $2", seed, fileName, time.Now())
	return err
}

func (dbs *DataBase) addTjFiles(fileName string, seed int64) error {
	_, err := dbs.DB.Exec("INSERT INTO tj_files (name, size, time) VALUES ($1, $2, $3)", fileName, seed, time.Now())
	return err
}

func (dbs *DataBase) loadTjFiles() ([]TjFiles, error) {

	curTjFiles := make([]TjFiles, 0)
	rows, err := db.DB.Query("SELECT name, size, time FROM tj_files")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var tjFile TjFiles
		err := rows.Scan(&tjFile.Name, &tjFile.Size, &tjFile.Time)
		if err != nil {
			return nil, err
		}
		curTjFiles = append(curTjFiles, tjFile)
	}
	return curTjFiles, nil
}

func (dbs *DataBase) saveEvents(fileName string, events []*Event) error {
	if events == nil || len(events) == 0 {
		return nil
	}
	tx, err := dbs.DB.Begin()
	if err != nil {
		return err
	}
	for _, event := range events {
		if event == nil {
			continue
		}
		_, err = tx.Exec(`INSERT INTO events (time, duration, name, eventlevel, log, ConnectString, ServiceName, res, OSThread, ExtData, SESN1process, ClientID, Err, Appl, DstId, pprocessName, DataBase, Url, Event, SrcId, ID, Info, process, ATTN0process, tclientID, IB, TargetCall, DBMS, Context, SrcName, tapplicationName, ApplicationExt, Data, Protected, ProcessId, tcomputerName, DstAddr, SessionID, AgentUrl, CONN0process, ClientComputerName, DstPid, DistribData, RmngrURL, CONN2process, CallID, Result, Request, Pid, InfoBase, Message, ServerComputerName, tconnectID, Usr, CONN1process, Administrator, SrcAddr, MName, EXCP0process, Ref, Nmb, UserName, Func, SrcPid, Calls, Txt, Descr, Exception, Level, SDBL, 
				appid, trans, rows, dstclientid, interface, iname, method,
				memory, memorypeak, inbytes, outbytes, cputime, waitconnections, dbpid, rowsaffected, body, status, 
				callwait, regions, locks, sql, uri, headers, phrase, first, ablename, prm, processname, srcprocessname, 
				tablename, callQlevel, retexcp, scallPlevel
				) 
			VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25, $26, $27, $28, $29, $30, $31, $32, $33, $34, $35, $36, $37, $38, $39, $40, $41, $42, $43, $44, $45, $46, $47, $48, $49, $50, $51, $52, $53, $54, $55, $56, $57, $58, $59, $60, $61, $62, $63, $64, $65, $66, $67, $68, $69, $70, $71, 
			$72,$73,$74,$75,$76,$77,
			$78,$79,$80,$81,$82,$83,$84,$85,$86,$87,$88,$89,$90,$91,$92,$93,$94,$95,$96,$97,$98,$99,
			$100,$101,$102,$103)`,
			event.Time, event.Duration, event.Name, event.EventLevel, event.Log, event.ConnectString, event.ServiceName, event.res, event.OSThread, event.ExtData, event.SESN1process, event.ClientID, event.Err, event.Appl, event.DstId, event.pprocessName, event.DataBase, event.Url, event.Event, event.SrcId, event.ID, event.Info, event.process, event.ATTN0process, event.tclientID, event.IB, event.TargetCall, event.DBMS, event.Context, event.SrcName, event.tapplicationName, event.ApplicationExt, event.Data, event.Protected, event.ProcessId, event.tcomputerName, event.DstAddr, event.SessionID, event.AgentUrl, event.CONN0process, event.ClientComputerName, event.DstPid, event.DistribData, event.RmngrURL, event.CONN2process, event.CallID, event.Result, event.Request, event.Pid, event.InfoBase, event.Message, event.ServerComputerName, event.tconnectID, event.Usr, event.CONN1process, event.Administrator, event.SrcAddr, event.MName, event.EXCP0process, event.Ref, event.Nmb, event.UserName, event.Func, event.SrcPid, event.Calls, event.Txt, event.Descr, event.Exception, event.Level, event.SDBL, event.appid,
			event.trans, event.rows, event.dstclientid, event._interface, event.iname, event.method,
			event.memory, event.memorypeak, event.inbytes, event.outbytes, event.cputime, event.waitconnections,
			event.dbpid, event.rowsaffected, event.body, event.status, event.callwait, event.regions, event.locks,
			event.sql, event.uri, event.headers, event.phrase, event.first, event.ablename, event.prm, event.processname,
			event.srcprocessname,
			event.tablename, event.callQlevel, event.retexcp, event.scallPlevel)
		if err != nil {
			tx.Rollback()
		}
	}
	tx.Commit()
	return err
}
