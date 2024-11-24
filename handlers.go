package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func (app *application) getVersion() {
	fmt.Printf("parser1c %s\n", app.version)
}

func helpHomeStr() string {
	var sb strings.Builder
	sb.WriteString("Приложение: parser1c\n")
	sb.WriteString("\tПарсинг логов технологического журнала 1С\n\n")

	sb.WriteString("Строка запуска: parser1c [Опции]\n\n")
	sb.WriteString("Опции:\n")
	sb.WriteString("-h --help - вызов справки\n")
	sb.WriteString("-v --version - версия приложения\n")
	sb.WriteString("--input - каталог с логами технологического журнала или имя файла с логами\n")
	sb.WriteString("--output - приемник (на данный момент только postgres)\n")
	sb.WriteString("--host - хост PostgreSQL (либо env PG_HOST)\n")
	sb.WriteString("--port - порт PostgreSQL (либо env PG_PORT)\n")
	sb.WriteString("--user - пользователь PostgreSQL (либо env PG_USER)\n")
	sb.WriteString("--password - пароль PostgreSQL (либо env PG_PASSWORD)\n")
	sb.WriteString("--dbname - база данных PostgreSQL (либо env PG_DBNAME)\n")
	sb.WriteString("--countRuner - количество потоков парсера, по умолчанию 1\n\n")

	sb.WriteString("Пример запуска:\n")
	sb.WriteString("./parser1c --input=/var/log/1c --output=postgres --host=localhost --port=5432 --user=postgres --password=postgres --countrunner=4 --dbname=alsu")
	return sb.String()
}

func (app *application) help_home() {
	fmt.Println(helpHomeStr())
}

func (app *application) parseArgs() error {

	if len(os.Args) < 1 || isArgs("--help") || isArgs("-h") {
		app.help_home()
	} else if isArgs("--version") || isArgs("-v") {
		app.getVersion()
	} else {
		app.parse()
	}
	return nil
}

func getArgs(a1 string) (string, error) {

	for _, s := range os.Args[1:] {
		if s == a1 {
			return "", nil
		}
		for i := 0; i < len(s); i++ {
			if s[i] == '=' && i > 0 {
				v := s[:i]
				if v == a1 {
					return s[i+1:], nil
				}
			}
		}

	}
	return "", errors.New("Не найдено флага " + a1)
}

func isArgs(a1 string) bool {

	_, err := getArgs(a1)
	return err == nil

}
func isArgsAll(ar string) bool {
	mas := strings.Split(ar, ",")
	res := true
	for _, a := range mas {
		res = res && isArgs(a)
	}
	return res
}

func (app *application) parse() {

	if !isArgsAll("--input,--output") {
		app.help_home()
		return
	}
	input, erri := getArgs("--input")
	output, erro := getArgs("--output")
	debug, _ := getArgs("--debug")
	countRunerStr, _ := getArgs("--countRuner")
	if input == "" || erri != nil || erro != nil || output == "" {
		app.help_home()
		return
	}
	countRuner, _ := strconv.Atoi(countRunerStr)
	p := parser{Input: input, Output: output, Debug: debug, CountRuner: countRuner}
	p.run()
}
