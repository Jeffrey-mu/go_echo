package main

import (
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

var (
	newFile *os.File
	err     error
)

func main() {
	newFile, err = os.Create("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	log.Println(newFile)
	newFile.Close()
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.GET("/users/:id", getUser)
	e.GET("/show", show)
	e.Logger.Fatal(e.Start(":1323"))
}

// e.GET("/users/:id", getUser)
func getUser(c echo.Context) error {
	// User ID from path `users/:id`
	id := c.Param("id")
	return c.String(http.StatusOK, id)
}
func show(c echo.Context) error {
	// Get name and data from the query string
	name := c.QueryParam("name")
	data := c.QueryParam("data")
	log.Printf("[DEBUG] User: %v Data:(%v)", name, data)
	return c.String(http.StatusOK, "name:"+name+", data:"+data)
}
