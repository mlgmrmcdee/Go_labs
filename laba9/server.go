package main

import (
	"crypto/rand"
	"database/sql"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	_ "github.com/lib/pq"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type AuthUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

var db *sql.DB
var sessions = make(map[string]string) // token -> username

func respondError(w http.ResponseWriter, status int, msg string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(map[string]string{"error": msg})
}

func validateUser(u *User) error {
	u.Name = strings.TrimSpace(u.Name)
	if u.Name == "" {
		return errors.New("имя не может быть пустым")
	}
	if len(u.Name) > 100 {
		return errors.New("максимум 100 символов")
	}
	if u.Age <= 0 {
		return errors.New("возраст должен быть больше 0")
	}
	if u.Age > 120 {
		return errors.New("максимум 120")
	}
	return nil
}

func generateToken() string {
	b := make([]byte, 16)
	rand.Read(b)
	return hex.EncodeToString(b)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		respondError(w, http.StatusMethodNotAllowed, "Только POST")
		return
	}

	var creds AuthUser
	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		respondError(w, http.StatusBadRequest, "Некорректный JSON")
		return
	}

	var storedPass string
	err := db.QueryRow("SELECT password FROM auth_users WHERE username=$1", creds.Username).Scan(&storedPass)
	if err == sql.ErrNoRows || storedPass != creds.Password {
		respondError(w, http.StatusUnauthorized, "Неверное имя пользователя или пароль")
		return
	} else if err != nil {
		respondError(w, 500, err.Error())
		return
	}

	token := generateToken()
	sessions[token] = creds.Username

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"token": token})
}

func logoutHandler(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("Authorization")
	if token == "" {
		respondError(w, http.StatusUnauthorized, "Нет токена")
		return
	}
	delete(sessions, token)
	w.WriteHeader(http.StatusNoContent)
}

func requireAuth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		if token == "" || sessions[token] == "" {
			respondError(w, http.StatusUnauthorized, "Необходима авторизация")
			return
		}
		next(w, r)
	}
}

func main() {
	var err error
	db, err = sql.Open("postgres", "host=localhost port=5432 user=admin password=admin dbname=laba8go sslmode=disable")
	if err != nil {
		log.Fatal("Ошибка подключения:", err)
	}
	if err = db.Ping(); err != nil {
		log.Fatal("База недоступна:", err)
	}
	log.Println("Успешно подключено к PostgreSQL")

	http.Handle("/", http.FileServer(http.Dir("./static")))
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/logout", logoutHandler)
	http.HandleFunc("/users", requireAuth(usersHandler))
	http.HandleFunc("/users/", requireAuth(userHandler))

	log.Println("Сервер запущен: http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

// --- Handlers ---

func usersHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getUsers(w, r)
	case http.MethodPost:
		createUser(w, r)
	default:
		respondError(w, http.StatusMethodNotAllowed, "Метод не поддерживается")
	}
}

func userHandler(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/users/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		respondError(w, http.StatusBadRequest, "Неверный ID")
		return
	}

	switch r.Method {
	case http.MethodGet:
		getUserByID(w, id)
	case http.MethodPut:
		updateUser(w, r, id)
	case http.MethodDelete:
		deleteUser(w, id)
	default:
		respondError(w, http.StatusMethodNotAllowed, "Метод не поддерживается")
	}
}

// --- DB operations ---

func getUsers(w http.ResponseWriter, r *http.Request) {
	// Query params: name (partial), min_age, max_age, page, limit
	q := r.URL.Query()
	name := strings.TrimSpace(q.Get("name"))

	var (
		minAge, maxAge int
		page, limit    int
		err            error
	)

	if s := q.Get("min_age"); s != "" {
		if minAge, err = strconv.Atoi(s); err != nil {
			respondError(w, http.StatusBadRequest, "min_age must be an integer")
			return
		}
	}
	if s := q.Get("max_age"); s != "" {
		if maxAge, err = strconv.Atoi(s); err != nil {
			respondError(w, http.StatusBadRequest, "max_age must be an integer")
			return
		}
	}

	// pagination
	page = 1
	limit = 10
	if s := q.Get("page"); s != "" {
		if page, err = strconv.Atoi(s); err != nil || page < 1 {
			respondError(w, http.StatusBadRequest, "page must be a positive integer")
			return
		}
	}
	if s := q.Get("limit"); s != "" {
		if limit, err = strconv.Atoi(s); err != nil || limit < 1 {
			respondError(w, http.StatusBadRequest, "limit must be a positive integer")
			return
		}
	}
	if limit > 100 {
		limit = 100
	}

	// Build query
	conditions := make([]string, 0)
	args := make([]interface{}, 0)
	idx := 1

	if name != "" {
		conditions = append(conditions, fmt.Sprintf("name ILIKE $%d", idx))
		args = append(args, "%"+name+"%")
		idx++
	}
	if q.Get("min_age") != "" {
		conditions = append(conditions, fmt.Sprintf("age >= $%d", idx))
		args = append(args, minAge)
		idx++
	}
	if q.Get("max_age") != "" {
		conditions = append(conditions, fmt.Sprintf("age <= $%d", idx))
		args = append(args, maxAge)
		idx++
	}

	where := ""
	if len(conditions) > 0 {
		where = "WHERE " + strings.Join(conditions, " AND ")
	}

	offset := (page - 1) * limit
	// add limit and offset
	query := fmt.Sprintf("SELECT id, name, age FROM users %s ORDER BY id LIMIT $%d OFFSET $%d", where, idx, idx+1)
	args = append(args, limit, offset)

	rows, err := db.Query(query, args...)
	if err != nil {
		respondError(w, 500, err.Error())
		return
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var u User
		if err := rows.Scan(&u.ID, &u.Name, &u.Age); err != nil {
			respondError(w, 500, err.Error())
			return
		}
		users = append(users, u)
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(users); err != nil {
		respondError(w, 500, err.Error())
		return
	}
}

func getUserByID(w http.ResponseWriter, id int) {
	var u User
	err := db.QueryRow("SELECT id, name, age FROM users WHERE id=$1", id).Scan(&u.ID, &u.Name, &u.Age)

	if err == sql.ErrNoRows {
		respondError(w, http.StatusNotFound, "Пользователь не найден")
		return
	} else if err != nil {
		respondError(w, 500, err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(u)
}

func createUser(w http.ResponseWriter, r *http.Request) {
	var u User
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		respondError(w, http.StatusBadRequest, "Некорректный JSON")
		return
	}

	if err := validateUser(&u); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}

	err := db.QueryRow("INSERT INTO users (name, age) VALUES ($1, $2) RETURNING id", u.Name, u.Age).Scan(&u.ID)
	if err != nil {
		respondError(w, 500, err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(u)
}

func updateUser(w http.ResponseWriter, r *http.Request, id int) {
	var u User
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		respondError(w, http.StatusBadRequest, "Некорректный JSON")
		return
	}

	if err := validateUser(&u); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}

	result, err := db.Exec("UPDATE users SET name=$1, age=$2 WHERE id=$3", u.Name, u.Age, id)
	if err != nil {
		respondError(w, 500, err.Error())
		return
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		respondError(w, http.StatusNotFound, "Пользователь не найден")
		return
	}

	u.ID = id
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(u)
}

func deleteUser(w http.ResponseWriter, id int) {
	result, err := db.Exec("DELETE FROM users WHERE id=$1", id)
	if err != nil {
		respondError(w, 500, err.Error())
		return
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		respondError(w, http.StatusNotFound, "Пользователь не найден")
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
