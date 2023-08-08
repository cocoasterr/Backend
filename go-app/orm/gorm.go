package orm

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
func ConnectDB(){
	var err error
	dsn := "postgresql://postgres:password@172.21.255.76:5439/go_app?sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	DB = db
	if err != nil {
		panic("Failed to connect to database")
	}
}