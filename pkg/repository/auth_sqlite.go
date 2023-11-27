package repository

import (
	"database/sql"
	"fmt"

	"todo"
)

type AuthSqlite struct {
	db *sql.DB
}

func NewAuthSqlite(db *sql.DB) *AuthSqlite {
	return &AuthSqlite{db: db}
}

func (r *AuthSqlite) CreateUser(user todo.User) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (name, username, password_hash) values ($1, $2, $3)", userTable)
	_, err := r.db.Exec(query, user.Name, user.Username, user.Password)
	if err != nil {
		return 0, err
	}
	err = r.db.QueryRow("SELECT last_insert_rowid() AS id").Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (r *AuthSqlite) GetUser(username, password string) (todo.User, error) {
	var user todo.User
	query := fmt.Sprintf("SELECT * FROM %s WHERE username = $1 AND password_hash = $2 ", userTable)
	err := r.db.QueryRow(query, username, password).Scan(&user.Id, &user.Name, &user.Username, &user.Password)
	return user, err
}
