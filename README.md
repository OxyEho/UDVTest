# Тестовое задание - библиотека

## Описание

REST API Web-сервис для каталогизации книг (библиотека)

Приложение написано на стеке: Golang (gorilla/mux, gorm), PostgreSQL

В приложении реализовано три эндпоинта

1. `GET /api/book[?page={page_number}&size={size}]` - получение списка книг с использованием пэйджинации
2. `GET /api/book/{id}` - получение информации о книге по id
3. `GET /api/book/{id}/items` - получение информации об экзеплярах книги по id

## Запуск

Приложение работает в Docker

Запустить приложение можно с помощью make `make run`

или с помощью `docker-compose`

В результате будут запущены два контейнера: с приложением и базой данных. 

У базы данных открыт порт `5432` для загрузки тестовых данных.

**Для корректной работы приложения нужно создать файл .env и указать 
все переменные окружения из .env.example**

## Модель данных

#### Таблица с информацией о книгах

| books                                           |   
|-------------------------------------------------|
| id  `int; primary key`                          |
| isbn `int; unique`                              |
| title `text`                                    |
| publisher_id `int; foreing key (publishers id)` |
| edition `int`                                   |
| description `text`                              |
| author `text`                                   |
| year `int`                                      |

#### Таблица с информацией об экземплярах книг

| book_entities                        |
|--------------------------------------|
| id  `int; primary key`               |
| book_id `int; foreing key (book id)` |
| is_taken `bool`                      |
| taker_id `id`                        |
| place `text`                         |

#### Таблица с информацией о пользователях

| users                 |
|-----------------------|
| id `int; primary key` |
| name `text`           |
| surname `text`        |
| patronymic `text`     |


#### Таблица с информацией об издателях

| publishers            |
|-----------------------|
| id `int; primary key` |
| name `text`           |
| addr `text`           |
| code `int; unique`    |

#### Таблица для сохранения истории выдачи экземпляров книг

| book_entity_users                                    |
|------------------------------------------------------|
| book_entity_id `int; foreing key (book_entities id)` |
| user_id `int; foreing key (users id)`                |
