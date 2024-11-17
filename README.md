# Парсер 1С

![GitHub Release](https://img.shields.io/github/v/release/farukshin/parser1c)
![GitHub build status](https://github.com/farukshin/parser1c/actions/workflows/parser1c.yml/badge.svg)
![Codecov](https://img.shields.io/codecov/c/github/farukshin/parser1c)
![GitHub Downloads (all assets, all releases)](https://img.shields.io/github/downloads/farukshin/parser1c/total?color=green)
![GitHub License](https://img.shields.io/github/license/farukshin/parser1c)


Многопоточный парсер логов технологического журнала 1С.

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
