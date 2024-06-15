package db

import (
	"github.com/rs/zerolog/log"

	"gorm.io/driver/sqlite" // Sqlite driver based on CGO
	"gorm.io/gorm"
)

type StudentHandler struct {
	DB *gorm.DB
}

type Student struct {
	gorm.Model
	Name   string `json:"name"`
	CPF    string `json:"cpf"`
	Email  string `json:"email"`
	Age    int    `json:"age"`
	Active bool   `json:"active"`
}

func Init() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("Student.db"), &gorm.Config{})

	if err != nil {
		log.Fatal().Err(err).Msgf("Failed to initialize SQLite: %s", err.Error())
	}

	db.AutoMigrate(&Student{})

	return db
}

func NewStudenHandler(db *gorm.DB) *StudentHandler {
	return &StudentHandler{DB: db}
}

func (s *StudentHandler) AddStudent(student Student) error {

	if result := s.DB.Create(&student); result.Error != nil {
		log.Error().Msg("Failed to add student")
		return result.Error
	}
	log.Info().Msgf("Student %s added", student.Name)
	return nil
}

func (s *StudentHandler) GetStudents() ([]Student, error) {
	students := []Student{}

	err := s.DB.Find(&students).Error

	return students, err
}

func (s *StudentHandler) GetStudent(id int) (Student, error) {
	var student Student
	err := s.DB.First(&student, id)

	return student, err.Error
}
