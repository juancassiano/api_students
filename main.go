package main

import (
	"net/http"

	"github.com/juancassiano/api_students/db"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/students", getStudents)
	e.POST("/students", createStudent)
	e.GET("/students/:id", getStudent)
	e.PUT("/students/:id", updateStudent)
	e.DELETE("/students/:id", deleteStudent)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}

// Handler
func getStudents(c echo.Context) error {
	return c.String(http.StatusOK, "List of all students")
}

func createStudent(c echo.Context) error {
	student := db.Student{}
	if err := c.Bind(&student); err != nil {
		return err
	}

	db.AddStudent(student)
	return c.String(http.StatusCreated, "Student created")
}

func getStudent(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, "Student with id: "+id)
}

func updateStudent(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, "Student updated with id: "+id)
}

func deleteStudent(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, "Student deleted with id: "+id)
}
