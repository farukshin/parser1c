# parser_tj_1c
Парсер логов технологического журнала 1С, в несколько потоков

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
./parser_tj_1c --input=./example/TJ/rphost_160/24051511.log --format=json --countRuner=4 --output=./log.json
```
