package database

import (
	"fmt"
	"github/ariifysp/go-101/config"
	"log"
	"sync"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type (
	DatabaseInterface interface {
		Connect() *gorm.DB
	}

	PostgresDatabase struct {
		*gorm.DB
	}
)

var (
	once                     sync.Once
	postgresDatabaseInstance *PostgresDatabase
)

func NewPostgresDatabase(config *config.Postgres) DatabaseInterface {
	once.Do(func() {
		connectionURL := GeneratePostgresConnectionURL(config)

		connect, err := gorm.Open(postgres.Open(connectionURL), &gorm.Config{})
		if err != nil {
			panic(err)
		}

		log.Printf("Connected to database %s", config.Database)

		postgresDatabaseInstance = &PostgresDatabase{connect}
	})

	return postgresDatabaseInstance
}

func GeneratePostgresConnectionURL(config *config.Postgres) string {
	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s search_path=%s",
		config.Host,
		config.Port,
		config.Username,
		config.Password,
		config.Database,
		config.SSLMode,
		config.Schema,
	)
}

func (db *PostgresDatabase) Connect() *gorm.DB {
	return db.DB
}
