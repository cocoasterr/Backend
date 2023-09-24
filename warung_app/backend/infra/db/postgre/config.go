package PgConfig

import (
	"log"
	"os"

	domain "github.com/cocoasterr/warung_app/domain"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func PGinit() (*gorm.DB, error) {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Failed to connect .env")
	}

	db_dsn := os.Getenv("PG_URL")
	db, err := gorm.Open(postgres.Open(db_dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Unable to connect Postgres - gorm")
		return nil, err
	}
	log.Println("Connected to Postgres - gorm")
	db.AutoMigrate(
		&domain.User{},
	)
	log.Println("migrationg to Postgres success!")
	return db, nil
}
