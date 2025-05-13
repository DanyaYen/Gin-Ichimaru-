package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"strconv"
)

type Todo struct {
	ID int `json:"id"`
	Title string `json:"title"`
	Completed bool `json:"completed"`
}

var (
	todos = map[int]Todo{}
	nextID = 1
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/todos", func(c *gin.Context) {
		list := make([]Todo, 0, len(todos))
		for _, t := range todos {
			list = append(list, t)
		}
		c.JSON(http.StatusOK, list)
	})

	r.POST("/todos", func(c *gin.Context) {
		var input Todo
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error":"invalid input"})
			return
		}		  
		id := nextID
		nextID++
		todo := Todo{ID: id, Title: input.Title, Completed: input.Completed}
		todos[id] = todo
		c.JSON(http.StatusCreated, todo)
	})

	r.GET("/todos/:id", func(c *gin.Context) {
		idSTR := c.Param("id")
		id, err := strconv.Atoi(idSTR)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
			return
		}
		todo, ok := todos[id]
		if !ok {
			c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
			return
		}
		c.JSON(http.StatusOK, todo)
	})

	r.PUT("/todos/:id", func(c *gin.Context) {
		idSTR := c.Param("id")
		id, err := strconv.Atoi(idSTR)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
			return
		}
		todo, ok := todos[id]
		if !ok {
			c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
			return
		}

		var input Todo
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error":"invalid input"})
			return
		}		  
		todo.Title = input.Title
		todo.Completed = input.Completed
		todos[id] = todo

		c.JSON(http.StatusOK, todo)
	})	
	
	r.DELETE("/todos/:id", func(c *gin.Context) {
		idSTR := c.Param("id")
		id, err := strconv.Atoi(idSTR)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
			return
		}
		_, ok := todos[id]
		if !ok {
			c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
			return
		}
		delete(todos, id)
		c.Status(http.StatusNoContent)

	})
	return r
}

func main() {
	r := setupRouter()
	r.Run()
}