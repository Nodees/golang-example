package postgres

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	host := "localhost"
	user := "postgres"
	dbpassword := "fpf2023"
	dbname := "Go"
	port := "5433"
	sslmode := "disable"

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s", host, user, dbpassword, dbname, port, sslmode)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Erro ao connectar ao banco de dados", err)
	}

	DB = db

	sqlDB, err := db.DB()

	// Verificar a conexão
	if err != nil {
		log.Fatal("Erro ao obter o DB:", err)
		sqlDB.Close()
	}
}
