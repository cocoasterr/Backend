package repository

import (
	"database/sql"
)


type RepositoryContract interface{
	DB() *sql.DB
}
