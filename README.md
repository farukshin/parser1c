# Парсер 1С

[![GitHub Release](https://img.shields.io/github/v/release/farukshin/parser1c?color=%231E90FF%09)](https://github.com/farukshin/parser1c/releases)
![GitHub build status](https://github.com/farukshin/parser1c/actions/workflows/parser1c.yml/badge.svg)
![Codecov](https://img.shields.io/codecov/c/github/farukshin/parser1c)
![GitHub Downloads (all assets, all releases)](https://img.shields.io/github/downloads/farukshin/parser1c/total?color=green)
[![GitHub License](https://img.shields.io/github/license/farukshin/parser1c)](https://github.com/farukshin/parser1c/blob/main/LICENSE.md)


Многопоточный парсер логов технологического журнала 1С с выгрузкой в PostgreSQL.

* [Установка](#install)
* * [Установка из релизов](#installRelease)
* * [Установка из исходников](#installSource)
* [Использование](#usage)
* [Анализ логов технологического журнала в SQL](#sql)
* [Лицензия](#lic)


<a name="install"></a> 

## Установка

<a name="installSource"></a> 

### Установка из исходников

```
git clone https://github.com/farukshin/parser1c.git
cd parser1c
go build .
./parser1c --version
```

<a name="installRelease"></a> 

### Установка из релизов

1. Получить версию [последнего релиза](https://github.com/farukshin/parser1c/releases).

``` bash
VERSION=$(curl -s "https://api.github.com/repos/farukshin/parser1c/releases/latest" | jq -r '.tag_name')
```
Или установить необходимую версию релиза:

``` bash
VERSION=vX.Y.Z
```

2. Загрузка релиза

``` bash
OS=Linux       # or Darwin, Windows
ARCH=x86_64    # or arm64, x86_64, armv6, i386, s390x
FILE=parser1c_${OS}_${ARCH}.tar.gz
curl -sL "https://github.com/farukshin/parser1c/releases/download/${VERSION}/${FILE}" > ${FILE}
```

3. Проверка контрольной суммы

``` bash
curl -sL https://github.com/farukshin/parser1c/releases/download/${VERSION}/parser1c_checksums.txt > parser1c_checksums.txt
shasum --check --ignore-missing ./parser1c_checksums.txt
```

4. Распаковать парсер

``` bash
tar -zxvf ${FILE} parser1c
./parser1c --version
```

<a name="usage"></a> 

## Использование

```
Строка запуска: parser1c [Опции]

Опции:
-h --help - вызов справки
-v --version - версия приложения
--input - каталог с логами технологического журнала или имя файла с логами
--output - приемник (на данный момент только postgres)
--host - хост PostgreSQL (либо env PG_HOST)
--port - порт PostgreSQL (либо env PG_PORT)
--user - пользователь PostgreSQL (либо env PG_USER)
--password - пароль PostgreSQL (либо env PG_PASSWORD)
--dbname - база данных PostgreSQL (либо env PG_DBNAME)
--countRuner - количество потоков парсера, по умолчанию 1
```

Допустим, в настройках сбора технологического журнала каталог сбора логв указан `/var/log/1c`:

```
<?xml version="1.0"?>
<config xmlns="http://v8.1c.ru/v8/tech-log">
    <log location="/var/log/1c" history="8">
    ...
```
И логи планируем хранить в PostgreSQL, запущенному по адресу `localhost` на порту `5432`, в базе данных `alsu`, пользователь и пароль `postgres`

Пример запуска парсера:

```
./parser1c --input=/var/log/1c --output=postgres --host=localhost --port=5432 --user=postgres --password=postgres --countrunner=4 --dbname=alsu
```

Параметры подключения к PostgreSQL можно указать в переменных окружения. 

``` bash
PG_HOST=localhost
PG_PORT=5432
PG_USER=postgres
PG_PASSWORD=postgres
PG_DBNAME=alsu
```

Тогда строка запуска парсера будут:
```
./parser1c --input=/var/log/1c --output=postgres --countrunner=8
```

<a name="sql"></a> 

## Анализ логов технологического журнала в SQL

Для просмотре и анализа логов можно подключиться к базе любым удобным способом, например используя [DBeaver](https://dbeaver.io/)

![](./static/dbeaver.png)


<a name="lic"></a> 

## Лицензия

Parser1c выпускается под лицензией MIT. Подробнее [LICENSE.md](https://github.com/farukshin/parser1c/blob/main/LICENSE.md)
