package postgres

import (
	"core/config"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitPostgresDB(conf *config.Config) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s", conf.DBHost, conf.DBUsername, conf.DBPassword, conf.DBName, conf.DBPort, conf.Sslmode)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Erro ao connectar ao banco de dados", err)
	}

	DB = db
}
