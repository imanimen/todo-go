package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type todo struct {
	ID        string `json:"id"`
	Item      string `json:"item"`
	Completed bool   `json:"completed"`
}

var todos = []todo{
	{ID: "1", Item: "Clean Room", Completed: false},
	{ID: "2", Item: "Do Homework", Completed: true},
	{ID: "3", Item: "Read Book", Completed: true},
}

func getAllTodos(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, todos) // to return json IntendedJSON
}

func addTodo(context *gin.Context) {
	var newTodo todo
	if err := context.BindJSON(&newTodo); err != nil { // to get json BindJSON
		return
	}

	todos = append(todos, newTodo)
	context.IndentedJSON(http.StatusCreated, todos)
}

func main() {
	router := gin.Default()
	router.GET("/api/rest/ToDo/todos", getAllTodos)
	router.POST("/api/rest/ToDo/todo", addTodo)
	err := router.Run("localhost:7000")
	if err != nil {
		return
	}
}
