# SberToDo

Simple CRUDL ToDo API made with Postgres, Gorm, Docker, Swagger, and Fiber.

- #### Для запуска:

Добавить файл .env и заполнить его в соответствии:

```
DB_HOST= your_host
DB_USER= your_user
DB_PORT= your_port
DB_PASS= your_password
DB_NAME= db_name
```

Запуск:

```
$ docker compose up
```

- #### API

Примеры запросов через программу Postman:

##### POST

##### Добавление новой задачи:

Запрос на адрес localhost:80/add

```json
{
    "title": "ToDel",
    "desc": "deleted",
    "date": "01.01.2024",
    "flag": true
}
```

Ответ:

```json
{
    "message": "todo created",
    "todo": {
        "title": "ToDel",
        "desc": "deleted",
        "date": "2024-01-01T00:00:00Z",
        "flag": true
    }
}
```

##### GET

##### Просмотр всех задач:

Запрос на адрес localhost:80/show

Ответ:

```json
[
    {
        "title": "Uni",
        "desc": "Get a diploma",
        "date": "2024-05-28T00:00:00Z",
        "flag": false
    },
    {
        "title": "Animal",
        "desc": "Pet a cat",
        "date": "2023-08-20T00:00:00Z",
        "flag": true
    },
    {
        "title": "Gym",
        "desc": "Get pumped up",
        "date": "2023-10-05T00:00:00Z",
        "flag": true
    },
    {
        "title": "Home",
        "desc": "Visit siblings",
        "date": "2023-11-11T00:00:00Z",
        "flag": false
    },
    {
        "title": "Language",
        "desc": "Learn Spanish",
        "date": "2023-12-19T00:00:00Z",
        "flag": true
    },
    {
        "title": "Shop",
        "desc": "Buy new clothes",
        "date": "2023-09-30T00:00:00Z",
        "flag": false
    },
    {
        "title": "Email",
        "desc": "Email team for updates",
        "date": "2023-10-14T00:00:00Z",
        "flag": true
    },
    {
        "title": "BD",
        "desc": "Visit friend's party",
        "date": "2023-09-14T00:00:00Z",
        "flag": true
    },
    {
        "title": "Clean",
        "desc": "Clean room",
        "date": "2023-09-22T00:00:00Z",
        "flag": false
    },
    {
        "title": "Phone",
        "desc": "Buy new phone",
        "date": "2024-01-01T00:00:00Z",
        "flag": false
    },
    {
        "title": "Job",
        "desc": "Get a job",
        "date": "2023-09-30T00:00:00Z",
        "flag": false
    },
    {
        "title": "ToDel",
        "desc": "deleted",
        "date": "2024-01-01T00:00:00Z",
        "flag": true
    }
]
```

##### Обновление статуса задачи по заголовку:

Запрос на адрес localhost:80/flag/ToDel

Ответ:

```json
{
    "title": "ToDel",
    "desc": "deleted",
    "date": "2024-01-01T00:00:00Z",
    "flag": false
}
```

##### Пагинация со статусом выполнено / не выполнено (true/false):

Запрос на адрес localhost:80/true&page=1

Ответ:

```json
[
    {
        "title": "Animal",
        "desc": "Pet a cat",
        "date": "2023-08-20T00:00:00Z",
        "flag": true
    },
    {
        "title": "Gym",
        "desc": "Get pumped up",
        "date": "2023-10-05T00:00:00Z",
        "flag": true
    },
    {
        "title": "Language",
        "desc": "Learn Spanish",
        "date": "2023-12-19T00:00:00Z",
        "flag": true
    }
]
```

Запрос на адрес localhost:80/false&page=3

Ответ:

```json
[
    {
        "title": "ToDel",
        "desc": "deleted",
        "date": "2024-01-01T00:00:00Z",
        "flag": false
    }
]
```

##### Просмотр по дате со статусом true/false

Запрос на адрес localhost:80/date/false

Ответ:

```json
[
    {
        "title": "Clean",
        "desc": "Clean room",
        "date": "2023-09-22T00:00:00Z",
        "flag": false
    },
    {
        "title": "Shop",
        "desc": "Buy new clothes",
        "date": "2023-09-30T00:00:00Z",
        "flag": false
    },
    {
        "title": "Job",
        "desc": "Get a job",
        "date": "2023-09-30T00:00:00Z",
        "flag": false
    },
    {
        "title": "Home",
        "desc": "Visit siblings",
        "date": "2023-11-11T00:00:00Z",
        "flag": false
    },
    {
        "title": "Phone",
        "desc": "Buy new phone",
        "date": "2024-01-01T00:00:00Z",
        "flag": false
    },
    {
        "title": "ToDel",
        "desc": "deleted",
        "date": "2024-01-01T00:00:00Z",
        "flag": false
    },
    {
        "title": "Uni",
        "desc": "Get a diploma",
        "date": "2024-05-28T00:00:00Z",
        "flag": false
    }
]
```

##### PATCH

##### Обновление задачи по заголовку:

Запрос на адрес localhost:80/change/ToDel

```json
{
    "desc": "patched",
    "date": "02.02.2010"
}
```

Ответ:

```json
{
    "title": "ToDel",
    "desc": "patched",
    "date": "2010-02-02T00:00:00Z",
    "flag": false
}
```

##### DELETE

##### Удаление задачи по заголовку:

Запрос на адрес localhost:80/del/ToDel

Ответ:

```json
{
    "Message": "todo deleted successfully",
    "Todo": {
        "title": "ToDel",
        "desc": "patched",
        "date": "2010-02-02T00:00:00Z",
        "flag": false
    }
}
```

##### SWAGGER

Запрос на адрес localhost:80/swagger/index.html

![image](https://github.com/losevs/SberToDo/assets/75357413/8bb28a9a-8492-4f24-9c67-3b0d6b864a15)
