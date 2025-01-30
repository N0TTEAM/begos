package db

import (
	"fmt"
	"log"

	"github.com/N0TTEAM/begos/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func NewConnection(cfg *config.Postgres) {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Dbname, cfg.SslMode,
	)
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect db: %v", err)
	}

	log.Println("Database Connected")

}

func GetDB() *gorm.DB {
	if DB == nil {
		log.Fatalf("Database connection is not initialized")
	}
	return DB
}
