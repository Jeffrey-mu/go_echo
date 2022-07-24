package main

import (
	"database/sql"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
)

var (
	err error
)

func main() {
	db, err := sql.Open("mysql", "admin:mc1009jf1018.@21729(bj-cynosdbmysql-grp-0o0dqcfy.sql.tencentcdb.com)/admin")
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	Db := db
	data, err := Db.Query("SELECT * FROM user where id = 1")
	// row, _ := data.Columns()
	// println(row[0])
	println(data.Columns())
	Db.Close()
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.GET("/users/:id", getUser)
	e.GET("/show", show)
	e.POST("/save", save)
	e.Static("/static", "static")
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
