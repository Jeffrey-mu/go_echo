package main

import (
	"io"
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
	e.POST("/save", save)
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

// func save(c echo.Context) error {
// 	name := c.FormValue("name")
// 	email := c.FormValue("email")
// 	log.Println(name, email)
// 	return c.String(http.StatusOK, "name:"+name+", email:"+email)
// }
func save(c echo.Context) error {
	// Get name
	name := c.FormValue("name")
	// Get avatar
	avatar, err := c.FormFile("avatar")
	if err != nil {
		return err
	}

	// Source
	src, err := avatar.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	// Destination
	dst, err := os.Create(avatar.Filename)
	if err != nil {
		return err
	}
	defer dst.Close()

	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		return err
	}

	return c.HTML(http.StatusOK, "<b>Thank you! "+name+"</b>")
}
