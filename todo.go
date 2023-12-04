package todo

import "errors"

type TodoList struct {
	Id          int    `json:"id" db: "id"`
	Title       string `json:"title" db: "title" binding:"required"`
	Description string `json:"description" db: "description"`
}

type UsersList struct {
	Id     int
	UserId int
	ListId int
}

type TodoItem struct {
	Id          int    `json:"id" db:"id"`
	Title       string `json:"title" db: "title" binding:"required"`
	Description string `json:"description" db:"descrition"`
	Done        bool   `json:"done" db:"done"`
}

type ListsItem struct {
	Id     int
	ListId int
	ItemId int
}

type UpdateList struct {
	Title       *string `json: "title"`
	Description *string `json:"description"`
}

func (i UpdateList) Validate() error {
	if i.Title == nil && i.Description == nil {
		return errors.New("update structure has no values")
	}
	return nil
}



type UpdateItem struct {
	Title       *string `json: "title"`
	Description *string `json:"description"`
	Done        *bool   `json:"done"`
}

func (i UpdateItem) Validate() error {
	if i.Title == nil && i.Description == nil && i.Done == nil {
		return errors.New("update structure has no values")
	}
	return nil
}
