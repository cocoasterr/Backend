package models

import (
	"database/sql"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var DB *sql.DB

func ConnectDB()(*sql.DB){
	if err := godotenv.Load(); err != nil{
		log.Fatal(".env failed!")
	}
	db_dsn := os.Getenv("DB_DSN")
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