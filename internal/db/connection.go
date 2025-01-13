package db

import (
	"log"
	"os"

	"github.com/ricardofedrigo/everyone_against_everyone/internal/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

var err error

func MigrateDB() {
	// Realiza a migração do modelo para o banco de dados
	err := db.AutoMigrate(&model.Answer{}, &model.Question{}, &model.Deck{})
	if err != nil {
		log.Fatalf("Erro while migrate: %v", err)
	}
	log.Println("Sucess to migrate!")
}

func InitDb() (*gorm.DB, error) {
	dns := os.Getenv("../../.env")
	if dns == "" {
		log.Fatalln("dot env not set")
	}

	db, err = gorm.Open(postgres.Open(dns), &gorm.Config{})

	if err != nil {
		log.Fatalln("Database not connect, error: %v", err)
		return nil, err
	}

	return db, nil
}
