package database

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func ConnectDB() {
	var err error
	connStr := "host=localhost user=root password=root dbname=root port=5433 sslmode=disable"
	db, err = gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		log.Fatalf("Erro ao conectar ao banco de dados: %v", err)
	}
	fmt.Println("Conectado ao banco de dados!")
}

func GetDB() *gorm.DB {
	return db
}
