package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

func (app *application) getVersion() {
	fmt.Printf("parser_tj_1c %s\n", app.version)
}

func (app *application) help_home() {

	var sb strings.Builder
	sb.WriteString("Не корректное использование\n\n")

	sb.WriteString("Приложение: parser_tj_1c\n")
	sb.WriteString("  Парсинг логов технологического журнала 1С\n\n")

	sb.WriteString("Строка запуска: parser_tj_1c [Опции]\n\n")
	sb.WriteString("Опции:\n")
	sb.WriteString("-h --help - вызов справки\n")
	sb.WriteString("-v --version - версия приложения\n")
	sb.WriteString("--input - каталог с логами технологического журнала или имя файла с логами\n")
	sb.WriteString("--format - формат вывода")

	fmt.Println(sb.String())
}

func (app *application) parseArgs() {

	if len(os.Args) < 1 || isArgs("--help") || isArgs("-h") {
		app.help_home()
	} else if isArgs("--version") || isArgs("-v") {
		app.getVersion()
	} else {
		app.parse()
	}
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

	if !isArgsAll("--input,--format") {
		app.help_home()
		return
	}
	input, erri := getArgs("--input")
	format, errf := getArgs("--format")
	debug, _ := getArgs("--debug")
	if input == "" || erri != nil || errf != nil {
		app.help_home()
		return
	}
	if format == "" {
		format = "json"
	}
	p := parser{input: input, format: format, debug: debug}
	p.run()
}
