/***********************************************************
* Dewdrop - A User Registration Web Service
*
* Module: Server.go
* This is the main module of Dewdrop Service.
* It consists of
*    (a) Echo Web Service instantiation
*    (b) Web HTTP method handlers
*
* DevOps Course, Nov 2021
*
* Tech Stack : Golang, Mongo
***********************************************************/

package main

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	// Create a new Echo web server instance
	e := echo.New()

	// Initialize required handlers
	log.Print("Initialize..")
	initialization(e)
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	// Start web server at port 3333
	e.Logger.Fatal(e.Start(":3333"))
}

func initialization(e *echo.Echo) {
	log.Print("Executing Initalization tasks")
	e.POST("/registerUser", registerUser)
	//e.GET("/users/:id", getUser)
	//e.PUT("/users/:id", updateUser)
	//e.DELETE("/users/:id", deleteUser)
}

// e.POST("/registerUser", registerUser)
func registerUser(c echo.Context) error {
	// Get name and email
	name := c.FormValue("name")
	email := c.FormValue("email")
	return c.String(http.StatusOK, "name:"+name+", email:"+email)
}
