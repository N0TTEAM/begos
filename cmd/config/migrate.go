package main

import (
	"flag"
	"log"

	"github.com/N0TTEAM/begos/internal/config"
	"github.com/N0TTEAM/begos/internal/db"
	"github.com/N0TTEAM/begos/internal/http/model"
	"gorm.io/gorm"
)

func main() {
	fresh := flag.Bool("fresh", false, "Drop all tables and recreate")
	flag.Parse()

	cfg := config.LoadConf()
	database := db.NewConnection(&cfg.Postgres)

	sqlDB, err := database.DB()
	if err != nil {
		log.Fatalf("Failed to get database instance: %v", err)
	}
	defer func() {
		if err := sqlDB.Close(); err != nil {
			log.Fatalf("Failed to close database connection: %v", err)
		}
	}()

	if *fresh {
		log.Println("Fresh Migrating")
		if err := dropAllTables(database); err != nil {
			log.Fatalf("failed to drop table: %v", err)
		}
	}

	log.Println("Running migrations...")
	err = database.AutoMigrate(
		model.GetAllModels()...,
	)

	if err != nil {
		log.Fatalf("Migration failed: %v", err)
	}
	log.Println("Migration completed success")
}

func dropAllTables(database *gorm.DB) error {
	err := database.Migrator().DropTable(
		model.GetAllModels()...,
	)
	return err
}
