package PGConfig

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var GORMDB *gorm.DB

func PGGormConfig() *gorm.DB {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Cant connect .env!")
	}
	db_url := os.Getenv("PG_DB_URL")
	db, err := gorm.Open(postgres.Open(db_url), &gorm.Config{})
	if err != nil {
		log.Fatal("Unable to connect Postgres!")
	}
	GORMDB = db
	log.Println("Postgres Connected!")
	return GORMDB
}
