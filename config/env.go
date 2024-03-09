package config

import (
	"os"
)

type PostgresEnv struct {
	host     string
	port     string
	database string
	user     string
	password string
}

func NewPostgresEnv() *PostgresEnv {
	return &PostgresEnv{
		host:     os.Getenv("POSTGRES_HOST"),
		port:     os.Getenv("POSTGRES_PORT"),
		database: os.Getenv("POSTGRES_DB"),
		user:     os.Getenv("POSTGRES_USER"),
		password: os.Getenv("POSTGRES_PASSWORD"),
	}
}

func (pe *PostgresEnv) GetHostOfPrivateValue() string {
	return pe.host
}

func (pe *PostgresEnv) GetPortOfPrivateValue() string {
	return pe.port
}

func (pe *PostgresEnv) GetDatabaseOfPrivateValue() string {
	return pe.database
}

func (pe *PostgresEnv) GetUserOfPrivateValue() string {
	return pe.user
}

func (pe *PostgresEnv) GetPasswordOfPrivateValue() string {
	return pe.password
}
