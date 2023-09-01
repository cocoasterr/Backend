package postgre

import (
	"database/sql"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func PGConfig() (*sql.DB, error){
	var err error
	if err = godotenv.Load(); err!=nil{
		log.Fatal("Failed to connect .env")
		return nil, err
	}
	DB_DSN := os.Getenv("PG_DB_URL")
	db,_ := sql.Open("postgres", DB_DSN)

	if err = db.Ping(); err != nil {
		db.Close()
		log.Fatal("Failed to ping database")
		return nil, err
	}
	log.Println("Postgres Connected!")
	return db, nil
}