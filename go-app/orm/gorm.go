package orm

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
func ConnectDB()(){
	var err error
	if err = godotenv.Load(); err != nil{
		log.Fatal(err.Error())
	}
	dsn := os.Getenv("DB_DSN")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	DB = db
	if err != nil {
		panic("Failed to connect to database")
	}
}