package sqlPostgreRepository

import (
	"database/sql"
)


type RepositoryContract interface{
	DB() *sql.DB
}
