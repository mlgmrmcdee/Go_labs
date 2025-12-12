## Документация API — laba8

Базовый URL: http://localhost:8080

Контент-тип: все ответы с JSON-телом имеют заголовок `Content-Type: application/json`.

Общие правила валидации (сервер):

- `name` — непустая строка, максимум 100 символов
- `age` — целое > 0 и <= 120

---

### GET /users

Описание: получить список пользователей с поддержкой фильтрации и пагинации.

Параметры query (все опциональны):

- `name` — строка. Поиск по подстроке (регистронезависимо, используется `ILIKE '%name%'`).
- `min_age` — целое. Включающий минимальный возраст (age >= min_age).
- `max_age` — целое. Включающий максимальный возраст (age <= max_age).
- `page` — целое, страница (по умолчанию 1). Нумерация с 1.
- `limit` — целое, количество элементов на странице (по умолчанию 10, максимум 100).

Пример запросов (PowerShell / curl):

```powershell
curl.exe "http://localhost:8080/users?page=1&limit=5"
curl.exe "http://localhost:8080/users?name=ivan&min_age=20&max_age=30"
```

Пример ответа (200 OK):

```json
[
  { "id": 1, "name": "Ivan", "age": 25 },
  { "id": 2, "name": "Anna", "age": 30 }
]
```

Ошибки:

- 400 Bad Request — неверные параметры (например, `page=abc`). Ответ JSON: `{ "error": "..." }`.
- 500 Internal Server Error — ошибка на сервере/БД.

---

### POST /users

Описание: создать нового пользователя.

Заголовки:

- `Content-Type: application/json`

Тело (JSON):

```json
{
  "name": "Иван",
  "age": 30
}
```

Пример (PowerShell):

```powershell
curl.exe -X POST "http://localhost:8080/users" -H "Content-Type: application/json" -d '{"name":"Иван","age":30}'
```

Успешный ответ (201 Created):

```json
{
  "id": 5,
  "name": "Иван",
  "age": 30
}
```

Ошибки:

- 400 Bad Request — некорректный JSON или валидация полей (возраст/имя).
- 500 Internal Server Error — ошибка БД.

---

### GET /users/{id}

Описание: получить пользователя по ID.

Параметры пути: `id` — целое.

Пример:

```powershell
curl.exe "http://localhost:8080/users/3"
```

Успешный ответ (200 OK):

```json
{ "id": 3, "name": "Petr", "age": 40 }
```

Ошибки:

- 400 Bad Request — неверный ID в пути.
- 404 Not Found — пользователь с таким ID не найден.
- 500 Internal Server Error — ошибка БД.

---

### PUT /users/{id}

Описание: обновить существующего пользователя.

Заголовки: `Content-Type: application/json`

Тело (JSON):

```json
{ "name": "Новый", "age": 35 }
```

Пример (PowerShell):

```powershell
curl.exe -X PUT "http://localhost:8080/users/3" -H "Content-Type: application/json" -d '{"name":"Новый","age":35}'
```

Успешный ответ (200 OK): обновлённый объект пользователя.

Ошибки:

- 400 Bad Request — неверный JSON/валидация/ID.
- 404 Not Found — пользователь не найден.
- 500 Internal Server Error — ошибка БД.

---

### DELETE /users/{id}

Описание: удалить пользователя по ID.

Пример:

```powershell
curl.exe -X DELETE "http://localhost:8080/users/3"
```

Успешный ответ: 204 No Content (тело отсутствует).

Ошибки:

- 400 Bad Request — неверный ID.
- 404 Not Found — пользователь не найден.
- 500 Internal Server Error — ошибка БД.

---

Примечания по отладке и тестированию

- UI: откройте `http://localhost:8080` — там есть простая тест-страница с кнопками и панелью фильтров/пагинации.
- Логи: сервер печатает сообщения при старте; в случае ошибок операции над БД возвращают JSON с полем `error`.
- Тесты: в репозитории добавлен файл `handlers_test.go` (юнит/интеграционные тесты), они ожидают доступной PostgreSQL с учётной записью `admin`/`admin` и базой `laba8go`. Тесты автоматически создают таблицу `users` и очищают её перед запуском.

Как запустить сервер:

```powershell
cd 'c:\Users\gadzh\OneDrive\Документы\GitHub\Go_labs\laba8'
go run server.go
```

Как запустить тесты:

```powershell
cd 'c:\Users\gadzh\OneDrive\Документы\GitHub\Go_labs\laba8'
go test -v
```
