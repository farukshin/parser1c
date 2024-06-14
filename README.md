# parser1c
Парсер логов технологического журнала 1С, в несколько потоков

## Установка

```
git clone https://github.com/farukshin/parser1c.git
cd parser1c
go build .
./parser1c --version
```

## Использование

Справка
```
./parser1c --help
```

Выгрузка в формат JSON
```
./parser1c --input=./example/TJ/ --format=json --countRuner=4 --output=./log.json
```
