package repository

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"todo"
)

type TodoListSqlite struct {
	db *sql.DB
}

func NewTodoListSqlite(db *sql.DB) *TodoListSqlite {
	return &TodoListSqlite{db: db}
}

func (r *TodoListSqlite) Create(userId int, list todo.TodoList) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}
	var listId int
	createListQuery := fmt.Sprintf("INSERT INTO %s (title, description) VALUES ($1, $2)", todoListTable)
	_, err = r.db.Exec(createListQuery, list.Title, list.Description)
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	r.db.QueryRow("SELECT last_insert_rowid() AS id").Scan(&listId)
	createUsersListsQuery := fmt.Sprintf("INSERT INTO %s (user_id, list_id) VALUES ($1, $2)", usersListsTable)
	_, err = r.db.Exec(createUsersListsQuery, userId, listId)
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	return listId, tx.Commit()
}

func (r *TodoListSqlite) GetAll(userId int) ([]todo.TodoList, error) {
	query := fmt.Sprintf("SELECT tl.id, tl.title, tl.description FROM %s tl INNER JOIN %s ul on tl.id = ul.list_id WHERE ul.user_id = $1 ", todoListTable, usersListsTable)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	rows, err := r.db.QueryContext(ctx, query, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var lists []todo.TodoList
	for rows.Next() {
		list := todo.TodoList{}
		err := rows.Scan(&list.Id, &list.Title, &list.Description)
		if err != nil {
			return nil, err
		}
		lists = append(lists, list)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return lists, nil
}
