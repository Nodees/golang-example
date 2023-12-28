package config

import (
	"fmt"

	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/gofiber/contrib/casbin"
	_ "gorm.io/driver/postgres"
)

func CasbinConfig(conf *Config) *casbin.Middleware {

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s", conf.DBHost, conf.DBUsername, conf.DBPassword, conf.DBName, conf.DBPort, conf.Sslmode)

	pollicyadapter, _ := gormadapter.NewAdapter(conf.Driver, dsn)

	return casbin.New(casbin.Config{
		ModelFilePath: "auth.model.conf",
		PolicyAdapter: pollicyadapter,
	})
}
