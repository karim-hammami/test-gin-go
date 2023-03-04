package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func getTodos(c *gin.Context) {
	c.JSON(http.StatusOK, todos)
}

func createTodos(c *gin.Context) {
	var todo Todo

	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	todos = append(todos, todo)

	c.JSON(http.StatusCreated, todo)
}
