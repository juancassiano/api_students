package db

import (
	"log"

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
		log.Fatalln(err)
	}

	db.AutoMigrate(&Student{})

	return db
}

func NewStudenHandler(db *gorm.DB) *StudentHandler {
	return &StudentHandler{DB: db}
}

func (s *StudentHandler) AddStudent(student Student) error {

	if result := s.DB.Create(&student); result.Error != nil {
		return result.Error
	}
	return nil
}

func (s *StudentHandler) GetStudents() ([]Student, error) {
	students := []Student{}

	err := s.DB.Find(&students).Error

	return students, err
}
