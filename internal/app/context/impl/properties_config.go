package impl

import (
	"github.com/joho/godotenv"
	"goproject/internal/app/context"
	"log"
	"os"
)

type propertiesConfig struct {
	profile     string
	databaseUrl string
}

func newPropertiesConfig() context.PropertiesConfig {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Failed to load env: ", err)
	}

	profile, _ := os.LookupEnv("profile")
	databaseUrl, _ := os.LookupEnv("database_url")

	return &propertiesConfig{
		profile,
		databaseUrl,
	}
}

func (config *propertiesConfig) GetProfile() string {
	return config.profile
}

func (config *propertiesConfig) GetDatabaseUrl() string {
	return config.databaseUrl
}
