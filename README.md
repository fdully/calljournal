# ТЗ - сервис "Журнал звонков"
[![Build Status](https://travis-ci.com/fdully/calljournal.svg?branch=master)](https://travis-ci.com/fdully/calljournal)

## Общее описание
Сервис предназначен для упорядоченного хранения аудио записей и мета информации о телефонных звонках.
А также для быстрого и удобного поиска аудио записей по номеру телефона и дате через API. 

Данные о звонках и аудио записи поступают из телефонной станции в json и wav форматах.

Пользовательский интерфейс не предусмотрен.

## Описание сущностей
### Звонок
* Данные звонка

### Аудио файл
* Данные аудиофайла


## Архитектура
Клиент/Сервер.
Клиент запущен на телефонном сервере и мониторит дериктории с аудио записями и метаданными по звонкам.
Клиентов может быть несколько, на каждом телефонном сервере по одному клиенту.
Сервер хранения звонков один. Клиенты вычитывают json файлы и аудио записи на телефонном сервере и отправляют
с помощью GRPC на сервер. Если сервер не доступен, то клиенты ждут пока восстановится связь и временно
прекращают вычитывать данные о звонках.
Микросервис состоит из API, базы для хранения данных о звонках и хранилища аудио записей
(можно использовать S3 интерфейс для хранения аудио файлов или складывать в папки самостоятельно).
Сервис должен предоставлять GRPC и REST API.

## Описание методов API

#### Добавление данных о звонке в базу
Запрос:
* Звонок

Ответ:
* ok

#### Загрузка аудио записи в хранилище
Запрос:
* Аудио файл

Ответ:
* ok

#### Поиск всех звонков по номеру телефона
#####TODO
Запрос:
* Номер телефона

Ответ:
* Список звонков

#### Поиск всех звонков за интервал времени
#####TODO
Запрос:
* Начало и конец временного интервала

Ответ:
* Список звонков

#### Скачивание файла записи
#####TODO
Запрещено скачивать файлы с тегом private

Запрос:
* ID звонка

Ответ:
* Аудио файл


## Конфигурация
Основные параметры конфигурации: хранилище для аудио записей, база данных по звонкам,
директории с json и wav файлами на сервере телефонии.

## Развертывание
Развертывание микросервиса должно осуществляться командой `make run` (внутри `docker compose up`)
в директории с проектом.

## Тестирование
Так же необходимо написать интеграционные тесты, проверяющие все вызовы API.
