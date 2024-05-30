package api

import (
	"net/http"

	"github.com/juancassiano/api_students/db"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"gorm.io/gorm"
)

type API struct {
	Echo *echo.Echo
	DB   *gorm.DB
}

func NewServer() *API {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	db := db.Init()

	return &API{
		Echo: e,
		DB:   db,
	}
}

func (api *API) Start() error {
	return api.Echo.Start(":8080")
}

func (api *API) ConfigureRoutes() {
	api.Echo.GET("/students", getStudents)
	api.Echo.POST("/students", createStudent)
	api.Echo.GET("/students/:id", getStudent)
	api.Echo.PUT("/students/:id", updateStudent)
	api.Echo.DELETE("/students/:id", deleteStudent)
}

func getStudents(c echo.Context) error {
	students, err := db.GetStudents()
	if err != nil {
		return c.String(http.StatusNotFound, "Students not found")

	}
	return c.JSON(http.StatusOK, students)
}

func createStudent(c echo.Context) error {
	student := db.Student{}
	if err := c.Bind(&student); err != nil {
		return err
	}

	if err := db.AddStudent(student); err != nil {
		return c.String(http.StatusInternalServerError, "Error creating student")
	}

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
