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
}
*/

type Todo struct {
	UserID    int    `json:"userId"`
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

/*
{
    "id": 1,
    "name": "Leanne Graham",
    "username": "Bret",
    "email": "Sincere@april.biz",
    "address": {
      "street": "Kulas Light",
      "suite": "Apt. 556",
      "city": "Gwenborough",
      "zipcode": "92998-3874",
      "geo": {
        "lat": "-37.3159",
        "lng": "81.1496"
      }
    },
    "phone": "1-770-736-8031 x56442",
    "website": "hildegard.org",
    "company": {
      "name": "Romaguera-Crona",
      "catchPhrase": "Multi-layered client-server neural-net",
      "bs": "harness real-time e-markets"
    }
  },
*/

type User struct {
	UserID   int    `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    bool   `json:"email"`
	Address Address `json:"address"`
	Company Company `json:"company"`
}

type Address struct {
	Street string `json:"street"`
	Suite string `json:"suite"`
	City string `json:"city"`
	Zipcode string `json:"zipcode"`
}

type Company struct {
	Name string `json:"name"`
}

func main() {
	gin.ForceConsoleColor()
	router := gin.Default()
	router.GET("/api/v1/todos", getTodos)
	router.GET("/api/v1/users", getUsers)
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
	var todo Todo
	json.Unmarshal(body, &todo)
	c.IndentedJSON(http.StatusOK, todo)
}

func getUsers(c *gin.Context) {
	log.Println("Received API call")
	cl := http.DefaultClient
	r, e := cl.Get("https://jsonplaceholder.typicode.com/users/1")
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
