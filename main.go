package main

import (
	"errors"
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

func getTodo(context *gin.Context) {
	id := context.Param("id")
	singleTodo, err := getTodoById(id)
	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Todo not found"})
		return
	}
	context.IndentedJSON(http.StatusOK, singleTodo)

}

// handler
func getTodoById(id string) (*todo, error) {
	for i, t := range todos {
		if t.ID == id {
			return &todos[i], nil
		}
	}
	return nil, errors.New("todo not found")
}

func main() {
	router := gin.Default()
	router.GET("/api/rest/ToDo/todos", getAllTodos)
	router.GET("/api/rest/ToDo/todos/:id", getTodo)
	router.POST("/api/rest/ToDo/todo", addTodo)
	err := router.Run("localhost:7000")
	if err != nil {
		return
	}
}
