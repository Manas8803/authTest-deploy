package storage

import (
	"auth-service/pkg/storage/postgres"
)

func ConnectDB() {
	postgres.Postgres()
}
