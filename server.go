package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
)

/*
{
	"userId": 1,
	"id": 1,
	"title": "delectus aut autem",
	"completed": false
}K
*/

type User struct {
	UserID    int    `json:"userId"`
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func main() {
	gin.ForceConsoleColor()
	router := gin.Default()
	router.GET("/api/v1/todos", getTodos)
	router.Run("localhost:8080")
}

func getTodos(c *gin.Context) {
	log.Println("Received API call")
	cl := http.DefaultClient
	r, e := cl.Get("https://jsonplaceholder.typicode.com/todos/1")
	if e != nil {
		log.Fatal("Error calling JSON API")
		c.Status(http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()
	body, e := ioutil.ReadAll(r.Body)
	var user User
	json.Unmarshal(body, &user)
	c.IndentedJSON(http.StatusOK, user)
}
