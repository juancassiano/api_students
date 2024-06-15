package api

import (
	"net/http"

	"github.com/juancassiano/api_students/db"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type API struct {
	Echo *echo.Echo
	DB   *db.StudentHandler
}

func NewServer() *API {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	database := db.Init()
	studentDB := db.NewStudenHandler(database)

	return &API{
		Echo: e,
		DB:   studentDB,
	}
}

func (api *API) Start() error {
	return api.Echo.Start(":8080")
}

func (api *API) ConfigureRoutes() {
	api.Echo.GET("/students", api.getStudents)
	api.Echo.POST("/students", api.createStudent)
	api.Echo.GET("/students/:id", api.getStudent)
	api.Echo.PUT("/students/:id", api.updateStudent)
	api.Echo.DELETE("/students/:id", api.deleteStudent)
}

func (api *API) getStudents(c echo.Context) error {
	students, err := api.DB.GetStudents()
	if err != nil {
		return c.String(http.StatusNotFound, "Students not found")

	}
	return c.JSON(http.StatusOK, students)
}

func (api *API) createStudent(c echo.Context) error {
	student := db.Student{}
	if err := c.Bind(&student); err != nil {
		return err
	}

	if err := api.DB.AddStudent(student); err != nil {
		return c.String(http.StatusInternalServerError, "Error creating student")
	}

	return c.String(http.StatusCreated, "Student created")
}

func (api *API) getStudent(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, "Student with id: "+id)
}

func (api *API) updateStudent(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, "Student updated with id: "+id)
}

func (api *API) deleteStudent(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, "Student deleted with id: "+id)
}
