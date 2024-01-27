package storage

import (
	"app/pkg/storage/postgres"
)

func ConnectDB() {
	postgres.Postgres()
}
