package db

import (
	"fmt"
	"log"
	"os"

	"github.com/mustafaakilll/go-site-exam/db/entity"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	database := os.Getenv("DB_DATABASE")
	port := os.Getenv("DB_PORT")
	var err error
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%v sslmode=disable",
		host,
		user,
		password,
		database,
		port,
	)
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Panicf("failed to connect database,err: %v", err)
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
