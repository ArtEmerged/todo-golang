package repository

import (
	"database/sql"
	"fmt"
	"strings"
	"todo"
)

type TodoItemsSqlite struct {
	db *sql.DB
}

func NewTodoItemSqlite(db *sql.DB) *TodoItemsSqlite {
	return &TodoItemsSqlite{db: db}
}

func (r *TodoItemsSqlite) Create(listId int, input todo.TodoItem) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	query := fmt.Sprintf("INSERT INTO %s (title, description, done) VALUES ($1, $2, $3)", todoItemsTable)
	result, err := tx.Exec(query, input.Title, input.Description, 0)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	itemId, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	createListsItemsQuery := fmt.Sprintf("INSERT INTO %s (item_id, list_id) VALUES ($1, $2)", listsItemsTable)
	_, err = tx.Exec(createListsItemsQuery, itemId, listId)
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	return int(itemId), tx.Commit()
}

func (r *TodoItemsSqlite) GetAll(listId int) ([]todo.TodoItem, error) {
	query := fmt.Sprintf("SELECT ti.id, ti.title, ti.description, ti.done FROM %s ti INNER JOIN %s il ON ti.id = il.item_id WHERE il.list_id = $1", todoItemsTable, listsItemsTable)
	rows, err := r.db.Query(query, listId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []todo.TodoItem
	for rows.Next() {
		var item todo.TodoItem
		err = rows.Scan(&item.Id, &item.Title, &item.Description, &item.Done)
		if err != nil {
			return nil, err
		}
		items = append(items, item)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

func (r *TodoItemsSqlite) GetItemById(userId, itemId int) (todo.TodoItem, error) {
	var item todo.TodoItem
	query := fmt.Sprintf("SELECT ti.id, ti.title, ti.description, ti.done FROM %s ti INNER JOIN %s li ON ti.id = li.item_id INNER JOIN %s ul ON li.list_id = ul.list_id WHERE ti.id = $1 AND ul.user_id = $2",
		todoItemsTable, listsItemsTable, usersListsTable)
	err := r.db.QueryRow(query, itemId, userId).Scan(&item.Id, &item.Title, &item.Description, &item.Done)
	if err != nil {
		return item, err
	}
	return item, nil
}

func (r *TodoItemsSqlite) Delete(itemId int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1", todoItemsTable)
	_, err := r.db.Exec(query, itemId)
	return err
}

func (r *TodoItemsSqlite) Update(itemId int, input todo.UpdateItem) error {
	set := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1
	if input.Title != nil {
		set = append(set, fmt.Sprintf("title = $%d", argId))
		args = append(args, input.Title)
		argId++
	}
	if input.Description != nil {
		set = append(set, fmt.Sprintf("description = $%d", argId))
		args = append(args, input.Description)
		argId++
	}
	if input.Done != nil {
		set = append(set, fmt.Sprintf("done = $%d", argId))
		args = append(args, input.Done)
		argId++
	}
	args = append(args, itemId)
	joinSet := strings.Join(set, ", ")
	query := fmt.Sprintf("UPDATE %s  SET %s WHERE id = $%d", todoItemsTable, joinSet, argId)
	_, err := r.db.Exec(query, args...)
	return err
}
