package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Store struct {
	ID    string `json: "id"`
	Book  string `json: "id"`
	Note  string `json: "id"`
	Paper string `json: "id"`
}

func main() {
	router := gin.Default()
	router.GET("/albums", Get_stores)
	router.GET("/albums/:id", Get_ID)
	router.POST("/albums", Post_stores)
	router.Run("localhost:8080")
}

var Stores = []Store{
	{ID: "7", Book: "Wingsoffire", Note: "classmate", Paper: "the hindu"},
	{ID: "8", Book: "bhagavatgita", Note: "flair", Paper: "indianexpress"},
	{ID: "9", Book: "bible", Note: "kaka", Paper: "Thanthi"},
}

func Get_stores(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, Stores)
}
func Post_stores(c *gin.Context) {
	var Newstore Store
	if error := c.BindJSON(&Newstore); error != nil {
		return
	}
	Stores = append(Stores, Newstore)
	c.IndentedJSON(http.StatusCreated, Newstore)
}
func Get_ID(c *gin.Context) {
	id := c.Param("id")
	for _, a := range Stores {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Id not found"})
}
