package config

import "os"

type Config struct {
	// AppEnv   string
	// IsDevEnv bool

	PostgresHost     string
	PostgresUser     string
	PostgresPassword string
	PostgresDb       string
	PostgresPort     string
}

var AppConfig Config

func LoadConfig() {
	AppConfig.PostgresHost = os.Getenv("POSTGRES_HOST")
	AppConfig.PostgresUser = os.Getenv("POSTGRES_USER")
	AppConfig.PostgresPassword = os.Getenv("POSTGRES_PASSWORD")
	AppConfig.PostgresDb = os.Getenv("POSTGRES_DB")
	AppConfig.PostgresPort = os.Getenv("POSTGRES_PORT")
}
