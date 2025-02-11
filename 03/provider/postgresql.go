package provider

import (
	"03/utils/config"
	"fmt"

	"03/utils/db"
)

func ProvidePostgres(config config.ServiceConfig) *db.Database {
	fmt.Println("Connecting to PostgreSQL...")
	return db.New(db.Connection{
		Host:     config.DBHost,
		Port:     config.DBPort,
		Database: config.DBName,
		User:     config.DBUser,
		Password: config.DBPwd,
		SSLMode:  db.Disable,
	})
}

type AppProvider struct {
	Postgres *db.Database
}

func NewAppProvider(config config.ServiceConfig) *AppProvider {
	return &AppProvider{
		Postgres: ProvidePostgres(config),
	}
}
