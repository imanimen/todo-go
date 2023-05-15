package main

import "github.com/gin-gonic/gin"

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

func main() {
	router := gin.Default()
	err := router.Run("localhost:7000")
	if err != nil {
		return
	}
}
