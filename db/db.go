package db

import (
	"fmt"
	"log"
	"os"

	"github.com/mustafaakilll/go-site-exam/db/entity"
	"github.com/mustafaakilll/go-site-exam/db/seeders"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/logger"
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

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

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
	if err := Seed(); err != nil {
		fmt.Println("failed to seed database")
	}
	fmt.Println("Database Seeded")
}

func Seed() error {
	userSeeder := seeders.UserTypeSeeder{}
	return DB.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "id"}},
		UpdateAll: true,
	}).Create(userSeeder.Run()).Error
}
