package config

import (
	"core/utils"
	"fmt"
	"log"
	"strings"

	// gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/gofiber/fiber/v2"
	_ "gorm.io/driver/postgres"
)

func Authenticate(conf *Config) func(*fiber.Ctx) error {
	// dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s", conf.DBHost, conf.DBUsername, conf.DBPassword, conf.DBName, conf.DBPort, conf.Sslmode)

	// pollicyadapter, _ := gormadapter.NewAdapter(conf.Driver, dsn)

	return func(c *fiber.Ctx) error {
		authorization := c.Get("Authorization")

		tokenString := strings.TrimPrefix(authorization, "Bearer ")
		method := c.Method()
		path := c.Path()

		if tokenString != "" {
			claim, _ := utils.GetUserFromToken(tokenString, conf.JwtSecret)
			log.Fatal(fmt.Sprint(claim))
			log.Fatal(fmt.Sprint(method))
			log.Fatal(fmt.Sprint(path))
		}
		return c.Next()
	}

	// return casbin.New(casbin.Config{
	// 	ModelFilePath: "auth.model.conf",
	// 	PolicyAdapter: pollicyadapter,
	// })
}
