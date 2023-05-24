package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type todo struct {
	ID        string `json:"id"`
	Item      string `json:"item"`
	Completed bool   `json:"completed"`
}

var todos = []todo{
	{"1", "Buy milk", false},
	{"2", "Buy eggs", false},
	{"3", "Buy bread", false},
	{"4", "Buy milk", false},
	{"5", "Buy eggs", false},
}

func getTodos(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, todos)
}
func addTodos(context *gin.Context) {
	var newTodo todo
	if err := context.BindJSON(&newTodo); err != nil {
		return
	}
	todos = append(todos, newTodo)
	context.IndentedJSON(http.StatusCreated, newTodo)
}

func getTodosByid(id string) (*todo, error) {
	for i, t := range todos {
		if t.ID == id {
			return &todos[i], nil
		}
	}
	return nil, errors.New("todo not found")
}
func getTodo(context *gin.Context) {
	id := context.Param("id")
	todo, err := getTodosByid(id)
	if err != nil {

		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "todo not found"})
		return
	}

	context.IndentedJSON(http.StatusOK, todo)

}
func updateTodoStatus(context *gin.Context) {

	id := context.Param("id")
	todo, err := getTodosByid(id)
	if err != nil {

		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "todo not found"})
		return
	}
	todo.Completed = !todo.Completed
	context.IndentedJSON(http.StatusOK, todo)

}
func main() {

	router := gin.Default()
	router.GET("/todos", getTodos)
	router.GET("/todos/:id", getTodo)
	router.POST("todos", addTodos)
	router.PATCH("/todos/:id", updateTodoStatus)
	router.Run("localhost:5000")

}
