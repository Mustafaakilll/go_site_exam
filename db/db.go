package db

import (
	"fmt"

	"src/github.com/mustafaakilll/go-site-exam/db/entity"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	var err error
	dsn := "host=localhost user=postgres password=password dbname=examsite port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	if err := DB.AutoMigrate(
		&entity.User{},
		&entity.Answer{},
		&entity.Quiz{},
		&entity.Question{},
		&entity.Lesson{},
		&entity.Code{},
		&entity.CodeAnswer{},
		&entity.CodeSubmission{},
		&entity.Choice{},
		&entity.UserQuiz{},
		&entity.UserAnswer{},
		&entity.UserType{},
	); err != nil {
		panic("failed to migrate database")
	}

	fmt.Println("Connection Opened to Database")
}

func Migrate() {
	// Migrate database
}
