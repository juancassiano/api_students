package db

import (
	"fmt"
	"log"

	"gorm.io/driver/sqlite" // Sqlite driver based on CGO
	"gorm.io/gorm"
)

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

func AddStudent(student Student) {
	db := Init()

	if result := db.Create(&student); result.Error != nil {
		log.Fatalln(result.Error)
	}
	fmt.Println("Student created with ID: ", student.ID)
}
