package models

import (
	"database/sql"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)



func loadEnv()string{
	if err := godotenv.Load(); err != nil{
		log.Fatal(".env failed!")
	}
	db_dsn := os.Getenv("DB_DSN")
	return db_dsn
}


//database sql config
var DB *sql.DB

func ConnectDB()(*sql.DB){
	db_dsn := loadEnv()
	db, err := sql.Open("postgres", db_dsn)
    if err != nil {
        return nil
    }
	if err = db.Ping(); err != nil {
        db.Close()
        return nil
    }
	DB = db
	log.Println("Database Connected!")
	return db
}


//Gorm conig
var DBORM *gorm.DB

func OrmConnectDB()(){
	db_dsn := loadEnv()
	db, err := gorm.Open(postgres.Open(db_dsn), &gorm.Config{})
	DBORM = db
	if err != nil {
		panic("Failed to connect to database")
	}
}