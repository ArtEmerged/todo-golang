# REST API Для Создания TODO Lists на Go


## В проекте разобранны следующие концепции:
- Разработка Веб-Приложений на Go, следуя дизайну REST API.
- Работа с фреймворком <a href="https://github.com/gin-gonic/gin">gin-gonic/gin</a>.
- Подход Чистой Архитектуры в построении структуры приложения. Техника внедрения зависимости.
- Работа с БД SQLite3. Генерация файлов миграций. 
- Конфигурация приложения с помощью библиотеки <a href="https://github.com/spf13/viper">spf13/viper</a>. Работа с переменными окружения.
- Регистрация и аутентификация. Работа с JWT. Middleware.
- Написание SQL запросов.
- Graceful Shutdown.

### Для запуска приложения:

```
make build && make run
```

Если приложение запускается впервые, необходимо применить миграции к базе данных:

```
make migrate
```

## Routing Requests



| HTTP Method | URL Pattern          | Handler        | Middleware        | Action                                 |
|-------------|----------------------|----------------|-------------------|----------------------------------------|
| POST        | /auth/sign-up        | signUp         | ----              | User registration                      |
| POST        | /auth/sign-in        | signIn         | ----              | User login                             |
|-------------|----------------------|----------------|-------------------|----------------------------------------|
| POST        | /api/lists/          | createList     | userIdentity      | Create a new list                      |
| GET         | /api/lists/          | getAllLists    | userIdentity      | Get all lists                          |
| GET         | /api/lists/:id       | getListById    | userIdentity      | Get a list by ID                       |
| PUT         | /api/lists/:id       | updateList     | userIdentity      | Update a list by ID                    |
| DELETE      | /api/lists/:id       | deleteList     | userIdentity      | Delete a list by ID                    |
| POST        | /api/lists/:id/items | createItem     | userIdentity      | Create a new item in a list by ID      |
| GET         | /api/lists/:id/items | getAllItems    | userIdentity      | Get all items in a list by ID          |
|-------------|----------------------|----------------|-------------------|----------------------------------------|
| GET         | /api/items/:id       | getItemById    | userIdentity      | Get an item by ID                      |
| PUT         | /api/items/:id       | updateItem     | userIdentity      | Update an item by ID                   |
| DELETE      | /api/items/:id       | deleteItem     | userIdentity      | Delete an item by ID                   |

