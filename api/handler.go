package api

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/juancassiano/api_students/schemas"
	"github.com/labstack/echo"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

func (api *API) getStudents(c echo.Context) error {
	students, err := api.DB.GetStudents()
	if err != nil {
		return c.String(http.StatusNotFound, "Students not found")

	}
	return c.JSON(http.StatusOK, students)
}

func (api *API) createStudent(c echo.Context) error {
	studentRequest := StudentRequest{}

	if err := c.Bind(&studentRequest); err != nil {
		return err
	}
	if err := studentRequest.Validate(); err != nil {
		log.Error().Err(err).Msgf("Failed to validate student request")
		return c.String(http.StatusBadRequest, err.Error())
	}

	student := schemas.Student{
		Name:   studentRequest.Name,
		CPF:    studentRequest.CPF,
		Email:  studentRequest.Email,
		Age:    studentRequest.Age,
		Active: *studentRequest.Active,
	}

	if err := api.DB.AddStudent(student); err != nil {
		return c.String(http.StatusInternalServerError, "Error creating student")
	}

	return c.JSON(http.StatusCreated, student)
}

func (api *API) getStudent(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.String(http.StatusInternalServerError, "Failed to get Student id")
	}

	student, err := api.DB.GetStudent(id)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return c.String(http.StatusNotFound, "Student not found")
	}

	return c.JSON(http.StatusOK, student)
}

func (api *API) updateStudent(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, "Student updated with id: "+id)
}

func (api *API) deleteStudent(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	student, err := api.DB.GetStudent(id)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Failed to get Student id")
	}
	api.DB.DeleteStudent(student)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return c.String(http.StatusNotFound, "Student not found")
	}

	return c.JSON(http.StatusNoContent, nil)
}
