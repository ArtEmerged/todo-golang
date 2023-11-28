package handler

import (
	"net/http"
	"strconv"

	"todo"

	"github.com/gin-gonic/gin"
)

func (h *Handler) createList(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "user id not found") // 500
		return
	}
	var input todo.TodoList
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error()) // 400
		return
	}
	id, err := h.services.TodoList.Create(userId, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error()) // 500
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

type getAllListsResponse struct {
	Data []todo.TodoList `json:"data"`
}

func (h *Handler) getAllLists(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "user id not found") // 500
		return
	}

	lists, err := h.services.TodoList.GetAll(userId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error()) // 500
		return
	}
	c.JSON(http.StatusOK, getAllListsResponse{Data: lists})
}

func (h *Handler) getListById(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "user id not found") // 500
		return
	}
	listId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param") // 400
		return
	}

	list, err := h.services.TodoList.GetListById(userId, listId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error()) // 500
		return
	}
	c.JSON(http.StatusOK, list)
}

func (h *Handler) updateList(c *gin.Context) {
}

func (h *Handler) deleteList(c *gin.Context) {
}
