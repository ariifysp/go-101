package config

import (
	"os"
	"sync"

	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

type Config struct {
	App      Fiber
	Postgres Postgres
	DB       *gorm.DB
	Redis    Redis
}

type Fiber struct {
	Host string
	Port string
}

type Postgres struct {
	Host     string
	Port     string
	Username string
	Password string
	Database string
	SSLMode  string
	Schema   string
}

type Redis struct {
	Host string
	Port string
}

var (
	once           sync.Once
	configInstance *Config
)

func GetConfig() *Config {
	once.Do(func() {
		if err := godotenv.Load(".env"); err != nil {
			panic(err.Error())
		}

		config := new(Config)

		config.App.Host = os.Getenv("FIBER_HOST")
		config.App.Port = os.Getenv("FIBER_PORT")

		config.Postgres.Host = os.Getenv("DB_HOST")
		config.Postgres.Port = os.Getenv("DB_PORT")
		config.Postgres.Username = os.Getenv("DB_USERNAME")
		config.Postgres.Password = os.Getenv("DB_PASSWORD")
		config.Postgres.Database = os.Getenv("DB_DATABASE")
		config.Postgres.SSLMode = os.Getenv("DB_SSLMODE")
		config.Postgres.Schema = os.Getenv("DB_SCHEMA")

		config.Redis.Host = os.Getenv("REDIS_HOST")
		config.Redis.Port = os.Getenv("REDIS_PORT")

		configInstance = config
	})

	return configInstance
}
