package postgres

import (
	"core/configs"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitPostgresDB(env *configs.Env) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s", env.DBHost, env.DBUsername, env.DBPassword, env.DBName, env.DBPort, env.Sslmode)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Erro ao connectar ao banco de dados", err)
	}

	DB = db
}
