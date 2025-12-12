package main

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	_ "github.com/lib/pq"
)

// setupTestDB устанавливает глобальную переменную db для тестов и очищает таблицу users
func setupTestDB(t *testing.T) {
	var err error
	db, err = sql.Open("postgres", "host=localhost port=5432 user=admin password=admin dbname=laba8go sslmode=disable")
	if err != nil {
		t.Fatalf("Ошибка подключения к БД: %v", err)
	}
	if err = db.Ping(); err != nil {
		t.Fatalf("БД недоступна: %v", err)
	}

	// Создадим таблицу, если её нет, и очистим данные
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS users (
        id serial PRIMARY KEY,
        name text NOT NULL,
        age int NOT NULL
    )`)
	if err != nil {
		t.Fatalf("Не удалось создать таблицу users: %v", err)
	}

	_, err = db.Exec("TRUNCATE TABLE users RESTART IDENTITY")
	if err != nil {
		t.Fatalf("Не удалось очистить таблицу users: %v", err)
	}
}

func TestCreateUser(t *testing.T) {
	setupTestDB(t)

	user := User{Name: "Иван", Age: 30}
	body, _ := json.Marshal(user)

	req := httptest.NewRequest(http.MethodPost, "/users", bytes.NewBuffer(body))
	rec := httptest.NewRecorder()

	usersHandler(rec, req)

	if rec.Code != http.StatusCreated {
		t.Fatalf("Ожидали 201, получили %d", rec.Code)
	}

	var created User
	json.Unmarshal(rec.Body.Bytes(), &created)

	if created.ID == 0 || created.Name != "Иван" || created.Age != 30 {
		t.Errorf("Некорректные данные пользователя: %+v", created)
	}
}

func TestGetUsers(t *testing.T) {
	setupTestDB(t)

	// Добавим тестовые данные
	db.Exec("INSERT INTO users (name, age) VALUES ('Alice', 25), ('Bob', 40)")

	req := httptest.NewRequest(http.MethodGet, "/users", nil)
	rec := httptest.NewRecorder()

	usersHandler(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("Ожидали 200, получили %d", rec.Code)
	}

	var users []User
	json.Unmarshal(rec.Body.Bytes(), &users)

	if len(users) < 2 {
		t.Errorf("Ожидали >= 2 пользователей, получили %d", len(users))
	}
}

func TestGetUserByID_NotFound(t *testing.T) {
	setupTestDB(t)

	req := httptest.NewRequest(http.MethodGet, "/users/9999", nil)
	rec := httptest.NewRecorder()

	userHandler(rec, req)

	if rec.Code != http.StatusNotFound {
		t.Fatalf("Ожидали 404, получили %d", rec.Code)
	}
}

func TestUpdateUser(t *testing.T) {
	setupTestDB(t)

	var id int
	db.QueryRow("INSERT INTO users (name, age) VALUES ('Old', 20) RETURNING id").Scan(&id)

	updated := User{Name: "New", Age: 35}
	body, _ := json.Marshal(updated)

	req := httptest.NewRequest(http.MethodPut, "/users/"+strconv.Itoa(id), bytes.NewBuffer(body))
	rec := httptest.NewRecorder()

	userHandler(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("Ожидали 200, получили %d", rec.Code)
	}

	var u User
	json.Unmarshal(rec.Body.Bytes(), &u)
	if u.Name != "New" || u.Age != 35 {
		t.Errorf("Обновление не сработало: %+v", u)
	}
}

func TestDeleteUser(t *testing.T) {
	setupTestDB(t)

	var id int
	db.QueryRow("INSERT INTO users (name, age) VALUES ('Del', 50) RETURNING id").Scan(&id)

	req := httptest.NewRequest(http.MethodDelete, "/users/"+strconv.Itoa(id), nil)
	rec := httptest.NewRecorder()

	userHandler(rec, req)

	if rec.Code != http.StatusNoContent {
		t.Fatalf("Ожидали 204, получили %d", rec.Code)
	}

	// Проверим, что действительно удален
	var count int
	db.QueryRow("SELECT COUNT(*) FROM users WHERE id=$1", id).Scan(&count)
	if count != 0 {
		t.Errorf("Пользователь не удален из БД")
	}
}
