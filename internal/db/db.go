package db

import (
	"fmt"
	"log"

	"github.com/N0TTEAM/begos/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewConnection(cfg *config.Postgres) *gorm.DB {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Dbname, cfg.SslMode,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect db: %v", err)
	}

	log.Println("connected")
	return db
}
