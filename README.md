# parser_tj_1c
Парсер логов технологического журнала 1С

## Установка

```
git clone https://github.com/farukshin/parser_tj_1c.git
cd parser_tj_1c
go build .
./parser_tj_1c --version
```

## Использование

Справка
```
./parser_tj_1c --help
```

Выгрузка в формат JSON
```
./parser_tj_1c -i="./example/TJ/" -f="json" > log.json
```
