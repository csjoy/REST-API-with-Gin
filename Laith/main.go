package main

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Todo struct {
	ID        string `json:"id"`
	Item      string `json:"item"`
	Completed bool   `json:"completed"`
}

var todos = []Todo{
	{ID: "1", Item: "Clean Room", Completed: false},
	{ID: "2", Item: "Read Book", Completed: false},
	{ID: "3", Item: "Record Video", Completed: false},
}

func getTodos(ctx *gin.Context) {
	ctx.IndentedJSON(http.StatusOK, todos)
}

func postTodo(ctx *gin.Context) {
	var newTodo Todo
	err := ctx.BindJSON(&newTodo)
	if err != nil {
		return
	}
	todos = append(todos, newTodo)
	ctx.IndentedJSON(http.StatusCreated, newTodo)
}

func getTodo(ctx *gin.Context) {
	id := ctx.Param("id")
	todo, err := getTodoByID(id)
	if err != nil {
		ctx.IndentedJSON(http.StatusNotFound, gin.H{
			"message": "todo not found",
		})
		return
	}
	ctx.IndentedJSON(http.StatusOK, todo)
}

func getTodoByID(id string) (*Todo, error) {
	for i, v := range todos {
		if v.ID == id {
			return &todos[i], nil
		}
	}
	return nil, errors.New("todo not found")
}

func toggleTodoStatus(ctx *gin.Context) {
	id := ctx.Param("id")
	todo, err := getTodoByID(id)
	if err != nil {
		ctx.IndentedJSON(http.StatusNotFound, gin.H{
			"message": "todo not found",
		})
		return
	}

	todo.Completed = !todo.Completed
	ctx.IndentedJSON(http.StatusOK, todo)
}

func main() {
	router := gin.Default()
	router.GET("/todos", getTodos)
	router.GET("/todos/:id", getTodo)
	router.PATCH("/todos/:id", toggleTodoStatus)
	router.POST("/todos", postTodo)
	router.Run("localhost:8080")
}
