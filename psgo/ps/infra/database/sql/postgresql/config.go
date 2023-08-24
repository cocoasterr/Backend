package configPostgres

import (
	"database/sql"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)


var DB *sql.DB


func ConnectDB()*sql.DB{
	if err := godotenv.Load();err != nil{
		log.Fatal("failed to connect .env")
	}
	db_dsn := os.Getenv("DB_DSN")
	db, err := sql.Open("postgres", db_dsn)
	if err != nil{
		log.Fatal(err)
	}
	DB = db
	log.Println("Postgres DB Connected!")
	return	DB
}